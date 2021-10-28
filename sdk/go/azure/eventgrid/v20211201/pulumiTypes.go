


package v20211201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureFunctionEventSubscriptionDestination struct {
	DeliveryAttributeMappings     []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType                  string        `pulumi:"endpointType"`
	MaxEventsPerBatch             *int          `pulumi:"maxEventsPerBatch"`
	PreferredBatchSizeInKilobytes *int          `pulumi:"preferredBatchSizeInKilobytes"`
	ResourceId                    *string       `pulumi:"resourceId"`
}





type AzureFunctionEventSubscriptionDestinationInput interface {
	pulumi.Input

	ToAzureFunctionEventSubscriptionDestinationOutput() AzureFunctionEventSubscriptionDestinationOutput
	ToAzureFunctionEventSubscriptionDestinationOutputWithContext(context.Context) AzureFunctionEventSubscriptionDestinationOutput
}

type AzureFunctionEventSubscriptionDestinationArgs struct {
	DeliveryAttributeMappings     pulumi.ArrayInput     `pulumi:"deliveryAttributeMappings"`
	EndpointType                  pulumi.StringInput    `pulumi:"endpointType"`
	MaxEventsPerBatch             pulumi.IntPtrInput    `pulumi:"maxEventsPerBatch"`
	PreferredBatchSizeInKilobytes pulumi.IntPtrInput    `pulumi:"preferredBatchSizeInKilobytes"`
	ResourceId                    pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (AzureFunctionEventSubscriptionDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFunctionEventSubscriptionDestination)(nil)).Elem()
}

func (i AzureFunctionEventSubscriptionDestinationArgs) ToAzureFunctionEventSubscriptionDestinationOutput() AzureFunctionEventSubscriptionDestinationOutput {
	return i.ToAzureFunctionEventSubscriptionDestinationOutputWithContext(context.Background())
}

func (i AzureFunctionEventSubscriptionDestinationArgs) ToAzureFunctionEventSubscriptionDestinationOutputWithContext(ctx context.Context) AzureFunctionEventSubscriptionDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFunctionEventSubscriptionDestinationOutput)
}

type AzureFunctionEventSubscriptionDestinationOutput struct{ *pulumi.OutputState }

func (AzureFunctionEventSubscriptionDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFunctionEventSubscriptionDestination)(nil)).Elem()
}

func (o AzureFunctionEventSubscriptionDestinationOutput) ToAzureFunctionEventSubscriptionDestinationOutput() AzureFunctionEventSubscriptionDestinationOutput {
	return o
}

func (o AzureFunctionEventSubscriptionDestinationOutput) ToAzureFunctionEventSubscriptionDestinationOutputWithContext(ctx context.Context) AzureFunctionEventSubscriptionDestinationOutput {
	return o
}

func (o AzureFunctionEventSubscriptionDestinationOutput) DeliveryAttributeMappings() pulumi.ArrayOutput {
	return o.ApplyT(func(v AzureFunctionEventSubscriptionDestination) []interface{} { return v.DeliveryAttributeMappings }).(pulumi.ArrayOutput)
}

func (o AzureFunctionEventSubscriptionDestinationOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFunctionEventSubscriptionDestination) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o AzureFunctionEventSubscriptionDestinationOutput) MaxEventsPerBatch() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureFunctionEventSubscriptionDestination) *int { return v.MaxEventsPerBatch }).(pulumi.IntPtrOutput)
}

func (o AzureFunctionEventSubscriptionDestinationOutput) PreferredBatchSizeInKilobytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureFunctionEventSubscriptionDestination) *int { return v.PreferredBatchSizeInKilobytes }).(pulumi.IntPtrOutput)
}

func (o AzureFunctionEventSubscriptionDestinationOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFunctionEventSubscriptionDestination) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type AzureFunctionEventSubscriptionDestinationResponse struct {
	DeliveryAttributeMappings     []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType                  string        `pulumi:"endpointType"`
	MaxEventsPerBatch             *int          `pulumi:"maxEventsPerBatch"`
	PreferredBatchSizeInKilobytes *int          `pulumi:"preferredBatchSizeInKilobytes"`
	ResourceId                    *string       `pulumi:"resourceId"`
}





type AzureFunctionEventSubscriptionDestinationResponseInput interface {
	pulumi.Input

	ToAzureFunctionEventSubscriptionDestinationResponseOutput() AzureFunctionEventSubscriptionDestinationResponseOutput
	ToAzureFunctionEventSubscriptionDestinationResponseOutputWithContext(context.Context) AzureFunctionEventSubscriptionDestinationResponseOutput
}

type AzureFunctionEventSubscriptionDestinationResponseArgs struct {
	DeliveryAttributeMappings     pulumi.ArrayInput     `pulumi:"deliveryAttributeMappings"`
	EndpointType                  pulumi.StringInput    `pulumi:"endpointType"`
	MaxEventsPerBatch             pulumi.IntPtrInput    `pulumi:"maxEventsPerBatch"`
	PreferredBatchSizeInKilobytes pulumi.IntPtrInput    `pulumi:"preferredBatchSizeInKilobytes"`
	ResourceId                    pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (AzureFunctionEventSubscriptionDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFunctionEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (i AzureFunctionEventSubscriptionDestinationResponseArgs) ToAzureFunctionEventSubscriptionDestinationResponseOutput() AzureFunctionEventSubscriptionDestinationResponseOutput {
	return i.ToAzureFunctionEventSubscriptionDestinationResponseOutputWithContext(context.Background())
}

func (i AzureFunctionEventSubscriptionDestinationResponseArgs) ToAzureFunctionEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) AzureFunctionEventSubscriptionDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFunctionEventSubscriptionDestinationResponseOutput)
}

type AzureFunctionEventSubscriptionDestinationResponseOutput struct{ *pulumi.OutputState }

func (AzureFunctionEventSubscriptionDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFunctionEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (o AzureFunctionEventSubscriptionDestinationResponseOutput) ToAzureFunctionEventSubscriptionDestinationResponseOutput() AzureFunctionEventSubscriptionDestinationResponseOutput {
	return o
}

func (o AzureFunctionEventSubscriptionDestinationResponseOutput) ToAzureFunctionEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) AzureFunctionEventSubscriptionDestinationResponseOutput {
	return o
}

func (o AzureFunctionEventSubscriptionDestinationResponseOutput) DeliveryAttributeMappings() pulumi.ArrayOutput {
	return o.ApplyT(func(v AzureFunctionEventSubscriptionDestinationResponse) []interface{} {
		return v.DeliveryAttributeMappings
	}).(pulumi.ArrayOutput)
}

func (o AzureFunctionEventSubscriptionDestinationResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFunctionEventSubscriptionDestinationResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o AzureFunctionEventSubscriptionDestinationResponseOutput) MaxEventsPerBatch() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureFunctionEventSubscriptionDestinationResponse) *int { return v.MaxEventsPerBatch }).(pulumi.IntPtrOutput)
}

func (o AzureFunctionEventSubscriptionDestinationResponseOutput) PreferredBatchSizeInKilobytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureFunctionEventSubscriptionDestinationResponse) *int { return v.PreferredBatchSizeInKilobytes }).(pulumi.IntPtrOutput)
}

func (o AzureFunctionEventSubscriptionDestinationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFunctionEventSubscriptionDestinationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type BoolEqualsAdvancedFilter struct {
	Key          *string `pulumi:"key"`
	OperatorType string  `pulumi:"operatorType"`
	Value        *bool   `pulumi:"value"`
}





type BoolEqualsAdvancedFilterInput interface {
	pulumi.Input

	ToBoolEqualsAdvancedFilterOutput() BoolEqualsAdvancedFilterOutput
	ToBoolEqualsAdvancedFilterOutputWithContext(context.Context) BoolEqualsAdvancedFilterOutput
}

type BoolEqualsAdvancedFilterArgs struct {
	Key          pulumi.StringPtrInput `pulumi:"key"`
	OperatorType pulumi.StringInput    `pulumi:"operatorType"`
	Value        pulumi.BoolPtrInput   `pulumi:"value"`
}

func (BoolEqualsAdvancedFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BoolEqualsAdvancedFilter)(nil)).Elem()
}

func (i BoolEqualsAdvancedFilterArgs) ToBoolEqualsAdvancedFilterOutput() BoolEqualsAdvancedFilterOutput {
	return i.ToBoolEqualsAdvancedFilterOutputWithContext(context.Background())
}

func (i BoolEqualsAdvancedFilterArgs) ToBoolEqualsAdvancedFilterOutputWithContext(ctx context.Context) BoolEqualsAdvancedFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoolEqualsAdvancedFilterOutput)
}

type BoolEqualsAdvancedFilterOutput struct{ *pulumi.OutputState }

func (BoolEqualsAdvancedFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BoolEqualsAdvancedFilter)(nil)).Elem()
}

func (o BoolEqualsAdvancedFilterOutput) ToBoolEqualsAdvancedFilterOutput() BoolEqualsAdvancedFilterOutput {
	return o
}

func (o BoolEqualsAdvancedFilterOutput) ToBoolEqualsAdvancedFilterOutputWithContext(ctx context.Context) BoolEqualsAdvancedFilterOutput {
	return o
}

func (o BoolEqualsAdvancedFilterOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BoolEqualsAdvancedFilter) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o BoolEqualsAdvancedFilterOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v BoolEqualsAdvancedFilter) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o BoolEqualsAdvancedFilterOutput) Value() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BoolEqualsAdvancedFilter) *bool { return v.Value }).(pulumi.BoolPtrOutput)
}

type BoolEqualsAdvancedFilterResponse struct {
	Key          *string `pulumi:"key"`
	OperatorType string  `pulumi:"operatorType"`
	Value        *bool   `pulumi:"value"`
}





type BoolEqualsAdvancedFilterResponseInput interface {
	pulumi.Input

	ToBoolEqualsAdvancedFilterResponseOutput() BoolEqualsAdvancedFilterResponseOutput
	ToBoolEqualsAdvancedFilterResponseOutputWithContext(context.Context) BoolEqualsAdvancedFilterResponseOutput
}

type BoolEqualsAdvancedFilterResponseArgs struct {
	Key          pulumi.StringPtrInput `pulumi:"key"`
	OperatorType pulumi.StringInput    `pulumi:"operatorType"`
	Value        pulumi.BoolPtrInput   `pulumi:"value"`
}

func (BoolEqualsAdvancedFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BoolEqualsAdvancedFilterResponse)(nil)).Elem()
}

func (i BoolEqualsAdvancedFilterResponseArgs) ToBoolEqualsAdvancedFilterResponseOutput() BoolEqualsAdvancedFilterResponseOutput {
	return i.ToBoolEqualsAdvancedFilterResponseOutputWithContext(context.Background())
}

func (i BoolEqualsAdvancedFilterResponseArgs) ToBoolEqualsAdvancedFilterResponseOutputWithContext(ctx context.Context) BoolEqualsAdvancedFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoolEqualsAdvancedFilterResponseOutput)
}

type BoolEqualsAdvancedFilterResponseOutput struct{ *pulumi.OutputState }

func (BoolEqualsAdvancedFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BoolEqualsAdvancedFilterResponse)(nil)).Elem()
}

func (o BoolEqualsAdvancedFilterResponseOutput) ToBoolEqualsAdvancedFilterResponseOutput() BoolEqualsAdvancedFilterResponseOutput {
	return o
}

func (o BoolEqualsAdvancedFilterResponseOutput) ToBoolEqualsAdvancedFilterResponseOutputWithContext(ctx context.Context) BoolEqualsAdvancedFilterResponseOutput {
	return o
}

func (o BoolEqualsAdvancedFilterResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BoolEqualsAdvancedFilterResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o BoolEqualsAdvancedFilterResponseOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v BoolEqualsAdvancedFilterResponse) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o BoolEqualsAdvancedFilterResponseOutput) Value() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BoolEqualsAdvancedFilterResponse) *bool { return v.Value }).(pulumi.BoolPtrOutput)
}

type ConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type ConnectionStateInput interface {
	pulumi.Input

	ToConnectionStateOutput() ConnectionStateOutput
	ToConnectionStateOutputWithContext(context.Context) ConnectionStateOutput
}

type ConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (ConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionState)(nil)).Elem()
}

func (i ConnectionStateArgs) ToConnectionStateOutput() ConnectionStateOutput {
	return i.ToConnectionStateOutputWithContext(context.Background())
}

func (i ConnectionStateArgs) ToConnectionStateOutputWithContext(ctx context.Context) ConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateOutput)
}

func (i ConnectionStateArgs) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return i.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (i ConnectionStateArgs) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateOutput).ToConnectionStatePtrOutputWithContext(ctx)
}









type ConnectionStatePtrInput interface {
	pulumi.Input

	ToConnectionStatePtrOutput() ConnectionStatePtrOutput
	ToConnectionStatePtrOutputWithContext(context.Context) ConnectionStatePtrOutput
}

type connectionStatePtrType ConnectionStateArgs

func ConnectionStatePtr(v *ConnectionStateArgs) ConnectionStatePtrInput {
	return (*connectionStatePtrType)(v)
}

func (*connectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionState)(nil)).Elem()
}

func (i *connectionStatePtrType) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return i.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (i *connectionStatePtrType) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStatePtrOutput)
}

type ConnectionStateOutput struct{ *pulumi.OutputState }

func (ConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionState)(nil)).Elem()
}

func (o ConnectionStateOutput) ToConnectionStateOutput() ConnectionStateOutput {
	return o
}

func (o ConnectionStateOutput) ToConnectionStateOutputWithContext(ctx context.Context) ConnectionStateOutput {
	return o
}

func (o ConnectionStateOutput) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return o.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (o ConnectionStateOutput) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionState) *ConnectionState {
		return &v
	}).(ConnectionStatePtrOutput)
}

func (o ConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o ConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (ConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionState)(nil)).Elem()
}

func (o ConnectionStatePtrOutput) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return o
}

func (o ConnectionStatePtrOutput) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return o
}

func (o ConnectionStatePtrOutput) Elem() ConnectionStateOutput {
	return o.ApplyT(func(v *ConnectionState) ConnectionState {
		if v != nil {
			return *v
		}
		var ret ConnectionState
		return ret
	}).(ConnectionStateOutput)
}

func (o ConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type ConnectionStateResponseInput interface {
	pulumi.Input

	ToConnectionStateResponseOutput() ConnectionStateResponseOutput
	ToConnectionStateResponseOutputWithContext(context.Context) ConnectionStateResponseOutput
}

type ConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (ConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStateResponse)(nil)).Elem()
}

func (i ConnectionStateResponseArgs) ToConnectionStateResponseOutput() ConnectionStateResponseOutput {
	return i.ToConnectionStateResponseOutputWithContext(context.Background())
}

func (i ConnectionStateResponseArgs) ToConnectionStateResponseOutputWithContext(ctx context.Context) ConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateResponseOutput)
}

func (i ConnectionStateResponseArgs) ToConnectionStateResponsePtrOutput() ConnectionStateResponsePtrOutput {
	return i.ToConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i ConnectionStateResponseArgs) ToConnectionStateResponsePtrOutputWithContext(ctx context.Context) ConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateResponseOutput).ToConnectionStateResponsePtrOutputWithContext(ctx)
}









type ConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToConnectionStateResponsePtrOutput() ConnectionStateResponsePtrOutput
	ToConnectionStateResponsePtrOutputWithContext(context.Context) ConnectionStateResponsePtrOutput
}

type connectionStateResponsePtrType ConnectionStateResponseArgs

func ConnectionStateResponsePtr(v *ConnectionStateResponseArgs) ConnectionStateResponsePtrInput {
	return (*connectionStateResponsePtrType)(v)
}

func (*connectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionStateResponse)(nil)).Elem()
}

func (i *connectionStateResponsePtrType) ToConnectionStateResponsePtrOutput() ConnectionStateResponsePtrOutput {
	return i.ToConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *connectionStateResponsePtrType) ToConnectionStateResponsePtrOutputWithContext(ctx context.Context) ConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateResponsePtrOutput)
}

type ConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (ConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStateResponse)(nil)).Elem()
}

func (o ConnectionStateResponseOutput) ToConnectionStateResponseOutput() ConnectionStateResponseOutput {
	return o
}

func (o ConnectionStateResponseOutput) ToConnectionStateResponseOutputWithContext(ctx context.Context) ConnectionStateResponseOutput {
	return o
}

func (o ConnectionStateResponseOutput) ToConnectionStateResponsePtrOutput() ConnectionStateResponsePtrOutput {
	return o.ToConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o ConnectionStateResponseOutput) ToConnectionStateResponsePtrOutputWithContext(ctx context.Context) ConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionStateResponse) *ConnectionStateResponse {
		return &v
	}).(ConnectionStateResponsePtrOutput)
}

func (o ConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o ConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionStateResponse)(nil)).Elem()
}

func (o ConnectionStateResponsePtrOutput) ToConnectionStateResponsePtrOutput() ConnectionStateResponsePtrOutput {
	return o
}

func (o ConnectionStateResponsePtrOutput) ToConnectionStateResponsePtrOutputWithContext(ctx context.Context) ConnectionStateResponsePtrOutput {
	return o
}

func (o ConnectionStateResponsePtrOutput) Elem() ConnectionStateResponseOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) ConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret ConnectionStateResponse
		return ret
	}).(ConnectionStateResponseOutput)
}

func (o ConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type DeadLetterWithResourceIdentity struct {
	DeadLetterDestination *StorageBlobDeadLetterDestination `pulumi:"deadLetterDestination"`
	Identity              *EventSubscriptionIdentity        `pulumi:"identity"`
}





type DeadLetterWithResourceIdentityInput interface {
	pulumi.Input

	ToDeadLetterWithResourceIdentityOutput() DeadLetterWithResourceIdentityOutput
	ToDeadLetterWithResourceIdentityOutputWithContext(context.Context) DeadLetterWithResourceIdentityOutput
}

type DeadLetterWithResourceIdentityArgs struct {
	DeadLetterDestination StorageBlobDeadLetterDestinationPtrInput `pulumi:"deadLetterDestination"`
	Identity              EventSubscriptionIdentityPtrInput        `pulumi:"identity"`
}

func (DeadLetterWithResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeadLetterWithResourceIdentity)(nil)).Elem()
}

func (i DeadLetterWithResourceIdentityArgs) ToDeadLetterWithResourceIdentityOutput() DeadLetterWithResourceIdentityOutput {
	return i.ToDeadLetterWithResourceIdentityOutputWithContext(context.Background())
}

func (i DeadLetterWithResourceIdentityArgs) ToDeadLetterWithResourceIdentityOutputWithContext(ctx context.Context) DeadLetterWithResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeadLetterWithResourceIdentityOutput)
}

func (i DeadLetterWithResourceIdentityArgs) ToDeadLetterWithResourceIdentityPtrOutput() DeadLetterWithResourceIdentityPtrOutput {
	return i.ToDeadLetterWithResourceIdentityPtrOutputWithContext(context.Background())
}

func (i DeadLetterWithResourceIdentityArgs) ToDeadLetterWithResourceIdentityPtrOutputWithContext(ctx context.Context) DeadLetterWithResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeadLetterWithResourceIdentityOutput).ToDeadLetterWithResourceIdentityPtrOutputWithContext(ctx)
}









type DeadLetterWithResourceIdentityPtrInput interface {
	pulumi.Input

	ToDeadLetterWithResourceIdentityPtrOutput() DeadLetterWithResourceIdentityPtrOutput
	ToDeadLetterWithResourceIdentityPtrOutputWithContext(context.Context) DeadLetterWithResourceIdentityPtrOutput
}

type deadLetterWithResourceIdentityPtrType DeadLetterWithResourceIdentityArgs

func DeadLetterWithResourceIdentityPtr(v *DeadLetterWithResourceIdentityArgs) DeadLetterWithResourceIdentityPtrInput {
	return (*deadLetterWithResourceIdentityPtrType)(v)
}

func (*deadLetterWithResourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeadLetterWithResourceIdentity)(nil)).Elem()
}

func (i *deadLetterWithResourceIdentityPtrType) ToDeadLetterWithResourceIdentityPtrOutput() DeadLetterWithResourceIdentityPtrOutput {
	return i.ToDeadLetterWithResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *deadLetterWithResourceIdentityPtrType) ToDeadLetterWithResourceIdentityPtrOutputWithContext(ctx context.Context) DeadLetterWithResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeadLetterWithResourceIdentityPtrOutput)
}

type DeadLetterWithResourceIdentityOutput struct{ *pulumi.OutputState }

func (DeadLetterWithResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeadLetterWithResourceIdentity)(nil)).Elem()
}

func (o DeadLetterWithResourceIdentityOutput) ToDeadLetterWithResourceIdentityOutput() DeadLetterWithResourceIdentityOutput {
	return o
}

func (o DeadLetterWithResourceIdentityOutput) ToDeadLetterWithResourceIdentityOutputWithContext(ctx context.Context) DeadLetterWithResourceIdentityOutput {
	return o
}

func (o DeadLetterWithResourceIdentityOutput) ToDeadLetterWithResourceIdentityPtrOutput() DeadLetterWithResourceIdentityPtrOutput {
	return o.ToDeadLetterWithResourceIdentityPtrOutputWithContext(context.Background())
}

func (o DeadLetterWithResourceIdentityOutput) ToDeadLetterWithResourceIdentityPtrOutputWithContext(ctx context.Context) DeadLetterWithResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeadLetterWithResourceIdentity) *DeadLetterWithResourceIdentity {
		return &v
	}).(DeadLetterWithResourceIdentityPtrOutput)
}

func (o DeadLetterWithResourceIdentityOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationPtrOutput {
	return o.ApplyT(func(v DeadLetterWithResourceIdentity) *StorageBlobDeadLetterDestination {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationPtrOutput)
}

func (o DeadLetterWithResourceIdentityOutput) Identity() EventSubscriptionIdentityPtrOutput {
	return o.ApplyT(func(v DeadLetterWithResourceIdentity) *EventSubscriptionIdentity { return v.Identity }).(EventSubscriptionIdentityPtrOutput)
}

type DeadLetterWithResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (DeadLetterWithResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeadLetterWithResourceIdentity)(nil)).Elem()
}

func (o DeadLetterWithResourceIdentityPtrOutput) ToDeadLetterWithResourceIdentityPtrOutput() DeadLetterWithResourceIdentityPtrOutput {
	return o
}

func (o DeadLetterWithResourceIdentityPtrOutput) ToDeadLetterWithResourceIdentityPtrOutputWithContext(ctx context.Context) DeadLetterWithResourceIdentityPtrOutput {
	return o
}

func (o DeadLetterWithResourceIdentityPtrOutput) Elem() DeadLetterWithResourceIdentityOutput {
	return o.ApplyT(func(v *DeadLetterWithResourceIdentity) DeadLetterWithResourceIdentity {
		if v != nil {
			return *v
		}
		var ret DeadLetterWithResourceIdentity
		return ret
	}).(DeadLetterWithResourceIdentityOutput)
}

func (o DeadLetterWithResourceIdentityPtrOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationPtrOutput {
	return o.ApplyT(func(v *DeadLetterWithResourceIdentity) *StorageBlobDeadLetterDestination {
		if v == nil {
			return nil
		}
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationPtrOutput)
}

func (o DeadLetterWithResourceIdentityPtrOutput) Identity() EventSubscriptionIdentityPtrOutput {
	return o.ApplyT(func(v *DeadLetterWithResourceIdentity) *EventSubscriptionIdentity {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(EventSubscriptionIdentityPtrOutput)
}

type DeadLetterWithResourceIdentityResponse struct {
	DeadLetterDestination *StorageBlobDeadLetterDestinationResponse `pulumi:"deadLetterDestination"`
	Identity              *EventSubscriptionIdentityResponse        `pulumi:"identity"`
}





type DeadLetterWithResourceIdentityResponseInput interface {
	pulumi.Input

	ToDeadLetterWithResourceIdentityResponseOutput() DeadLetterWithResourceIdentityResponseOutput
	ToDeadLetterWithResourceIdentityResponseOutputWithContext(context.Context) DeadLetterWithResourceIdentityResponseOutput
}

type DeadLetterWithResourceIdentityResponseArgs struct {
	DeadLetterDestination StorageBlobDeadLetterDestinationResponsePtrInput `pulumi:"deadLetterDestination"`
	Identity              EventSubscriptionIdentityResponsePtrInput        `pulumi:"identity"`
}

func (DeadLetterWithResourceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeadLetterWithResourceIdentityResponse)(nil)).Elem()
}

func (i DeadLetterWithResourceIdentityResponseArgs) ToDeadLetterWithResourceIdentityResponseOutput() DeadLetterWithResourceIdentityResponseOutput {
	return i.ToDeadLetterWithResourceIdentityResponseOutputWithContext(context.Background())
}

func (i DeadLetterWithResourceIdentityResponseArgs) ToDeadLetterWithResourceIdentityResponseOutputWithContext(ctx context.Context) DeadLetterWithResourceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeadLetterWithResourceIdentityResponseOutput)
}

func (i DeadLetterWithResourceIdentityResponseArgs) ToDeadLetterWithResourceIdentityResponsePtrOutput() DeadLetterWithResourceIdentityResponsePtrOutput {
	return i.ToDeadLetterWithResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i DeadLetterWithResourceIdentityResponseArgs) ToDeadLetterWithResourceIdentityResponsePtrOutputWithContext(ctx context.Context) DeadLetterWithResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeadLetterWithResourceIdentityResponseOutput).ToDeadLetterWithResourceIdentityResponsePtrOutputWithContext(ctx)
}









type DeadLetterWithResourceIdentityResponsePtrInput interface {
	pulumi.Input

	ToDeadLetterWithResourceIdentityResponsePtrOutput() DeadLetterWithResourceIdentityResponsePtrOutput
	ToDeadLetterWithResourceIdentityResponsePtrOutputWithContext(context.Context) DeadLetterWithResourceIdentityResponsePtrOutput
}

type deadLetterWithResourceIdentityResponsePtrType DeadLetterWithResourceIdentityResponseArgs

func DeadLetterWithResourceIdentityResponsePtr(v *DeadLetterWithResourceIdentityResponseArgs) DeadLetterWithResourceIdentityResponsePtrInput {
	return (*deadLetterWithResourceIdentityResponsePtrType)(v)
}

func (*deadLetterWithResourceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeadLetterWithResourceIdentityResponse)(nil)).Elem()
}

func (i *deadLetterWithResourceIdentityResponsePtrType) ToDeadLetterWithResourceIdentityResponsePtrOutput() DeadLetterWithResourceIdentityResponsePtrOutput {
	return i.ToDeadLetterWithResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *deadLetterWithResourceIdentityResponsePtrType) ToDeadLetterWithResourceIdentityResponsePtrOutputWithContext(ctx context.Context) DeadLetterWithResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

type DeadLetterWithResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (DeadLetterWithResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeadLetterWithResourceIdentityResponse)(nil)).Elem()
}

func (o DeadLetterWithResourceIdentityResponseOutput) ToDeadLetterWithResourceIdentityResponseOutput() DeadLetterWithResourceIdentityResponseOutput {
	return o
}

func (o DeadLetterWithResourceIdentityResponseOutput) ToDeadLetterWithResourceIdentityResponseOutputWithContext(ctx context.Context) DeadLetterWithResourceIdentityResponseOutput {
	return o
}

func (o DeadLetterWithResourceIdentityResponseOutput) ToDeadLetterWithResourceIdentityResponsePtrOutput() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ToDeadLetterWithResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o DeadLetterWithResourceIdentityResponseOutput) ToDeadLetterWithResourceIdentityResponsePtrOutputWithContext(ctx context.Context) DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeadLetterWithResourceIdentityResponse) *DeadLetterWithResourceIdentityResponse {
		return &v
	}).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

func (o DeadLetterWithResourceIdentityResponseOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v DeadLetterWithResourceIdentityResponse) *StorageBlobDeadLetterDestinationResponse {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

func (o DeadLetterWithResourceIdentityResponseOutput) Identity() EventSubscriptionIdentityResponsePtrOutput {
	return o.ApplyT(func(v DeadLetterWithResourceIdentityResponse) *EventSubscriptionIdentityResponse { return v.Identity }).(EventSubscriptionIdentityResponsePtrOutput)
}

type DeadLetterWithResourceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (DeadLetterWithResourceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeadLetterWithResourceIdentityResponse)(nil)).Elem()
}

func (o DeadLetterWithResourceIdentityResponsePtrOutput) ToDeadLetterWithResourceIdentityResponsePtrOutput() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o
}

func (o DeadLetterWithResourceIdentityResponsePtrOutput) ToDeadLetterWithResourceIdentityResponsePtrOutputWithContext(ctx context.Context) DeadLetterWithResourceIdentityResponsePtrOutput {
	return o
}

func (o DeadLetterWithResourceIdentityResponsePtrOutput) Elem() DeadLetterWithResourceIdentityResponseOutput {
	return o.ApplyT(func(v *DeadLetterWithResourceIdentityResponse) DeadLetterWithResourceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret DeadLetterWithResourceIdentityResponse
		return ret
	}).(DeadLetterWithResourceIdentityResponseOutput)
}

func (o DeadLetterWithResourceIdentityResponsePtrOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v *DeadLetterWithResourceIdentityResponse) *StorageBlobDeadLetterDestinationResponse {
		if v == nil {
			return nil
		}
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

func (o DeadLetterWithResourceIdentityResponsePtrOutput) Identity() EventSubscriptionIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DeadLetterWithResourceIdentityResponse) *EventSubscriptionIdentityResponse {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(EventSubscriptionIdentityResponsePtrOutput)
}

type DeliveryWithResourceIdentity struct {
	Destination interface{}                `pulumi:"destination"`
	Identity    *EventSubscriptionIdentity `pulumi:"identity"`
}





type DeliveryWithResourceIdentityInput interface {
	pulumi.Input

	ToDeliveryWithResourceIdentityOutput() DeliveryWithResourceIdentityOutput
	ToDeliveryWithResourceIdentityOutputWithContext(context.Context) DeliveryWithResourceIdentityOutput
}

type DeliveryWithResourceIdentityArgs struct {
	Destination pulumi.Input                      `pulumi:"destination"`
	Identity    EventSubscriptionIdentityPtrInput `pulumi:"identity"`
}

func (DeliveryWithResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryWithResourceIdentity)(nil)).Elem()
}

func (i DeliveryWithResourceIdentityArgs) ToDeliveryWithResourceIdentityOutput() DeliveryWithResourceIdentityOutput {
	return i.ToDeliveryWithResourceIdentityOutputWithContext(context.Background())
}

func (i DeliveryWithResourceIdentityArgs) ToDeliveryWithResourceIdentityOutputWithContext(ctx context.Context) DeliveryWithResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryWithResourceIdentityOutput)
}

func (i DeliveryWithResourceIdentityArgs) ToDeliveryWithResourceIdentityPtrOutput() DeliveryWithResourceIdentityPtrOutput {
	return i.ToDeliveryWithResourceIdentityPtrOutputWithContext(context.Background())
}

func (i DeliveryWithResourceIdentityArgs) ToDeliveryWithResourceIdentityPtrOutputWithContext(ctx context.Context) DeliveryWithResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryWithResourceIdentityOutput).ToDeliveryWithResourceIdentityPtrOutputWithContext(ctx)
}









type DeliveryWithResourceIdentityPtrInput interface {
	pulumi.Input

	ToDeliveryWithResourceIdentityPtrOutput() DeliveryWithResourceIdentityPtrOutput
	ToDeliveryWithResourceIdentityPtrOutputWithContext(context.Context) DeliveryWithResourceIdentityPtrOutput
}

type deliveryWithResourceIdentityPtrType DeliveryWithResourceIdentityArgs

func DeliveryWithResourceIdentityPtr(v *DeliveryWithResourceIdentityArgs) DeliveryWithResourceIdentityPtrInput {
	return (*deliveryWithResourceIdentityPtrType)(v)
}

func (*deliveryWithResourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryWithResourceIdentity)(nil)).Elem()
}

func (i *deliveryWithResourceIdentityPtrType) ToDeliveryWithResourceIdentityPtrOutput() DeliveryWithResourceIdentityPtrOutput {
	return i.ToDeliveryWithResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *deliveryWithResourceIdentityPtrType) ToDeliveryWithResourceIdentityPtrOutputWithContext(ctx context.Context) DeliveryWithResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryWithResourceIdentityPtrOutput)
}

type DeliveryWithResourceIdentityOutput struct{ *pulumi.OutputState }

func (DeliveryWithResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryWithResourceIdentity)(nil)).Elem()
}

func (o DeliveryWithResourceIdentityOutput) ToDeliveryWithResourceIdentityOutput() DeliveryWithResourceIdentityOutput {
	return o
}

func (o DeliveryWithResourceIdentityOutput) ToDeliveryWithResourceIdentityOutputWithContext(ctx context.Context) DeliveryWithResourceIdentityOutput {
	return o
}

func (o DeliveryWithResourceIdentityOutput) ToDeliveryWithResourceIdentityPtrOutput() DeliveryWithResourceIdentityPtrOutput {
	return o.ToDeliveryWithResourceIdentityPtrOutputWithContext(context.Background())
}

func (o DeliveryWithResourceIdentityOutput) ToDeliveryWithResourceIdentityPtrOutputWithContext(ctx context.Context) DeliveryWithResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeliveryWithResourceIdentity) *DeliveryWithResourceIdentity {
		return &v
	}).(DeliveryWithResourceIdentityPtrOutput)
}

func (o DeliveryWithResourceIdentityOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v DeliveryWithResourceIdentity) interface{} { return v.Destination }).(pulumi.AnyOutput)
}

func (o DeliveryWithResourceIdentityOutput) Identity() EventSubscriptionIdentityPtrOutput {
	return o.ApplyT(func(v DeliveryWithResourceIdentity) *EventSubscriptionIdentity { return v.Identity }).(EventSubscriptionIdentityPtrOutput)
}

type DeliveryWithResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (DeliveryWithResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryWithResourceIdentity)(nil)).Elem()
}

func (o DeliveryWithResourceIdentityPtrOutput) ToDeliveryWithResourceIdentityPtrOutput() DeliveryWithResourceIdentityPtrOutput {
	return o
}

func (o DeliveryWithResourceIdentityPtrOutput) ToDeliveryWithResourceIdentityPtrOutputWithContext(ctx context.Context) DeliveryWithResourceIdentityPtrOutput {
	return o
}

func (o DeliveryWithResourceIdentityPtrOutput) Elem() DeliveryWithResourceIdentityOutput {
	return o.ApplyT(func(v *DeliveryWithResourceIdentity) DeliveryWithResourceIdentity {
		if v != nil {
			return *v
		}
		var ret DeliveryWithResourceIdentity
		return ret
	}).(DeliveryWithResourceIdentityOutput)
}

func (o DeliveryWithResourceIdentityPtrOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v *DeliveryWithResourceIdentity) interface{} {
		if v == nil {
			return nil
		}
		return v.Destination
	}).(pulumi.AnyOutput)
}

func (o DeliveryWithResourceIdentityPtrOutput) Identity() EventSubscriptionIdentityPtrOutput {
	return o.ApplyT(func(v *DeliveryWithResourceIdentity) *EventSubscriptionIdentity {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(EventSubscriptionIdentityPtrOutput)
}

type DeliveryWithResourceIdentityResponse struct {
	Destination interface{}                        `pulumi:"destination"`
	Identity    *EventSubscriptionIdentityResponse `pulumi:"identity"`
}





type DeliveryWithResourceIdentityResponseInput interface {
	pulumi.Input

	ToDeliveryWithResourceIdentityResponseOutput() DeliveryWithResourceIdentityResponseOutput
	ToDeliveryWithResourceIdentityResponseOutputWithContext(context.Context) DeliveryWithResourceIdentityResponseOutput
}

type DeliveryWithResourceIdentityResponseArgs struct {
	Destination pulumi.Input                              `pulumi:"destination"`
	Identity    EventSubscriptionIdentityResponsePtrInput `pulumi:"identity"`
}

func (DeliveryWithResourceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryWithResourceIdentityResponse)(nil)).Elem()
}

func (i DeliveryWithResourceIdentityResponseArgs) ToDeliveryWithResourceIdentityResponseOutput() DeliveryWithResourceIdentityResponseOutput {
	return i.ToDeliveryWithResourceIdentityResponseOutputWithContext(context.Background())
}

func (i DeliveryWithResourceIdentityResponseArgs) ToDeliveryWithResourceIdentityResponseOutputWithContext(ctx context.Context) DeliveryWithResourceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryWithResourceIdentityResponseOutput)
}

func (i DeliveryWithResourceIdentityResponseArgs) ToDeliveryWithResourceIdentityResponsePtrOutput() DeliveryWithResourceIdentityResponsePtrOutput {
	return i.ToDeliveryWithResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i DeliveryWithResourceIdentityResponseArgs) ToDeliveryWithResourceIdentityResponsePtrOutputWithContext(ctx context.Context) DeliveryWithResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryWithResourceIdentityResponseOutput).ToDeliveryWithResourceIdentityResponsePtrOutputWithContext(ctx)
}









type DeliveryWithResourceIdentityResponsePtrInput interface {
	pulumi.Input

	ToDeliveryWithResourceIdentityResponsePtrOutput() DeliveryWithResourceIdentityResponsePtrOutput
	ToDeliveryWithResourceIdentityResponsePtrOutputWithContext(context.Context) DeliveryWithResourceIdentityResponsePtrOutput
}

type deliveryWithResourceIdentityResponsePtrType DeliveryWithResourceIdentityResponseArgs

func DeliveryWithResourceIdentityResponsePtr(v *DeliveryWithResourceIdentityResponseArgs) DeliveryWithResourceIdentityResponsePtrInput {
	return (*deliveryWithResourceIdentityResponsePtrType)(v)
}

func (*deliveryWithResourceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryWithResourceIdentityResponse)(nil)).Elem()
}

func (i *deliveryWithResourceIdentityResponsePtrType) ToDeliveryWithResourceIdentityResponsePtrOutput() DeliveryWithResourceIdentityResponsePtrOutput {
	return i.ToDeliveryWithResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *deliveryWithResourceIdentityResponsePtrType) ToDeliveryWithResourceIdentityResponsePtrOutputWithContext(ctx context.Context) DeliveryWithResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryWithResourceIdentityResponsePtrOutput)
}

type DeliveryWithResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (DeliveryWithResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryWithResourceIdentityResponse)(nil)).Elem()
}

func (o DeliveryWithResourceIdentityResponseOutput) ToDeliveryWithResourceIdentityResponseOutput() DeliveryWithResourceIdentityResponseOutput {
	return o
}

func (o DeliveryWithResourceIdentityResponseOutput) ToDeliveryWithResourceIdentityResponseOutputWithContext(ctx context.Context) DeliveryWithResourceIdentityResponseOutput {
	return o
}

func (o DeliveryWithResourceIdentityResponseOutput) ToDeliveryWithResourceIdentityResponsePtrOutput() DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ToDeliveryWithResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o DeliveryWithResourceIdentityResponseOutput) ToDeliveryWithResourceIdentityResponsePtrOutputWithContext(ctx context.Context) DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeliveryWithResourceIdentityResponse) *DeliveryWithResourceIdentityResponse {
		return &v
	}).(DeliveryWithResourceIdentityResponsePtrOutput)
}

func (o DeliveryWithResourceIdentityResponseOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v DeliveryWithResourceIdentityResponse) interface{} { return v.Destination }).(pulumi.AnyOutput)
}

func (o DeliveryWithResourceIdentityResponseOutput) Identity() EventSubscriptionIdentityResponsePtrOutput {
	return o.ApplyT(func(v DeliveryWithResourceIdentityResponse) *EventSubscriptionIdentityResponse { return v.Identity }).(EventSubscriptionIdentityResponsePtrOutput)
}

type DeliveryWithResourceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (DeliveryWithResourceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryWithResourceIdentityResponse)(nil)).Elem()
}

func (o DeliveryWithResourceIdentityResponsePtrOutput) ToDeliveryWithResourceIdentityResponsePtrOutput() DeliveryWithResourceIdentityResponsePtrOutput {
	return o
}

func (o DeliveryWithResourceIdentityResponsePtrOutput) ToDeliveryWithResourceIdentityResponsePtrOutputWithContext(ctx context.Context) DeliveryWithResourceIdentityResponsePtrOutput {
	return o
}

func (o DeliveryWithResourceIdentityResponsePtrOutput) Elem() DeliveryWithResourceIdentityResponseOutput {
	return o.ApplyT(func(v *DeliveryWithResourceIdentityResponse) DeliveryWithResourceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret DeliveryWithResourceIdentityResponse
		return ret
	}).(DeliveryWithResourceIdentityResponseOutput)
}

func (o DeliveryWithResourceIdentityResponsePtrOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v *DeliveryWithResourceIdentityResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Destination
	}).(pulumi.AnyOutput)
}

func (o DeliveryWithResourceIdentityResponsePtrOutput) Identity() EventSubscriptionIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DeliveryWithResourceIdentityResponse) *EventSubscriptionIdentityResponse {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(EventSubscriptionIdentityResponsePtrOutput)
}

type DynamicDeliveryAttributeMapping struct {
	Name        *string `pulumi:"name"`
	SourceField *string `pulumi:"sourceField"`
	Type        string  `pulumi:"type"`
}





type DynamicDeliveryAttributeMappingInput interface {
	pulumi.Input

	ToDynamicDeliveryAttributeMappingOutput() DynamicDeliveryAttributeMappingOutput
	ToDynamicDeliveryAttributeMappingOutputWithContext(context.Context) DynamicDeliveryAttributeMappingOutput
}

type DynamicDeliveryAttributeMappingArgs struct {
	Name        pulumi.StringPtrInput `pulumi:"name"`
	SourceField pulumi.StringPtrInput `pulumi:"sourceField"`
	Type        pulumi.StringInput    `pulumi:"type"`
}

func (DynamicDeliveryAttributeMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicDeliveryAttributeMapping)(nil)).Elem()
}

func (i DynamicDeliveryAttributeMappingArgs) ToDynamicDeliveryAttributeMappingOutput() DynamicDeliveryAttributeMappingOutput {
	return i.ToDynamicDeliveryAttributeMappingOutputWithContext(context.Background())
}

func (i DynamicDeliveryAttributeMappingArgs) ToDynamicDeliveryAttributeMappingOutputWithContext(ctx context.Context) DynamicDeliveryAttributeMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamicDeliveryAttributeMappingOutput)
}

type DynamicDeliveryAttributeMappingOutput struct{ *pulumi.OutputState }

func (DynamicDeliveryAttributeMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicDeliveryAttributeMapping)(nil)).Elem()
}

func (o DynamicDeliveryAttributeMappingOutput) ToDynamicDeliveryAttributeMappingOutput() DynamicDeliveryAttributeMappingOutput {
	return o
}

func (o DynamicDeliveryAttributeMappingOutput) ToDynamicDeliveryAttributeMappingOutputWithContext(ctx context.Context) DynamicDeliveryAttributeMappingOutput {
	return o
}

func (o DynamicDeliveryAttributeMappingOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DynamicDeliveryAttributeMapping) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DynamicDeliveryAttributeMappingOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DynamicDeliveryAttributeMapping) *string { return v.SourceField }).(pulumi.StringPtrOutput)
}

func (o DynamicDeliveryAttributeMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DynamicDeliveryAttributeMapping) string { return v.Type }).(pulumi.StringOutput)
}

type DynamicDeliveryAttributeMappingResponse struct {
	Name        *string `pulumi:"name"`
	SourceField *string `pulumi:"sourceField"`
	Type        string  `pulumi:"type"`
}





type DynamicDeliveryAttributeMappingResponseInput interface {
	pulumi.Input

	ToDynamicDeliveryAttributeMappingResponseOutput() DynamicDeliveryAttributeMappingResponseOutput
	ToDynamicDeliveryAttributeMappingResponseOutputWithContext(context.Context) DynamicDeliveryAttributeMappingResponseOutput
}

type DynamicDeliveryAttributeMappingResponseArgs struct {
	Name        pulumi.StringPtrInput `pulumi:"name"`
	SourceField pulumi.StringPtrInput `pulumi:"sourceField"`
	Type        pulumi.StringInput    `pulumi:"type"`
}

func (DynamicDeliveryAttributeMappingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicDeliveryAttributeMappingResponse)(nil)).Elem()
}

func (i DynamicDeliveryAttributeMappingResponseArgs) ToDynamicDeliveryAttributeMappingResponseOutput() DynamicDeliveryAttributeMappingResponseOutput {
	return i.ToDynamicDeliveryAttributeMappingResponseOutputWithContext(context.Background())
}

func (i DynamicDeliveryAttributeMappingResponseArgs) ToDynamicDeliveryAttributeMappingResponseOutputWithContext(ctx context.Context) DynamicDeliveryAttributeMappingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamicDeliveryAttributeMappingResponseOutput)
}

type DynamicDeliveryAttributeMappingResponseOutput struct{ *pulumi.OutputState }

func (DynamicDeliveryAttributeMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicDeliveryAttributeMappingResponse)(nil)).Elem()
}

func (o DynamicDeliveryAttributeMappingResponseOutput) ToDynamicDeliveryAttributeMappingResponseOutput() DynamicDeliveryAttributeMappingResponseOutput {
	return o
}

func (o DynamicDeliveryAttributeMappingResponseOutput) ToDynamicDeliveryAttributeMappingResponseOutputWithContext(ctx context.Context) DynamicDeliveryAttributeMappingResponseOutput {
	return o
}

func (o DynamicDeliveryAttributeMappingResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DynamicDeliveryAttributeMappingResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DynamicDeliveryAttributeMappingResponseOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DynamicDeliveryAttributeMappingResponse) *string { return v.SourceField }).(pulumi.StringPtrOutput)
}

func (o DynamicDeliveryAttributeMappingResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DynamicDeliveryAttributeMappingResponse) string { return v.Type }).(pulumi.StringOutput)
}

type EventHubEventSubscriptionDestination struct {
	DeliveryAttributeMappings []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType              string        `pulumi:"endpointType"`
	ResourceId                *string       `pulumi:"resourceId"`
}





type EventHubEventSubscriptionDestinationInput interface {
	pulumi.Input

	ToEventHubEventSubscriptionDestinationOutput() EventHubEventSubscriptionDestinationOutput
	ToEventHubEventSubscriptionDestinationOutputWithContext(context.Context) EventHubEventSubscriptionDestinationOutput
}

type EventHubEventSubscriptionDestinationArgs struct {
	DeliveryAttributeMappings pulumi.ArrayInput     `pulumi:"deliveryAttributeMappings"`
	EndpointType              pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId                pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (EventHubEventSubscriptionDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubEventSubscriptionDestination)(nil)).Elem()
}

func (i EventHubEventSubscriptionDestinationArgs) ToEventHubEventSubscriptionDestinationOutput() EventHubEventSubscriptionDestinationOutput {
	return i.ToEventHubEventSubscriptionDestinationOutputWithContext(context.Background())
}

func (i EventHubEventSubscriptionDestinationArgs) ToEventHubEventSubscriptionDestinationOutputWithContext(ctx context.Context) EventHubEventSubscriptionDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubEventSubscriptionDestinationOutput)
}

type EventHubEventSubscriptionDestinationOutput struct{ *pulumi.OutputState }

func (EventHubEventSubscriptionDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubEventSubscriptionDestination)(nil)).Elem()
}

func (o EventHubEventSubscriptionDestinationOutput) ToEventHubEventSubscriptionDestinationOutput() EventHubEventSubscriptionDestinationOutput {
	return o
}

func (o EventHubEventSubscriptionDestinationOutput) ToEventHubEventSubscriptionDestinationOutputWithContext(ctx context.Context) EventHubEventSubscriptionDestinationOutput {
	return o
}

func (o EventHubEventSubscriptionDestinationOutput) DeliveryAttributeMappings() pulumi.ArrayOutput {
	return o.ApplyT(func(v EventHubEventSubscriptionDestination) []interface{} { return v.DeliveryAttributeMappings }).(pulumi.ArrayOutput)
}

func (o EventHubEventSubscriptionDestinationOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubEventSubscriptionDestination) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o EventHubEventSubscriptionDestinationOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubEventSubscriptionDestination) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type EventHubEventSubscriptionDestinationResponse struct {
	DeliveryAttributeMappings []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType              string        `pulumi:"endpointType"`
	ResourceId                *string       `pulumi:"resourceId"`
}





type EventHubEventSubscriptionDestinationResponseInput interface {
	pulumi.Input

	ToEventHubEventSubscriptionDestinationResponseOutput() EventHubEventSubscriptionDestinationResponseOutput
	ToEventHubEventSubscriptionDestinationResponseOutputWithContext(context.Context) EventHubEventSubscriptionDestinationResponseOutput
}

type EventHubEventSubscriptionDestinationResponseArgs struct {
	DeliveryAttributeMappings pulumi.ArrayInput     `pulumi:"deliveryAttributeMappings"`
	EndpointType              pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId                pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (EventHubEventSubscriptionDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (i EventHubEventSubscriptionDestinationResponseArgs) ToEventHubEventSubscriptionDestinationResponseOutput() EventHubEventSubscriptionDestinationResponseOutput {
	return i.ToEventHubEventSubscriptionDestinationResponseOutputWithContext(context.Background())
}

func (i EventHubEventSubscriptionDestinationResponseArgs) ToEventHubEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) EventHubEventSubscriptionDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubEventSubscriptionDestinationResponseOutput)
}

type EventHubEventSubscriptionDestinationResponseOutput struct{ *pulumi.OutputState }

func (EventHubEventSubscriptionDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (o EventHubEventSubscriptionDestinationResponseOutput) ToEventHubEventSubscriptionDestinationResponseOutput() EventHubEventSubscriptionDestinationResponseOutput {
	return o
}

func (o EventHubEventSubscriptionDestinationResponseOutput) ToEventHubEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) EventHubEventSubscriptionDestinationResponseOutput {
	return o
}

func (o EventHubEventSubscriptionDestinationResponseOutput) DeliveryAttributeMappings() pulumi.ArrayOutput {
	return o.ApplyT(func(v EventHubEventSubscriptionDestinationResponse) []interface{} { return v.DeliveryAttributeMappings }).(pulumi.ArrayOutput)
}

func (o EventHubEventSubscriptionDestinationResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubEventSubscriptionDestinationResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o EventHubEventSubscriptionDestinationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubEventSubscriptionDestinationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type EventSubscriptionFilter struct {
	AdvancedFilters                 []interface{} `pulumi:"advancedFilters"`
	EnableAdvancedFilteringOnArrays *bool         `pulumi:"enableAdvancedFilteringOnArrays"`
	IncludedEventTypes              []string      `pulumi:"includedEventTypes"`
	IsSubjectCaseSensitive          *bool         `pulumi:"isSubjectCaseSensitive"`
	SubjectBeginsWith               *string       `pulumi:"subjectBeginsWith"`
	SubjectEndsWith                 *string       `pulumi:"subjectEndsWith"`
}





type EventSubscriptionFilterInput interface {
	pulumi.Input

	ToEventSubscriptionFilterOutput() EventSubscriptionFilterOutput
	ToEventSubscriptionFilterOutputWithContext(context.Context) EventSubscriptionFilterOutput
}

type EventSubscriptionFilterArgs struct {
	AdvancedFilters                 pulumi.ArrayInput       `pulumi:"advancedFilters"`
	EnableAdvancedFilteringOnArrays pulumi.BoolPtrInput     `pulumi:"enableAdvancedFilteringOnArrays"`
	IncludedEventTypes              pulumi.StringArrayInput `pulumi:"includedEventTypes"`
	IsSubjectCaseSensitive          pulumi.BoolPtrInput     `pulumi:"isSubjectCaseSensitive"`
	SubjectBeginsWith               pulumi.StringPtrInput   `pulumi:"subjectBeginsWith"`
	SubjectEndsWith                 pulumi.StringPtrInput   `pulumi:"subjectEndsWith"`
}

func (EventSubscriptionFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionFilter)(nil)).Elem()
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterOutput() EventSubscriptionFilterOutput {
	return i.ToEventSubscriptionFilterOutputWithContext(context.Background())
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterOutputWithContext(ctx context.Context) EventSubscriptionFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterOutput)
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput {
	return i.ToEventSubscriptionFilterPtrOutputWithContext(context.Background())
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterPtrOutputWithContext(ctx context.Context) EventSubscriptionFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterOutput).ToEventSubscriptionFilterPtrOutputWithContext(ctx)
}









type EventSubscriptionFilterPtrInput interface {
	pulumi.Input

	ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput
	ToEventSubscriptionFilterPtrOutputWithContext(context.Context) EventSubscriptionFilterPtrOutput
}

type eventSubscriptionFilterPtrType EventSubscriptionFilterArgs

func EventSubscriptionFilterPtr(v *EventSubscriptionFilterArgs) EventSubscriptionFilterPtrInput {
	return (*eventSubscriptionFilterPtrType)(v)
}

func (*eventSubscriptionFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionFilter)(nil)).Elem()
}

func (i *eventSubscriptionFilterPtrType) ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput {
	return i.ToEventSubscriptionFilterPtrOutputWithContext(context.Background())
}

func (i *eventSubscriptionFilterPtrType) ToEventSubscriptionFilterPtrOutputWithContext(ctx context.Context) EventSubscriptionFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterPtrOutput)
}

type EventSubscriptionFilterOutput struct{ *pulumi.OutputState }

func (EventSubscriptionFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionFilter)(nil)).Elem()
}

func (o EventSubscriptionFilterOutput) ToEventSubscriptionFilterOutput() EventSubscriptionFilterOutput {
	return o
}

func (o EventSubscriptionFilterOutput) ToEventSubscriptionFilterOutputWithContext(ctx context.Context) EventSubscriptionFilterOutput {
	return o
}

func (o EventSubscriptionFilterOutput) ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput {
	return o.ToEventSubscriptionFilterPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionFilterOutput) ToEventSubscriptionFilterPtrOutputWithContext(ctx context.Context) EventSubscriptionFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSubscriptionFilter) *EventSubscriptionFilter {
		return &v
	}).(EventSubscriptionFilterPtrOutput)
}

func (o EventSubscriptionFilterOutput) AdvancedFilters() pulumi.ArrayOutput {
	return o.ApplyT(func(v EventSubscriptionFilter) []interface{} { return v.AdvancedFilters }).(pulumi.ArrayOutput)
}

func (o EventSubscriptionFilterOutput) EnableAdvancedFilteringOnArrays() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilter) *bool { return v.EnableAdvancedFilteringOnArrays }).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterOutput) IncludedEventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EventSubscriptionFilter) []string { return v.IncludedEventTypes }).(pulumi.StringArrayOutput)
}

func (o EventSubscriptionFilterOutput) IsSubjectCaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilter) *bool { return v.IsSubjectCaseSensitive }).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterOutput) SubjectBeginsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilter) *string { return v.SubjectBeginsWith }).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionFilterOutput) SubjectEndsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilter) *string { return v.SubjectEndsWith }).(pulumi.StringPtrOutput)
}

type EventSubscriptionFilterPtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionFilter)(nil)).Elem()
}

func (o EventSubscriptionFilterPtrOutput) ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput {
	return o
}

func (o EventSubscriptionFilterPtrOutput) ToEventSubscriptionFilterPtrOutputWithContext(ctx context.Context) EventSubscriptionFilterPtrOutput {
	return o
}

func (o EventSubscriptionFilterPtrOutput) Elem() EventSubscriptionFilterOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) EventSubscriptionFilter {
		if v != nil {
			return *v
		}
		var ret EventSubscriptionFilter
		return ret
	}).(EventSubscriptionFilterOutput)
}

func (o EventSubscriptionFilterPtrOutput) AdvancedFilters() pulumi.ArrayOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) []interface{} {
		if v == nil {
			return nil
		}
		return v.AdvancedFilters
	}).(pulumi.ArrayOutput)
}

func (o EventSubscriptionFilterPtrOutput) EnableAdvancedFilteringOnArrays() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAdvancedFilteringOnArrays
	}).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterPtrOutput) IncludedEventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) []string {
		if v == nil {
			return nil
		}
		return v.IncludedEventTypes
	}).(pulumi.StringArrayOutput)
}

func (o EventSubscriptionFilterPtrOutput) IsSubjectCaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) *bool {
		if v == nil {
			return nil
		}
		return v.IsSubjectCaseSensitive
	}).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterPtrOutput) SubjectBeginsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) *string {
		if v == nil {
			return nil
		}
		return v.SubjectBeginsWith
	}).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionFilterPtrOutput) SubjectEndsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) *string {
		if v == nil {
			return nil
		}
		return v.SubjectEndsWith
	}).(pulumi.StringPtrOutput)
}

type EventSubscriptionFilterResponse struct {
	AdvancedFilters                 []interface{} `pulumi:"advancedFilters"`
	EnableAdvancedFilteringOnArrays *bool         `pulumi:"enableAdvancedFilteringOnArrays"`
	IncludedEventTypes              []string      `pulumi:"includedEventTypes"`
	IsSubjectCaseSensitive          *bool         `pulumi:"isSubjectCaseSensitive"`
	SubjectBeginsWith               *string       `pulumi:"subjectBeginsWith"`
	SubjectEndsWith                 *string       `pulumi:"subjectEndsWith"`
}





type EventSubscriptionFilterResponseInput interface {
	pulumi.Input

	ToEventSubscriptionFilterResponseOutput() EventSubscriptionFilterResponseOutput
	ToEventSubscriptionFilterResponseOutputWithContext(context.Context) EventSubscriptionFilterResponseOutput
}

type EventSubscriptionFilterResponseArgs struct {
	AdvancedFilters                 pulumi.ArrayInput       `pulumi:"advancedFilters"`
	EnableAdvancedFilteringOnArrays pulumi.BoolPtrInput     `pulumi:"enableAdvancedFilteringOnArrays"`
	IncludedEventTypes              pulumi.StringArrayInput `pulumi:"includedEventTypes"`
	IsSubjectCaseSensitive          pulumi.BoolPtrInput     `pulumi:"isSubjectCaseSensitive"`
	SubjectBeginsWith               pulumi.StringPtrInput   `pulumi:"subjectBeginsWith"`
	SubjectEndsWith                 pulumi.StringPtrInput   `pulumi:"subjectEndsWith"`
}

func (EventSubscriptionFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionFilterResponse)(nil)).Elem()
}

func (i EventSubscriptionFilterResponseArgs) ToEventSubscriptionFilterResponseOutput() EventSubscriptionFilterResponseOutput {
	return i.ToEventSubscriptionFilterResponseOutputWithContext(context.Background())
}

func (i EventSubscriptionFilterResponseArgs) ToEventSubscriptionFilterResponseOutputWithContext(ctx context.Context) EventSubscriptionFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterResponseOutput)
}

func (i EventSubscriptionFilterResponseArgs) ToEventSubscriptionFilterResponsePtrOutput() EventSubscriptionFilterResponsePtrOutput {
	return i.ToEventSubscriptionFilterResponsePtrOutputWithContext(context.Background())
}

func (i EventSubscriptionFilterResponseArgs) ToEventSubscriptionFilterResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterResponseOutput).ToEventSubscriptionFilterResponsePtrOutputWithContext(ctx)
}









type EventSubscriptionFilterResponsePtrInput interface {
	pulumi.Input

	ToEventSubscriptionFilterResponsePtrOutput() EventSubscriptionFilterResponsePtrOutput
	ToEventSubscriptionFilterResponsePtrOutputWithContext(context.Context) EventSubscriptionFilterResponsePtrOutput
}

type eventSubscriptionFilterResponsePtrType EventSubscriptionFilterResponseArgs

func EventSubscriptionFilterResponsePtr(v *EventSubscriptionFilterResponseArgs) EventSubscriptionFilterResponsePtrInput {
	return (*eventSubscriptionFilterResponsePtrType)(v)
}

func (*eventSubscriptionFilterResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionFilterResponse)(nil)).Elem()
}

func (i *eventSubscriptionFilterResponsePtrType) ToEventSubscriptionFilterResponsePtrOutput() EventSubscriptionFilterResponsePtrOutput {
	return i.ToEventSubscriptionFilterResponsePtrOutputWithContext(context.Background())
}

func (i *eventSubscriptionFilterResponsePtrType) ToEventSubscriptionFilterResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterResponsePtrOutput)
}

type EventSubscriptionFilterResponseOutput struct{ *pulumi.OutputState }

func (EventSubscriptionFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionFilterResponse)(nil)).Elem()
}

func (o EventSubscriptionFilterResponseOutput) ToEventSubscriptionFilterResponseOutput() EventSubscriptionFilterResponseOutput {
	return o
}

func (o EventSubscriptionFilterResponseOutput) ToEventSubscriptionFilterResponseOutputWithContext(ctx context.Context) EventSubscriptionFilterResponseOutput {
	return o
}

func (o EventSubscriptionFilterResponseOutput) ToEventSubscriptionFilterResponsePtrOutput() EventSubscriptionFilterResponsePtrOutput {
	return o.ToEventSubscriptionFilterResponsePtrOutputWithContext(context.Background())
}

func (o EventSubscriptionFilterResponseOutput) ToEventSubscriptionFilterResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSubscriptionFilterResponse) *EventSubscriptionFilterResponse {
		return &v
	}).(EventSubscriptionFilterResponsePtrOutput)
}

func (o EventSubscriptionFilterResponseOutput) AdvancedFilters() pulumi.ArrayOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) []interface{} { return v.AdvancedFilters }).(pulumi.ArrayOutput)
}

func (o EventSubscriptionFilterResponseOutput) EnableAdvancedFilteringOnArrays() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) *bool { return v.EnableAdvancedFilteringOnArrays }).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterResponseOutput) IncludedEventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) []string { return v.IncludedEventTypes }).(pulumi.StringArrayOutput)
}

func (o EventSubscriptionFilterResponseOutput) IsSubjectCaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) *bool { return v.IsSubjectCaseSensitive }).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterResponseOutput) SubjectBeginsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) *string { return v.SubjectBeginsWith }).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionFilterResponseOutput) SubjectEndsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) *string { return v.SubjectEndsWith }).(pulumi.StringPtrOutput)
}

type EventSubscriptionFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionFilterResponse)(nil)).Elem()
}

func (o EventSubscriptionFilterResponsePtrOutput) ToEventSubscriptionFilterResponsePtrOutput() EventSubscriptionFilterResponsePtrOutput {
	return o
}

func (o EventSubscriptionFilterResponsePtrOutput) ToEventSubscriptionFilterResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionFilterResponsePtrOutput {
	return o
}

func (o EventSubscriptionFilterResponsePtrOutput) Elem() EventSubscriptionFilterResponseOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) EventSubscriptionFilterResponse {
		if v != nil {
			return *v
		}
		var ret EventSubscriptionFilterResponse
		return ret
	}).(EventSubscriptionFilterResponseOutput)
}

func (o EventSubscriptionFilterResponsePtrOutput) AdvancedFilters() pulumi.ArrayOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) []interface{} {
		if v == nil {
			return nil
		}
		return v.AdvancedFilters
	}).(pulumi.ArrayOutput)
}

func (o EventSubscriptionFilterResponsePtrOutput) EnableAdvancedFilteringOnArrays() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAdvancedFilteringOnArrays
	}).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterResponsePtrOutput) IncludedEventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) []string {
		if v == nil {
			return nil
		}
		return v.IncludedEventTypes
	}).(pulumi.StringArrayOutput)
}

func (o EventSubscriptionFilterResponsePtrOutput) IsSubjectCaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsSubjectCaseSensitive
	}).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterResponsePtrOutput) SubjectBeginsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubjectBeginsWith
	}).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionFilterResponsePtrOutput) SubjectEndsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubjectEndsWith
	}).(pulumi.StringPtrOutput)
}

type EventSubscriptionIdentity struct {
	Type                 *string `pulumi:"type"`
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}





type EventSubscriptionIdentityInput interface {
	pulumi.Input

	ToEventSubscriptionIdentityOutput() EventSubscriptionIdentityOutput
	ToEventSubscriptionIdentityOutputWithContext(context.Context) EventSubscriptionIdentityOutput
}

type EventSubscriptionIdentityArgs struct {
	Type                 pulumi.StringPtrInput `pulumi:"type"`
	UserAssignedIdentity pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (EventSubscriptionIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionIdentity)(nil)).Elem()
}

func (i EventSubscriptionIdentityArgs) ToEventSubscriptionIdentityOutput() EventSubscriptionIdentityOutput {
	return i.ToEventSubscriptionIdentityOutputWithContext(context.Background())
}

func (i EventSubscriptionIdentityArgs) ToEventSubscriptionIdentityOutputWithContext(ctx context.Context) EventSubscriptionIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionIdentityOutput)
}

func (i EventSubscriptionIdentityArgs) ToEventSubscriptionIdentityPtrOutput() EventSubscriptionIdentityPtrOutput {
	return i.ToEventSubscriptionIdentityPtrOutputWithContext(context.Background())
}

func (i EventSubscriptionIdentityArgs) ToEventSubscriptionIdentityPtrOutputWithContext(ctx context.Context) EventSubscriptionIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionIdentityOutput).ToEventSubscriptionIdentityPtrOutputWithContext(ctx)
}









type EventSubscriptionIdentityPtrInput interface {
	pulumi.Input

	ToEventSubscriptionIdentityPtrOutput() EventSubscriptionIdentityPtrOutput
	ToEventSubscriptionIdentityPtrOutputWithContext(context.Context) EventSubscriptionIdentityPtrOutput
}

type eventSubscriptionIdentityPtrType EventSubscriptionIdentityArgs

func EventSubscriptionIdentityPtr(v *EventSubscriptionIdentityArgs) EventSubscriptionIdentityPtrInput {
	return (*eventSubscriptionIdentityPtrType)(v)
}

func (*eventSubscriptionIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionIdentity)(nil)).Elem()
}

func (i *eventSubscriptionIdentityPtrType) ToEventSubscriptionIdentityPtrOutput() EventSubscriptionIdentityPtrOutput {
	return i.ToEventSubscriptionIdentityPtrOutputWithContext(context.Background())
}

func (i *eventSubscriptionIdentityPtrType) ToEventSubscriptionIdentityPtrOutputWithContext(ctx context.Context) EventSubscriptionIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionIdentityPtrOutput)
}

type EventSubscriptionIdentityOutput struct{ *pulumi.OutputState }

func (EventSubscriptionIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionIdentity)(nil)).Elem()
}

func (o EventSubscriptionIdentityOutput) ToEventSubscriptionIdentityOutput() EventSubscriptionIdentityOutput {
	return o
}

func (o EventSubscriptionIdentityOutput) ToEventSubscriptionIdentityOutputWithContext(ctx context.Context) EventSubscriptionIdentityOutput {
	return o
}

func (o EventSubscriptionIdentityOutput) ToEventSubscriptionIdentityPtrOutput() EventSubscriptionIdentityPtrOutput {
	return o.ToEventSubscriptionIdentityPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionIdentityOutput) ToEventSubscriptionIdentityPtrOutputWithContext(ctx context.Context) EventSubscriptionIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSubscriptionIdentity) *EventSubscriptionIdentity {
		return &v
	}).(EventSubscriptionIdentityPtrOutput)
}

func (o EventSubscriptionIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionIdentityOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionIdentity) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type EventSubscriptionIdentityPtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionIdentity)(nil)).Elem()
}

func (o EventSubscriptionIdentityPtrOutput) ToEventSubscriptionIdentityPtrOutput() EventSubscriptionIdentityPtrOutput {
	return o
}

func (o EventSubscriptionIdentityPtrOutput) ToEventSubscriptionIdentityPtrOutputWithContext(ctx context.Context) EventSubscriptionIdentityPtrOutput {
	return o
}

func (o EventSubscriptionIdentityPtrOutput) Elem() EventSubscriptionIdentityOutput {
	return o.ApplyT(func(v *EventSubscriptionIdentity) EventSubscriptionIdentity {
		if v != nil {
			return *v
		}
		var ret EventSubscriptionIdentity
		return ret
	}).(EventSubscriptionIdentityOutput)
}

func (o EventSubscriptionIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionIdentityPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionIdentity) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type EventSubscriptionIdentityResponse struct {
	Type                 *string `pulumi:"type"`
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}





type EventSubscriptionIdentityResponseInput interface {
	pulumi.Input

	ToEventSubscriptionIdentityResponseOutput() EventSubscriptionIdentityResponseOutput
	ToEventSubscriptionIdentityResponseOutputWithContext(context.Context) EventSubscriptionIdentityResponseOutput
}

type EventSubscriptionIdentityResponseArgs struct {
	Type                 pulumi.StringPtrInput `pulumi:"type"`
	UserAssignedIdentity pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (EventSubscriptionIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionIdentityResponse)(nil)).Elem()
}

func (i EventSubscriptionIdentityResponseArgs) ToEventSubscriptionIdentityResponseOutput() EventSubscriptionIdentityResponseOutput {
	return i.ToEventSubscriptionIdentityResponseOutputWithContext(context.Background())
}

func (i EventSubscriptionIdentityResponseArgs) ToEventSubscriptionIdentityResponseOutputWithContext(ctx context.Context) EventSubscriptionIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionIdentityResponseOutput)
}

func (i EventSubscriptionIdentityResponseArgs) ToEventSubscriptionIdentityResponsePtrOutput() EventSubscriptionIdentityResponsePtrOutput {
	return i.ToEventSubscriptionIdentityResponsePtrOutputWithContext(context.Background())
}

func (i EventSubscriptionIdentityResponseArgs) ToEventSubscriptionIdentityResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionIdentityResponseOutput).ToEventSubscriptionIdentityResponsePtrOutputWithContext(ctx)
}









type EventSubscriptionIdentityResponsePtrInput interface {
	pulumi.Input

	ToEventSubscriptionIdentityResponsePtrOutput() EventSubscriptionIdentityResponsePtrOutput
	ToEventSubscriptionIdentityResponsePtrOutputWithContext(context.Context) EventSubscriptionIdentityResponsePtrOutput
}

type eventSubscriptionIdentityResponsePtrType EventSubscriptionIdentityResponseArgs

func EventSubscriptionIdentityResponsePtr(v *EventSubscriptionIdentityResponseArgs) EventSubscriptionIdentityResponsePtrInput {
	return (*eventSubscriptionIdentityResponsePtrType)(v)
}

func (*eventSubscriptionIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionIdentityResponse)(nil)).Elem()
}

func (i *eventSubscriptionIdentityResponsePtrType) ToEventSubscriptionIdentityResponsePtrOutput() EventSubscriptionIdentityResponsePtrOutput {
	return i.ToEventSubscriptionIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *eventSubscriptionIdentityResponsePtrType) ToEventSubscriptionIdentityResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionIdentityResponsePtrOutput)
}

type EventSubscriptionIdentityResponseOutput struct{ *pulumi.OutputState }

func (EventSubscriptionIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionIdentityResponse)(nil)).Elem()
}

func (o EventSubscriptionIdentityResponseOutput) ToEventSubscriptionIdentityResponseOutput() EventSubscriptionIdentityResponseOutput {
	return o
}

func (o EventSubscriptionIdentityResponseOutput) ToEventSubscriptionIdentityResponseOutputWithContext(ctx context.Context) EventSubscriptionIdentityResponseOutput {
	return o
}

func (o EventSubscriptionIdentityResponseOutput) ToEventSubscriptionIdentityResponsePtrOutput() EventSubscriptionIdentityResponsePtrOutput {
	return o.ToEventSubscriptionIdentityResponsePtrOutputWithContext(context.Background())
}

func (o EventSubscriptionIdentityResponseOutput) ToEventSubscriptionIdentityResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSubscriptionIdentityResponse) *EventSubscriptionIdentityResponse {
		return &v
	}).(EventSubscriptionIdentityResponsePtrOutput)
}

func (o EventSubscriptionIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionIdentityResponseOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionIdentityResponse) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type EventSubscriptionIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionIdentityResponse)(nil)).Elem()
}

func (o EventSubscriptionIdentityResponsePtrOutput) ToEventSubscriptionIdentityResponsePtrOutput() EventSubscriptionIdentityResponsePtrOutput {
	return o
}

func (o EventSubscriptionIdentityResponsePtrOutput) ToEventSubscriptionIdentityResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionIdentityResponsePtrOutput {
	return o
}

func (o EventSubscriptionIdentityResponsePtrOutput) Elem() EventSubscriptionIdentityResponseOutput {
	return o.ApplyT(func(v *EventSubscriptionIdentityResponse) EventSubscriptionIdentityResponse {
		if v != nil {
			return *v
		}
		var ret EventSubscriptionIdentityResponse
		return ret
	}).(EventSubscriptionIdentityResponseOutput)
}

func (o EventSubscriptionIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionIdentityResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type HybridConnectionEventSubscriptionDestination struct {
	DeliveryAttributeMappings []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType              string        `pulumi:"endpointType"`
	ResourceId                *string       `pulumi:"resourceId"`
}





type HybridConnectionEventSubscriptionDestinationInput interface {
	pulumi.Input

	ToHybridConnectionEventSubscriptionDestinationOutput() HybridConnectionEventSubscriptionDestinationOutput
	ToHybridConnectionEventSubscriptionDestinationOutputWithContext(context.Context) HybridConnectionEventSubscriptionDestinationOutput
}

type HybridConnectionEventSubscriptionDestinationArgs struct {
	DeliveryAttributeMappings pulumi.ArrayInput     `pulumi:"deliveryAttributeMappings"`
	EndpointType              pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId                pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (HybridConnectionEventSubscriptionDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnectionEventSubscriptionDestination)(nil)).Elem()
}

func (i HybridConnectionEventSubscriptionDestinationArgs) ToHybridConnectionEventSubscriptionDestinationOutput() HybridConnectionEventSubscriptionDestinationOutput {
	return i.ToHybridConnectionEventSubscriptionDestinationOutputWithContext(context.Background())
}

func (i HybridConnectionEventSubscriptionDestinationArgs) ToHybridConnectionEventSubscriptionDestinationOutputWithContext(ctx context.Context) HybridConnectionEventSubscriptionDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridConnectionEventSubscriptionDestinationOutput)
}

type HybridConnectionEventSubscriptionDestinationOutput struct{ *pulumi.OutputState }

func (HybridConnectionEventSubscriptionDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnectionEventSubscriptionDestination)(nil)).Elem()
}

func (o HybridConnectionEventSubscriptionDestinationOutput) ToHybridConnectionEventSubscriptionDestinationOutput() HybridConnectionEventSubscriptionDestinationOutput {
	return o
}

func (o HybridConnectionEventSubscriptionDestinationOutput) ToHybridConnectionEventSubscriptionDestinationOutputWithContext(ctx context.Context) HybridConnectionEventSubscriptionDestinationOutput {
	return o
}

func (o HybridConnectionEventSubscriptionDestinationOutput) DeliveryAttributeMappings() pulumi.ArrayOutput {
	return o.ApplyT(func(v HybridConnectionEventSubscriptionDestination) []interface{} { return v.DeliveryAttributeMappings }).(pulumi.ArrayOutput)
}

func (o HybridConnectionEventSubscriptionDestinationOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v HybridConnectionEventSubscriptionDestination) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o HybridConnectionEventSubscriptionDestinationOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridConnectionEventSubscriptionDestination) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type HybridConnectionEventSubscriptionDestinationResponse struct {
	DeliveryAttributeMappings []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType              string        `pulumi:"endpointType"`
	ResourceId                *string       `pulumi:"resourceId"`
}





type HybridConnectionEventSubscriptionDestinationResponseInput interface {
	pulumi.Input

	ToHybridConnectionEventSubscriptionDestinationResponseOutput() HybridConnectionEventSubscriptionDestinationResponseOutput
	ToHybridConnectionEventSubscriptionDestinationResponseOutputWithContext(context.Context) HybridConnectionEventSubscriptionDestinationResponseOutput
}

type HybridConnectionEventSubscriptionDestinationResponseArgs struct {
	DeliveryAttributeMappings pulumi.ArrayInput     `pulumi:"deliveryAttributeMappings"`
	EndpointType              pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId                pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (HybridConnectionEventSubscriptionDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnectionEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (i HybridConnectionEventSubscriptionDestinationResponseArgs) ToHybridConnectionEventSubscriptionDestinationResponseOutput() HybridConnectionEventSubscriptionDestinationResponseOutput {
	return i.ToHybridConnectionEventSubscriptionDestinationResponseOutputWithContext(context.Background())
}

func (i HybridConnectionEventSubscriptionDestinationResponseArgs) ToHybridConnectionEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) HybridConnectionEventSubscriptionDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridConnectionEventSubscriptionDestinationResponseOutput)
}

type HybridConnectionEventSubscriptionDestinationResponseOutput struct{ *pulumi.OutputState }

func (HybridConnectionEventSubscriptionDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnectionEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (o HybridConnectionEventSubscriptionDestinationResponseOutput) ToHybridConnectionEventSubscriptionDestinationResponseOutput() HybridConnectionEventSubscriptionDestinationResponseOutput {
	return o
}

func (o HybridConnectionEventSubscriptionDestinationResponseOutput) ToHybridConnectionEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) HybridConnectionEventSubscriptionDestinationResponseOutput {
	return o
}

func (o HybridConnectionEventSubscriptionDestinationResponseOutput) DeliveryAttributeMappings() pulumi.ArrayOutput {
	return o.ApplyT(func(v HybridConnectionEventSubscriptionDestinationResponse) []interface{} {
		return v.DeliveryAttributeMappings
	}).(pulumi.ArrayOutput)
}

func (o HybridConnectionEventSubscriptionDestinationResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v HybridConnectionEventSubscriptionDestinationResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o HybridConnectionEventSubscriptionDestinationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridConnectionEventSubscriptionDestinationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type IdentityInfo struct {
	PrincipalId            *string                           `pulumi:"principalId"`
	TenantId               *string                           `pulumi:"tenantId"`
	Type                   *string                           `pulumi:"type"`
	UserAssignedIdentities map[string]UserIdentityProperties `pulumi:"userAssignedIdentities"`
}





type IdentityInfoInput interface {
	pulumi.Input

	ToIdentityInfoOutput() IdentityInfoOutput
	ToIdentityInfoOutputWithContext(context.Context) IdentityInfoOutput
}

type IdentityInfoArgs struct {
	PrincipalId            pulumi.StringPtrInput          `pulumi:"principalId"`
	TenantId               pulumi.StringPtrInput          `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput          `pulumi:"type"`
	UserAssignedIdentities UserIdentityPropertiesMapInput `pulumi:"userAssignedIdentities"`
}

func (IdentityInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityInfo)(nil)).Elem()
}

func (i IdentityInfoArgs) ToIdentityInfoOutput() IdentityInfoOutput {
	return i.ToIdentityInfoOutputWithContext(context.Background())
}

func (i IdentityInfoArgs) ToIdentityInfoOutputWithContext(ctx context.Context) IdentityInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityInfoOutput)
}

func (i IdentityInfoArgs) ToIdentityInfoPtrOutput() IdentityInfoPtrOutput {
	return i.ToIdentityInfoPtrOutputWithContext(context.Background())
}

func (i IdentityInfoArgs) ToIdentityInfoPtrOutputWithContext(ctx context.Context) IdentityInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityInfoOutput).ToIdentityInfoPtrOutputWithContext(ctx)
}









type IdentityInfoPtrInput interface {
	pulumi.Input

	ToIdentityInfoPtrOutput() IdentityInfoPtrOutput
	ToIdentityInfoPtrOutputWithContext(context.Context) IdentityInfoPtrOutput
}

type identityInfoPtrType IdentityInfoArgs

func IdentityInfoPtr(v *IdentityInfoArgs) IdentityInfoPtrInput {
	return (*identityInfoPtrType)(v)
}

func (*identityInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityInfo)(nil)).Elem()
}

func (i *identityInfoPtrType) ToIdentityInfoPtrOutput() IdentityInfoPtrOutput {
	return i.ToIdentityInfoPtrOutputWithContext(context.Background())
}

func (i *identityInfoPtrType) ToIdentityInfoPtrOutputWithContext(ctx context.Context) IdentityInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityInfoPtrOutput)
}

type IdentityInfoOutput struct{ *pulumi.OutputState }

func (IdentityInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityInfo)(nil)).Elem()
}

func (o IdentityInfoOutput) ToIdentityInfoOutput() IdentityInfoOutput {
	return o
}

func (o IdentityInfoOutput) ToIdentityInfoOutputWithContext(ctx context.Context) IdentityInfoOutput {
	return o
}

func (o IdentityInfoOutput) ToIdentityInfoPtrOutput() IdentityInfoPtrOutput {
	return o.ToIdentityInfoPtrOutputWithContext(context.Background())
}

func (o IdentityInfoOutput) ToIdentityInfoPtrOutputWithContext(ctx context.Context) IdentityInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityInfo) *IdentityInfo {
		return &v
	}).(IdentityInfoPtrOutput)
}

func (o IdentityInfoOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityInfo) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityInfoOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityInfo) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o IdentityInfoOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityInfo) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityInfoOutput) UserAssignedIdentities() UserIdentityPropertiesMapOutput {
	return o.ApplyT(func(v IdentityInfo) map[string]UserIdentityProperties { return v.UserAssignedIdentities }).(UserIdentityPropertiesMapOutput)
}

type IdentityInfoPtrOutput struct{ *pulumi.OutputState }

func (IdentityInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityInfo)(nil)).Elem()
}

func (o IdentityInfoPtrOutput) ToIdentityInfoPtrOutput() IdentityInfoPtrOutput {
	return o
}

func (o IdentityInfoPtrOutput) ToIdentityInfoPtrOutputWithContext(ctx context.Context) IdentityInfoPtrOutput {
	return o
}

func (o IdentityInfoPtrOutput) Elem() IdentityInfoOutput {
	return o.ApplyT(func(v *IdentityInfo) IdentityInfo {
		if v != nil {
			return *v
		}
		var ret IdentityInfo
		return ret
	}).(IdentityInfoOutput)
}

func (o IdentityInfoPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityInfo) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityInfoPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityInfo) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityInfoPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityInfo) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityInfoPtrOutput) UserAssignedIdentities() UserIdentityPropertiesMapOutput {
	return o.ApplyT(func(v *IdentityInfo) map[string]UserIdentityProperties {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserIdentityPropertiesMapOutput)
}

type IdentityInfoResponse struct {
	PrincipalId            *string                                   `pulumi:"principalId"`
	TenantId               *string                                   `pulumi:"tenantId"`
	Type                   *string                                   `pulumi:"type"`
	UserAssignedIdentities map[string]UserIdentityPropertiesResponse `pulumi:"userAssignedIdentities"`
}





type IdentityInfoResponseInput interface {
	pulumi.Input

	ToIdentityInfoResponseOutput() IdentityInfoResponseOutput
	ToIdentityInfoResponseOutputWithContext(context.Context) IdentityInfoResponseOutput
}

type IdentityInfoResponseArgs struct {
	PrincipalId            pulumi.StringPtrInput                  `pulumi:"principalId"`
	TenantId               pulumi.StringPtrInput                  `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                  `pulumi:"type"`
	UserAssignedIdentities UserIdentityPropertiesResponseMapInput `pulumi:"userAssignedIdentities"`
}

func (IdentityInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityInfoResponse)(nil)).Elem()
}

func (i IdentityInfoResponseArgs) ToIdentityInfoResponseOutput() IdentityInfoResponseOutput {
	return i.ToIdentityInfoResponseOutputWithContext(context.Background())
}

func (i IdentityInfoResponseArgs) ToIdentityInfoResponseOutputWithContext(ctx context.Context) IdentityInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityInfoResponseOutput)
}

func (i IdentityInfoResponseArgs) ToIdentityInfoResponsePtrOutput() IdentityInfoResponsePtrOutput {
	return i.ToIdentityInfoResponsePtrOutputWithContext(context.Background())
}

func (i IdentityInfoResponseArgs) ToIdentityInfoResponsePtrOutputWithContext(ctx context.Context) IdentityInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityInfoResponseOutput).ToIdentityInfoResponsePtrOutputWithContext(ctx)
}









type IdentityInfoResponsePtrInput interface {
	pulumi.Input

	ToIdentityInfoResponsePtrOutput() IdentityInfoResponsePtrOutput
	ToIdentityInfoResponsePtrOutputWithContext(context.Context) IdentityInfoResponsePtrOutput
}

type identityInfoResponsePtrType IdentityInfoResponseArgs

func IdentityInfoResponsePtr(v *IdentityInfoResponseArgs) IdentityInfoResponsePtrInput {
	return (*identityInfoResponsePtrType)(v)
}

func (*identityInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityInfoResponse)(nil)).Elem()
}

func (i *identityInfoResponsePtrType) ToIdentityInfoResponsePtrOutput() IdentityInfoResponsePtrOutput {
	return i.ToIdentityInfoResponsePtrOutputWithContext(context.Background())
}

func (i *identityInfoResponsePtrType) ToIdentityInfoResponsePtrOutputWithContext(ctx context.Context) IdentityInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityInfoResponsePtrOutput)
}

type IdentityInfoResponseOutput struct{ *pulumi.OutputState }

func (IdentityInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityInfoResponse)(nil)).Elem()
}

func (o IdentityInfoResponseOutput) ToIdentityInfoResponseOutput() IdentityInfoResponseOutput {
	return o
}

func (o IdentityInfoResponseOutput) ToIdentityInfoResponseOutputWithContext(ctx context.Context) IdentityInfoResponseOutput {
	return o
}

func (o IdentityInfoResponseOutput) ToIdentityInfoResponsePtrOutput() IdentityInfoResponsePtrOutput {
	return o.ToIdentityInfoResponsePtrOutputWithContext(context.Background())
}

func (o IdentityInfoResponseOutput) ToIdentityInfoResponsePtrOutputWithContext(ctx context.Context) IdentityInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityInfoResponse) *IdentityInfoResponse {
		return &v
	}).(IdentityInfoResponsePtrOutput)
}

func (o IdentityInfoResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityInfoResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityInfoResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityInfoResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o IdentityInfoResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityInfoResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityInfoResponseOutput) UserAssignedIdentities() UserIdentityPropertiesResponseMapOutput {
	return o.ApplyT(func(v IdentityInfoResponse) map[string]UserIdentityPropertiesResponse {
		return v.UserAssignedIdentities
	}).(UserIdentityPropertiesResponseMapOutput)
}

type IdentityInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityInfoResponse)(nil)).Elem()
}

func (o IdentityInfoResponsePtrOutput) ToIdentityInfoResponsePtrOutput() IdentityInfoResponsePtrOutput {
	return o
}

func (o IdentityInfoResponsePtrOutput) ToIdentityInfoResponsePtrOutputWithContext(ctx context.Context) IdentityInfoResponsePtrOutput {
	return o
}

func (o IdentityInfoResponsePtrOutput) Elem() IdentityInfoResponseOutput {
	return o.ApplyT(func(v *IdentityInfoResponse) IdentityInfoResponse {
		if v != nil {
			return *v
		}
		var ret IdentityInfoResponse
		return ret
	}).(IdentityInfoResponseOutput)
}

func (o IdentityInfoResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityInfoResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityInfoResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityInfoResponsePtrOutput) UserAssignedIdentities() UserIdentityPropertiesResponseMapOutput {
	return o.ApplyT(func(v *IdentityInfoResponse) map[string]UserIdentityPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserIdentityPropertiesResponseMapOutput)
}

type InboundIpRule struct {
	Action *string `pulumi:"action"`
	IpMask *string `pulumi:"ipMask"`
}





type InboundIpRuleInput interface {
	pulumi.Input

	ToInboundIpRuleOutput() InboundIpRuleOutput
	ToInboundIpRuleOutputWithContext(context.Context) InboundIpRuleOutput
}

type InboundIpRuleArgs struct {
	Action pulumi.StringPtrInput `pulumi:"action"`
	IpMask pulumi.StringPtrInput `pulumi:"ipMask"`
}

func (InboundIpRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundIpRule)(nil)).Elem()
}

func (i InboundIpRuleArgs) ToInboundIpRuleOutput() InboundIpRuleOutput {
	return i.ToInboundIpRuleOutputWithContext(context.Background())
}

func (i InboundIpRuleArgs) ToInboundIpRuleOutputWithContext(ctx context.Context) InboundIpRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundIpRuleOutput)
}





type InboundIpRuleArrayInput interface {
	pulumi.Input

	ToInboundIpRuleArrayOutput() InboundIpRuleArrayOutput
	ToInboundIpRuleArrayOutputWithContext(context.Context) InboundIpRuleArrayOutput
}

type InboundIpRuleArray []InboundIpRuleInput

func (InboundIpRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundIpRule)(nil)).Elem()
}

func (i InboundIpRuleArray) ToInboundIpRuleArrayOutput() InboundIpRuleArrayOutput {
	return i.ToInboundIpRuleArrayOutputWithContext(context.Background())
}

func (i InboundIpRuleArray) ToInboundIpRuleArrayOutputWithContext(ctx context.Context) InboundIpRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundIpRuleArrayOutput)
}

type InboundIpRuleOutput struct{ *pulumi.OutputState }

func (InboundIpRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundIpRule)(nil)).Elem()
}

func (o InboundIpRuleOutput) ToInboundIpRuleOutput() InboundIpRuleOutput {
	return o
}

func (o InboundIpRuleOutput) ToInboundIpRuleOutputWithContext(ctx context.Context) InboundIpRuleOutput {
	return o
}

func (o InboundIpRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundIpRule) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o InboundIpRuleOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundIpRule) *string { return v.IpMask }).(pulumi.StringPtrOutput)
}

type InboundIpRuleArrayOutput struct{ *pulumi.OutputState }

func (InboundIpRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundIpRule)(nil)).Elem()
}

func (o InboundIpRuleArrayOutput) ToInboundIpRuleArrayOutput() InboundIpRuleArrayOutput {
	return o
}

func (o InboundIpRuleArrayOutput) ToInboundIpRuleArrayOutputWithContext(ctx context.Context) InboundIpRuleArrayOutput {
	return o
}

func (o InboundIpRuleArrayOutput) Index(i pulumi.IntInput) InboundIpRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundIpRule {
		return vs[0].([]InboundIpRule)[vs[1].(int)]
	}).(InboundIpRuleOutput)
}

type InboundIpRuleResponse struct {
	Action *string `pulumi:"action"`
	IpMask *string `pulumi:"ipMask"`
}





type InboundIpRuleResponseInput interface {
	pulumi.Input

	ToInboundIpRuleResponseOutput() InboundIpRuleResponseOutput
	ToInboundIpRuleResponseOutputWithContext(context.Context) InboundIpRuleResponseOutput
}

type InboundIpRuleResponseArgs struct {
	Action pulumi.StringPtrInput `pulumi:"action"`
	IpMask pulumi.StringPtrInput `pulumi:"ipMask"`
}

func (InboundIpRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundIpRuleResponse)(nil)).Elem()
}

func (i InboundIpRuleResponseArgs) ToInboundIpRuleResponseOutput() InboundIpRuleResponseOutput {
	return i.ToInboundIpRuleResponseOutputWithContext(context.Background())
}

func (i InboundIpRuleResponseArgs) ToInboundIpRuleResponseOutputWithContext(ctx context.Context) InboundIpRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundIpRuleResponseOutput)
}





type InboundIpRuleResponseArrayInput interface {
	pulumi.Input

	ToInboundIpRuleResponseArrayOutput() InboundIpRuleResponseArrayOutput
	ToInboundIpRuleResponseArrayOutputWithContext(context.Context) InboundIpRuleResponseArrayOutput
}

type InboundIpRuleResponseArray []InboundIpRuleResponseInput

func (InboundIpRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundIpRuleResponse)(nil)).Elem()
}

func (i InboundIpRuleResponseArray) ToInboundIpRuleResponseArrayOutput() InboundIpRuleResponseArrayOutput {
	return i.ToInboundIpRuleResponseArrayOutputWithContext(context.Background())
}

func (i InboundIpRuleResponseArray) ToInboundIpRuleResponseArrayOutputWithContext(ctx context.Context) InboundIpRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundIpRuleResponseArrayOutput)
}

type InboundIpRuleResponseOutput struct{ *pulumi.OutputState }

func (InboundIpRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundIpRuleResponse)(nil)).Elem()
}

func (o InboundIpRuleResponseOutput) ToInboundIpRuleResponseOutput() InboundIpRuleResponseOutput {
	return o
}

func (o InboundIpRuleResponseOutput) ToInboundIpRuleResponseOutputWithContext(ctx context.Context) InboundIpRuleResponseOutput {
	return o
}

func (o InboundIpRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundIpRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o InboundIpRuleResponseOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundIpRuleResponse) *string { return v.IpMask }).(pulumi.StringPtrOutput)
}

type InboundIpRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (InboundIpRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundIpRuleResponse)(nil)).Elem()
}

func (o InboundIpRuleResponseArrayOutput) ToInboundIpRuleResponseArrayOutput() InboundIpRuleResponseArrayOutput {
	return o
}

func (o InboundIpRuleResponseArrayOutput) ToInboundIpRuleResponseArrayOutputWithContext(ctx context.Context) InboundIpRuleResponseArrayOutput {
	return o
}

func (o InboundIpRuleResponseArrayOutput) Index(i pulumi.IntInput) InboundIpRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundIpRuleResponse {
		return vs[0].([]InboundIpRuleResponse)[vs[1].(int)]
	}).(InboundIpRuleResponseOutput)
}

type JsonField struct {
	SourceField *string `pulumi:"sourceField"`
}





type JsonFieldInput interface {
	pulumi.Input

	ToJsonFieldOutput() JsonFieldOutput
	ToJsonFieldOutputWithContext(context.Context) JsonFieldOutput
}

type JsonFieldArgs struct {
	SourceField pulumi.StringPtrInput `pulumi:"sourceField"`
}

func (JsonFieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonField)(nil)).Elem()
}

func (i JsonFieldArgs) ToJsonFieldOutput() JsonFieldOutput {
	return i.ToJsonFieldOutputWithContext(context.Background())
}

func (i JsonFieldArgs) ToJsonFieldOutputWithContext(ctx context.Context) JsonFieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldOutput)
}

func (i JsonFieldArgs) ToJsonFieldPtrOutput() JsonFieldPtrOutput {
	return i.ToJsonFieldPtrOutputWithContext(context.Background())
}

func (i JsonFieldArgs) ToJsonFieldPtrOutputWithContext(ctx context.Context) JsonFieldPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldOutput).ToJsonFieldPtrOutputWithContext(ctx)
}









type JsonFieldPtrInput interface {
	pulumi.Input

	ToJsonFieldPtrOutput() JsonFieldPtrOutput
	ToJsonFieldPtrOutputWithContext(context.Context) JsonFieldPtrOutput
}

type jsonFieldPtrType JsonFieldArgs

func JsonFieldPtr(v *JsonFieldArgs) JsonFieldPtrInput {
	return (*jsonFieldPtrType)(v)
}

func (*jsonFieldPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonField)(nil)).Elem()
}

func (i *jsonFieldPtrType) ToJsonFieldPtrOutput() JsonFieldPtrOutput {
	return i.ToJsonFieldPtrOutputWithContext(context.Background())
}

func (i *jsonFieldPtrType) ToJsonFieldPtrOutputWithContext(ctx context.Context) JsonFieldPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldPtrOutput)
}

type JsonFieldOutput struct{ *pulumi.OutputState }

func (JsonFieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonField)(nil)).Elem()
}

func (o JsonFieldOutput) ToJsonFieldOutput() JsonFieldOutput {
	return o
}

func (o JsonFieldOutput) ToJsonFieldOutputWithContext(ctx context.Context) JsonFieldOutput {
	return o
}

func (o JsonFieldOutput) ToJsonFieldPtrOutput() JsonFieldPtrOutput {
	return o.ToJsonFieldPtrOutputWithContext(context.Background())
}

func (o JsonFieldOutput) ToJsonFieldPtrOutputWithContext(ctx context.Context) JsonFieldPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JsonField) *JsonField {
		return &v
	}).(JsonFieldPtrOutput)
}

func (o JsonFieldOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonField) *string { return v.SourceField }).(pulumi.StringPtrOutput)
}

type JsonFieldPtrOutput struct{ *pulumi.OutputState }

func (JsonFieldPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonField)(nil)).Elem()
}

func (o JsonFieldPtrOutput) ToJsonFieldPtrOutput() JsonFieldPtrOutput {
	return o
}

func (o JsonFieldPtrOutput) ToJsonFieldPtrOutputWithContext(ctx context.Context) JsonFieldPtrOutput {
	return o
}

func (o JsonFieldPtrOutput) Elem() JsonFieldOutput {
	return o.ApplyT(func(v *JsonField) JsonField {
		if v != nil {
			return *v
		}
		var ret JsonField
		return ret
	}).(JsonFieldOutput)
}

func (o JsonFieldPtrOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonField) *string {
		if v == nil {
			return nil
		}
		return v.SourceField
	}).(pulumi.StringPtrOutput)
}

type JsonFieldResponse struct {
	SourceField *string `pulumi:"sourceField"`
}





type JsonFieldResponseInput interface {
	pulumi.Input

	ToJsonFieldResponseOutput() JsonFieldResponseOutput
	ToJsonFieldResponseOutputWithContext(context.Context) JsonFieldResponseOutput
}

type JsonFieldResponseArgs struct {
	SourceField pulumi.StringPtrInput `pulumi:"sourceField"`
}

func (JsonFieldResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonFieldResponse)(nil)).Elem()
}

func (i JsonFieldResponseArgs) ToJsonFieldResponseOutput() JsonFieldResponseOutput {
	return i.ToJsonFieldResponseOutputWithContext(context.Background())
}

func (i JsonFieldResponseArgs) ToJsonFieldResponseOutputWithContext(ctx context.Context) JsonFieldResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldResponseOutput)
}

func (i JsonFieldResponseArgs) ToJsonFieldResponsePtrOutput() JsonFieldResponsePtrOutput {
	return i.ToJsonFieldResponsePtrOutputWithContext(context.Background())
}

func (i JsonFieldResponseArgs) ToJsonFieldResponsePtrOutputWithContext(ctx context.Context) JsonFieldResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldResponseOutput).ToJsonFieldResponsePtrOutputWithContext(ctx)
}









type JsonFieldResponsePtrInput interface {
	pulumi.Input

	ToJsonFieldResponsePtrOutput() JsonFieldResponsePtrOutput
	ToJsonFieldResponsePtrOutputWithContext(context.Context) JsonFieldResponsePtrOutput
}

type jsonFieldResponsePtrType JsonFieldResponseArgs

func JsonFieldResponsePtr(v *JsonFieldResponseArgs) JsonFieldResponsePtrInput {
	return (*jsonFieldResponsePtrType)(v)
}

func (*jsonFieldResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonFieldResponse)(nil)).Elem()
}

func (i *jsonFieldResponsePtrType) ToJsonFieldResponsePtrOutput() JsonFieldResponsePtrOutput {
	return i.ToJsonFieldResponsePtrOutputWithContext(context.Background())
}

func (i *jsonFieldResponsePtrType) ToJsonFieldResponsePtrOutputWithContext(ctx context.Context) JsonFieldResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldResponsePtrOutput)
}

type JsonFieldResponseOutput struct{ *pulumi.OutputState }

func (JsonFieldResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonFieldResponse)(nil)).Elem()
}

func (o JsonFieldResponseOutput) ToJsonFieldResponseOutput() JsonFieldResponseOutput {
	return o
}

func (o JsonFieldResponseOutput) ToJsonFieldResponseOutputWithContext(ctx context.Context) JsonFieldResponseOutput {
	return o
}

func (o JsonFieldResponseOutput) ToJsonFieldResponsePtrOutput() JsonFieldResponsePtrOutput {
	return o.ToJsonFieldResponsePtrOutputWithContext(context.Background())
}

func (o JsonFieldResponseOutput) ToJsonFieldResponsePtrOutputWithContext(ctx context.Context) JsonFieldResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JsonFieldResponse) *JsonFieldResponse {
		return &v
	}).(JsonFieldResponsePtrOutput)
}

func (o JsonFieldResponseOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonFieldResponse) *string { return v.SourceField }).(pulumi.StringPtrOutput)
}

type JsonFieldResponsePtrOutput struct{ *pulumi.OutputState }

func (JsonFieldResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonFieldResponse)(nil)).Elem()
}

func (o JsonFieldResponsePtrOutput) ToJsonFieldResponsePtrOutput() JsonFieldResponsePtrOutput {
	return o
}

func (o JsonFieldResponsePtrOutput) ToJsonFieldResponsePtrOutputWithContext(ctx context.Context) JsonFieldResponsePtrOutput {
	return o
}

func (o JsonFieldResponsePtrOutput) Elem() JsonFieldResponseOutput {
	return o.ApplyT(func(v *JsonFieldResponse) JsonFieldResponse {
		if v != nil {
			return *v
		}
		var ret JsonFieldResponse
		return ret
	}).(JsonFieldResponseOutput)
}

func (o JsonFieldResponsePtrOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonFieldResponse) *string {
		if v == nil {
			return nil
		}
		return v.SourceField
	}).(pulumi.StringPtrOutput)
}

type JsonFieldWithDefault struct {
	DefaultValue *string `pulumi:"defaultValue"`
	SourceField  *string `pulumi:"sourceField"`
}





type JsonFieldWithDefaultInput interface {
	pulumi.Input

	ToJsonFieldWithDefaultOutput() JsonFieldWithDefaultOutput
	ToJsonFieldWithDefaultOutputWithContext(context.Context) JsonFieldWithDefaultOutput
}

type JsonFieldWithDefaultArgs struct {
	DefaultValue pulumi.StringPtrInput `pulumi:"defaultValue"`
	SourceField  pulumi.StringPtrInput `pulumi:"sourceField"`
}

func (JsonFieldWithDefaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonFieldWithDefault)(nil)).Elem()
}

func (i JsonFieldWithDefaultArgs) ToJsonFieldWithDefaultOutput() JsonFieldWithDefaultOutput {
	return i.ToJsonFieldWithDefaultOutputWithContext(context.Background())
}

func (i JsonFieldWithDefaultArgs) ToJsonFieldWithDefaultOutputWithContext(ctx context.Context) JsonFieldWithDefaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldWithDefaultOutput)
}

func (i JsonFieldWithDefaultArgs) ToJsonFieldWithDefaultPtrOutput() JsonFieldWithDefaultPtrOutput {
	return i.ToJsonFieldWithDefaultPtrOutputWithContext(context.Background())
}

func (i JsonFieldWithDefaultArgs) ToJsonFieldWithDefaultPtrOutputWithContext(ctx context.Context) JsonFieldWithDefaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldWithDefaultOutput).ToJsonFieldWithDefaultPtrOutputWithContext(ctx)
}









type JsonFieldWithDefaultPtrInput interface {
	pulumi.Input

	ToJsonFieldWithDefaultPtrOutput() JsonFieldWithDefaultPtrOutput
	ToJsonFieldWithDefaultPtrOutputWithContext(context.Context) JsonFieldWithDefaultPtrOutput
}

type jsonFieldWithDefaultPtrType JsonFieldWithDefaultArgs

func JsonFieldWithDefaultPtr(v *JsonFieldWithDefaultArgs) JsonFieldWithDefaultPtrInput {
	return (*jsonFieldWithDefaultPtrType)(v)
}

func (*jsonFieldWithDefaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonFieldWithDefault)(nil)).Elem()
}

func (i *jsonFieldWithDefaultPtrType) ToJsonFieldWithDefaultPtrOutput() JsonFieldWithDefaultPtrOutput {
	return i.ToJsonFieldWithDefaultPtrOutputWithContext(context.Background())
}

func (i *jsonFieldWithDefaultPtrType) ToJsonFieldWithDefaultPtrOutputWithContext(ctx context.Context) JsonFieldWithDefaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldWithDefaultPtrOutput)
}

type JsonFieldWithDefaultOutput struct{ *pulumi.OutputState }

func (JsonFieldWithDefaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonFieldWithDefault)(nil)).Elem()
}

func (o JsonFieldWithDefaultOutput) ToJsonFieldWithDefaultOutput() JsonFieldWithDefaultOutput {
	return o
}

func (o JsonFieldWithDefaultOutput) ToJsonFieldWithDefaultOutputWithContext(ctx context.Context) JsonFieldWithDefaultOutput {
	return o
}

func (o JsonFieldWithDefaultOutput) ToJsonFieldWithDefaultPtrOutput() JsonFieldWithDefaultPtrOutput {
	return o.ToJsonFieldWithDefaultPtrOutputWithContext(context.Background())
}

func (o JsonFieldWithDefaultOutput) ToJsonFieldWithDefaultPtrOutputWithContext(ctx context.Context) JsonFieldWithDefaultPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JsonFieldWithDefault) *JsonFieldWithDefault {
		return &v
	}).(JsonFieldWithDefaultPtrOutput)
}

func (o JsonFieldWithDefaultOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonFieldWithDefault) *string { return v.DefaultValue }).(pulumi.StringPtrOutput)
}

func (o JsonFieldWithDefaultOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonFieldWithDefault) *string { return v.SourceField }).(pulumi.StringPtrOutput)
}

type JsonFieldWithDefaultPtrOutput struct{ *pulumi.OutputState }

func (JsonFieldWithDefaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonFieldWithDefault)(nil)).Elem()
}

func (o JsonFieldWithDefaultPtrOutput) ToJsonFieldWithDefaultPtrOutput() JsonFieldWithDefaultPtrOutput {
	return o
}

func (o JsonFieldWithDefaultPtrOutput) ToJsonFieldWithDefaultPtrOutputWithContext(ctx context.Context) JsonFieldWithDefaultPtrOutput {
	return o
}

func (o JsonFieldWithDefaultPtrOutput) Elem() JsonFieldWithDefaultOutput {
	return o.ApplyT(func(v *JsonFieldWithDefault) JsonFieldWithDefault {
		if v != nil {
			return *v
		}
		var ret JsonFieldWithDefault
		return ret
	}).(JsonFieldWithDefaultOutput)
}

func (o JsonFieldWithDefaultPtrOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonFieldWithDefault) *string {
		if v == nil {
			return nil
		}
		return v.DefaultValue
	}).(pulumi.StringPtrOutput)
}

func (o JsonFieldWithDefaultPtrOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonFieldWithDefault) *string {
		if v == nil {
			return nil
		}
		return v.SourceField
	}).(pulumi.StringPtrOutput)
}

type JsonFieldWithDefaultResponse struct {
	DefaultValue *string `pulumi:"defaultValue"`
	SourceField  *string `pulumi:"sourceField"`
}





type JsonFieldWithDefaultResponseInput interface {
	pulumi.Input

	ToJsonFieldWithDefaultResponseOutput() JsonFieldWithDefaultResponseOutput
	ToJsonFieldWithDefaultResponseOutputWithContext(context.Context) JsonFieldWithDefaultResponseOutput
}

type JsonFieldWithDefaultResponseArgs struct {
	DefaultValue pulumi.StringPtrInput `pulumi:"defaultValue"`
	SourceField  pulumi.StringPtrInput `pulumi:"sourceField"`
}

func (JsonFieldWithDefaultResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonFieldWithDefaultResponse)(nil)).Elem()
}

func (i JsonFieldWithDefaultResponseArgs) ToJsonFieldWithDefaultResponseOutput() JsonFieldWithDefaultResponseOutput {
	return i.ToJsonFieldWithDefaultResponseOutputWithContext(context.Background())
}

func (i JsonFieldWithDefaultResponseArgs) ToJsonFieldWithDefaultResponseOutputWithContext(ctx context.Context) JsonFieldWithDefaultResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldWithDefaultResponseOutput)
}

func (i JsonFieldWithDefaultResponseArgs) ToJsonFieldWithDefaultResponsePtrOutput() JsonFieldWithDefaultResponsePtrOutput {
	return i.ToJsonFieldWithDefaultResponsePtrOutputWithContext(context.Background())
}

func (i JsonFieldWithDefaultResponseArgs) ToJsonFieldWithDefaultResponsePtrOutputWithContext(ctx context.Context) JsonFieldWithDefaultResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldWithDefaultResponseOutput).ToJsonFieldWithDefaultResponsePtrOutputWithContext(ctx)
}









type JsonFieldWithDefaultResponsePtrInput interface {
	pulumi.Input

	ToJsonFieldWithDefaultResponsePtrOutput() JsonFieldWithDefaultResponsePtrOutput
	ToJsonFieldWithDefaultResponsePtrOutputWithContext(context.Context) JsonFieldWithDefaultResponsePtrOutput
}

type jsonFieldWithDefaultResponsePtrType JsonFieldWithDefaultResponseArgs

func JsonFieldWithDefaultResponsePtr(v *JsonFieldWithDefaultResponseArgs) JsonFieldWithDefaultResponsePtrInput {
	return (*jsonFieldWithDefaultResponsePtrType)(v)
}

func (*jsonFieldWithDefaultResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonFieldWithDefaultResponse)(nil)).Elem()
}

func (i *jsonFieldWithDefaultResponsePtrType) ToJsonFieldWithDefaultResponsePtrOutput() JsonFieldWithDefaultResponsePtrOutput {
	return i.ToJsonFieldWithDefaultResponsePtrOutputWithContext(context.Background())
}

func (i *jsonFieldWithDefaultResponsePtrType) ToJsonFieldWithDefaultResponsePtrOutputWithContext(ctx context.Context) JsonFieldWithDefaultResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldWithDefaultResponsePtrOutput)
}

type JsonFieldWithDefaultResponseOutput struct{ *pulumi.OutputState }

func (JsonFieldWithDefaultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonFieldWithDefaultResponse)(nil)).Elem()
}

func (o JsonFieldWithDefaultResponseOutput) ToJsonFieldWithDefaultResponseOutput() JsonFieldWithDefaultResponseOutput {
	return o
}

func (o JsonFieldWithDefaultResponseOutput) ToJsonFieldWithDefaultResponseOutputWithContext(ctx context.Context) JsonFieldWithDefaultResponseOutput {
	return o
}

func (o JsonFieldWithDefaultResponseOutput) ToJsonFieldWithDefaultResponsePtrOutput() JsonFieldWithDefaultResponsePtrOutput {
	return o.ToJsonFieldWithDefaultResponsePtrOutputWithContext(context.Background())
}

func (o JsonFieldWithDefaultResponseOutput) ToJsonFieldWithDefaultResponsePtrOutputWithContext(ctx context.Context) JsonFieldWithDefaultResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JsonFieldWithDefaultResponse) *JsonFieldWithDefaultResponse {
		return &v
	}).(JsonFieldWithDefaultResponsePtrOutput)
}

func (o JsonFieldWithDefaultResponseOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonFieldWithDefaultResponse) *string { return v.DefaultValue }).(pulumi.StringPtrOutput)
}

func (o JsonFieldWithDefaultResponseOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonFieldWithDefaultResponse) *string { return v.SourceField }).(pulumi.StringPtrOutput)
}

type JsonFieldWithDefaultResponsePtrOutput struct{ *pulumi.OutputState }

func (JsonFieldWithDefaultResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonFieldWithDefaultResponse)(nil)).Elem()
}

func (o JsonFieldWithDefaultResponsePtrOutput) ToJsonFieldWithDefaultResponsePtrOutput() JsonFieldWithDefaultResponsePtrOutput {
	return o
}

func (o JsonFieldWithDefaultResponsePtrOutput) ToJsonFieldWithDefaultResponsePtrOutputWithContext(ctx context.Context) JsonFieldWithDefaultResponsePtrOutput {
	return o
}

func (o JsonFieldWithDefaultResponsePtrOutput) Elem() JsonFieldWithDefaultResponseOutput {
	return o.ApplyT(func(v *JsonFieldWithDefaultResponse) JsonFieldWithDefaultResponse {
		if v != nil {
			return *v
		}
		var ret JsonFieldWithDefaultResponse
		return ret
	}).(JsonFieldWithDefaultResponseOutput)
}

func (o JsonFieldWithDefaultResponsePtrOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonFieldWithDefaultResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultValue
	}).(pulumi.StringPtrOutput)
}

func (o JsonFieldWithDefaultResponsePtrOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonFieldWithDefaultResponse) *string {
		if v == nil {
			return nil
		}
		return v.SourceField
	}).(pulumi.StringPtrOutput)
}

type JsonInputSchemaMapping struct {
	DataVersion            *JsonFieldWithDefault `pulumi:"dataVersion"`
	EventTime              *JsonField            `pulumi:"eventTime"`
	EventType              *JsonFieldWithDefault `pulumi:"eventType"`
	Id                     *JsonField            `pulumi:"id"`
	InputSchemaMappingType string                `pulumi:"inputSchemaMappingType"`
	Subject                *JsonFieldWithDefault `pulumi:"subject"`
	Topic                  *JsonField            `pulumi:"topic"`
}





type JsonInputSchemaMappingInput interface {
	pulumi.Input

	ToJsonInputSchemaMappingOutput() JsonInputSchemaMappingOutput
	ToJsonInputSchemaMappingOutputWithContext(context.Context) JsonInputSchemaMappingOutput
}

type JsonInputSchemaMappingArgs struct {
	DataVersion            JsonFieldWithDefaultPtrInput `pulumi:"dataVersion"`
	EventTime              JsonFieldPtrInput            `pulumi:"eventTime"`
	EventType              JsonFieldWithDefaultPtrInput `pulumi:"eventType"`
	Id                     JsonFieldPtrInput            `pulumi:"id"`
	InputSchemaMappingType pulumi.StringInput           `pulumi:"inputSchemaMappingType"`
	Subject                JsonFieldWithDefaultPtrInput `pulumi:"subject"`
	Topic                  JsonFieldPtrInput            `pulumi:"topic"`
}

func (JsonInputSchemaMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonInputSchemaMapping)(nil)).Elem()
}

func (i JsonInputSchemaMappingArgs) ToJsonInputSchemaMappingOutput() JsonInputSchemaMappingOutput {
	return i.ToJsonInputSchemaMappingOutputWithContext(context.Background())
}

func (i JsonInputSchemaMappingArgs) ToJsonInputSchemaMappingOutputWithContext(ctx context.Context) JsonInputSchemaMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonInputSchemaMappingOutput)
}

func (i JsonInputSchemaMappingArgs) ToJsonInputSchemaMappingPtrOutput() JsonInputSchemaMappingPtrOutput {
	return i.ToJsonInputSchemaMappingPtrOutputWithContext(context.Background())
}

func (i JsonInputSchemaMappingArgs) ToJsonInputSchemaMappingPtrOutputWithContext(ctx context.Context) JsonInputSchemaMappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonInputSchemaMappingOutput).ToJsonInputSchemaMappingPtrOutputWithContext(ctx)
}









type JsonInputSchemaMappingPtrInput interface {
	pulumi.Input

	ToJsonInputSchemaMappingPtrOutput() JsonInputSchemaMappingPtrOutput
	ToJsonInputSchemaMappingPtrOutputWithContext(context.Context) JsonInputSchemaMappingPtrOutput
}

type jsonInputSchemaMappingPtrType JsonInputSchemaMappingArgs

func JsonInputSchemaMappingPtr(v *JsonInputSchemaMappingArgs) JsonInputSchemaMappingPtrInput {
	return (*jsonInputSchemaMappingPtrType)(v)
}

func (*jsonInputSchemaMappingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonInputSchemaMapping)(nil)).Elem()
}

func (i *jsonInputSchemaMappingPtrType) ToJsonInputSchemaMappingPtrOutput() JsonInputSchemaMappingPtrOutput {
	return i.ToJsonInputSchemaMappingPtrOutputWithContext(context.Background())
}

func (i *jsonInputSchemaMappingPtrType) ToJsonInputSchemaMappingPtrOutputWithContext(ctx context.Context) JsonInputSchemaMappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonInputSchemaMappingPtrOutput)
}

type JsonInputSchemaMappingOutput struct{ *pulumi.OutputState }

func (JsonInputSchemaMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonInputSchemaMapping)(nil)).Elem()
}

func (o JsonInputSchemaMappingOutput) ToJsonInputSchemaMappingOutput() JsonInputSchemaMappingOutput {
	return o
}

func (o JsonInputSchemaMappingOutput) ToJsonInputSchemaMappingOutputWithContext(ctx context.Context) JsonInputSchemaMappingOutput {
	return o
}

func (o JsonInputSchemaMappingOutput) ToJsonInputSchemaMappingPtrOutput() JsonInputSchemaMappingPtrOutput {
	return o.ToJsonInputSchemaMappingPtrOutputWithContext(context.Background())
}

func (o JsonInputSchemaMappingOutput) ToJsonInputSchemaMappingPtrOutputWithContext(ctx context.Context) JsonInputSchemaMappingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JsonInputSchemaMapping) *JsonInputSchemaMapping {
		return &v
	}).(JsonInputSchemaMappingPtrOutput)
}

func (o JsonInputSchemaMappingOutput) DataVersion() JsonFieldWithDefaultPtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMapping) *JsonFieldWithDefault { return v.DataVersion }).(JsonFieldWithDefaultPtrOutput)
}

func (o JsonInputSchemaMappingOutput) EventTime() JsonFieldPtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMapping) *JsonField { return v.EventTime }).(JsonFieldPtrOutput)
}

func (o JsonInputSchemaMappingOutput) EventType() JsonFieldWithDefaultPtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMapping) *JsonFieldWithDefault { return v.EventType }).(JsonFieldWithDefaultPtrOutput)
}

func (o JsonInputSchemaMappingOutput) Id() JsonFieldPtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMapping) *JsonField { return v.Id }).(JsonFieldPtrOutput)
}

func (o JsonInputSchemaMappingOutput) InputSchemaMappingType() pulumi.StringOutput {
	return o.ApplyT(func(v JsonInputSchemaMapping) string { return v.InputSchemaMappingType }).(pulumi.StringOutput)
}

func (o JsonInputSchemaMappingOutput) Subject() JsonFieldWithDefaultPtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMapping) *JsonFieldWithDefault { return v.Subject }).(JsonFieldWithDefaultPtrOutput)
}

func (o JsonInputSchemaMappingOutput) Topic() JsonFieldPtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMapping) *JsonField { return v.Topic }).(JsonFieldPtrOutput)
}

type JsonInputSchemaMappingPtrOutput struct{ *pulumi.OutputState }

func (JsonInputSchemaMappingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonInputSchemaMapping)(nil)).Elem()
}

func (o JsonInputSchemaMappingPtrOutput) ToJsonInputSchemaMappingPtrOutput() JsonInputSchemaMappingPtrOutput {
	return o
}

func (o JsonInputSchemaMappingPtrOutput) ToJsonInputSchemaMappingPtrOutputWithContext(ctx context.Context) JsonInputSchemaMappingPtrOutput {
	return o
}

func (o JsonInputSchemaMappingPtrOutput) Elem() JsonInputSchemaMappingOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) JsonInputSchemaMapping {
		if v != nil {
			return *v
		}
		var ret JsonInputSchemaMapping
		return ret
	}).(JsonInputSchemaMappingOutput)
}

func (o JsonInputSchemaMappingPtrOutput) DataVersion() JsonFieldWithDefaultPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) *JsonFieldWithDefault {
		if v == nil {
			return nil
		}
		return v.DataVersion
	}).(JsonFieldWithDefaultPtrOutput)
}

func (o JsonInputSchemaMappingPtrOutput) EventTime() JsonFieldPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) *JsonField {
		if v == nil {
			return nil
		}
		return v.EventTime
	}).(JsonFieldPtrOutput)
}

func (o JsonInputSchemaMappingPtrOutput) EventType() JsonFieldWithDefaultPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) *JsonFieldWithDefault {
		if v == nil {
			return nil
		}
		return v.EventType
	}).(JsonFieldWithDefaultPtrOutput)
}

func (o JsonInputSchemaMappingPtrOutput) Id() JsonFieldPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) *JsonField {
		if v == nil {
			return nil
		}
		return v.Id
	}).(JsonFieldPtrOutput)
}

func (o JsonInputSchemaMappingPtrOutput) InputSchemaMappingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) *string {
		if v == nil {
			return nil
		}
		return &v.InputSchemaMappingType
	}).(pulumi.StringPtrOutput)
}

func (o JsonInputSchemaMappingPtrOutput) Subject() JsonFieldWithDefaultPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) *JsonFieldWithDefault {
		if v == nil {
			return nil
		}
		return v.Subject
	}).(JsonFieldWithDefaultPtrOutput)
}

func (o JsonInputSchemaMappingPtrOutput) Topic() JsonFieldPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) *JsonField {
		if v == nil {
			return nil
		}
		return v.Topic
	}).(JsonFieldPtrOutput)
}

type JsonInputSchemaMappingResponse struct {
	DataVersion            *JsonFieldWithDefaultResponse `pulumi:"dataVersion"`
	EventTime              *JsonFieldResponse            `pulumi:"eventTime"`
	EventType              *JsonFieldWithDefaultResponse `pulumi:"eventType"`
	Id                     *JsonFieldResponse            `pulumi:"id"`
	InputSchemaMappingType string                        `pulumi:"inputSchemaMappingType"`
	Subject                *JsonFieldWithDefaultResponse `pulumi:"subject"`
	Topic                  *JsonFieldResponse            `pulumi:"topic"`
}





type JsonInputSchemaMappingResponseInput interface {
	pulumi.Input

	ToJsonInputSchemaMappingResponseOutput() JsonInputSchemaMappingResponseOutput
	ToJsonInputSchemaMappingResponseOutputWithContext(context.Context) JsonInputSchemaMappingResponseOutput
}

type JsonInputSchemaMappingResponseArgs struct {
	DataVersion            JsonFieldWithDefaultResponsePtrInput `pulumi:"dataVersion"`
	EventTime              JsonFieldResponsePtrInput            `pulumi:"eventTime"`
	EventType              JsonFieldWithDefaultResponsePtrInput `pulumi:"eventType"`
	Id                     JsonFieldResponsePtrInput            `pulumi:"id"`
	InputSchemaMappingType pulumi.StringInput                   `pulumi:"inputSchemaMappingType"`
	Subject                JsonFieldWithDefaultResponsePtrInput `pulumi:"subject"`
	Topic                  JsonFieldResponsePtrInput            `pulumi:"topic"`
}

func (JsonInputSchemaMappingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonInputSchemaMappingResponse)(nil)).Elem()
}

func (i JsonInputSchemaMappingResponseArgs) ToJsonInputSchemaMappingResponseOutput() JsonInputSchemaMappingResponseOutput {
	return i.ToJsonInputSchemaMappingResponseOutputWithContext(context.Background())
}

func (i JsonInputSchemaMappingResponseArgs) ToJsonInputSchemaMappingResponseOutputWithContext(ctx context.Context) JsonInputSchemaMappingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonInputSchemaMappingResponseOutput)
}

func (i JsonInputSchemaMappingResponseArgs) ToJsonInputSchemaMappingResponsePtrOutput() JsonInputSchemaMappingResponsePtrOutput {
	return i.ToJsonInputSchemaMappingResponsePtrOutputWithContext(context.Background())
}

func (i JsonInputSchemaMappingResponseArgs) ToJsonInputSchemaMappingResponsePtrOutputWithContext(ctx context.Context) JsonInputSchemaMappingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonInputSchemaMappingResponseOutput).ToJsonInputSchemaMappingResponsePtrOutputWithContext(ctx)
}









type JsonInputSchemaMappingResponsePtrInput interface {
	pulumi.Input

	ToJsonInputSchemaMappingResponsePtrOutput() JsonInputSchemaMappingResponsePtrOutput
	ToJsonInputSchemaMappingResponsePtrOutputWithContext(context.Context) JsonInputSchemaMappingResponsePtrOutput
}

type jsonInputSchemaMappingResponsePtrType JsonInputSchemaMappingResponseArgs

func JsonInputSchemaMappingResponsePtr(v *JsonInputSchemaMappingResponseArgs) JsonInputSchemaMappingResponsePtrInput {
	return (*jsonInputSchemaMappingResponsePtrType)(v)
}

func (*jsonInputSchemaMappingResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonInputSchemaMappingResponse)(nil)).Elem()
}

func (i *jsonInputSchemaMappingResponsePtrType) ToJsonInputSchemaMappingResponsePtrOutput() JsonInputSchemaMappingResponsePtrOutput {
	return i.ToJsonInputSchemaMappingResponsePtrOutputWithContext(context.Background())
}

func (i *jsonInputSchemaMappingResponsePtrType) ToJsonInputSchemaMappingResponsePtrOutputWithContext(ctx context.Context) JsonInputSchemaMappingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonInputSchemaMappingResponsePtrOutput)
}

type JsonInputSchemaMappingResponseOutput struct{ *pulumi.OutputState }

func (JsonInputSchemaMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonInputSchemaMappingResponse)(nil)).Elem()
}

func (o JsonInputSchemaMappingResponseOutput) ToJsonInputSchemaMappingResponseOutput() JsonInputSchemaMappingResponseOutput {
	return o
}

func (o JsonInputSchemaMappingResponseOutput) ToJsonInputSchemaMappingResponseOutputWithContext(ctx context.Context) JsonInputSchemaMappingResponseOutput {
	return o
}

func (o JsonInputSchemaMappingResponseOutput) ToJsonInputSchemaMappingResponsePtrOutput() JsonInputSchemaMappingResponsePtrOutput {
	return o.ToJsonInputSchemaMappingResponsePtrOutputWithContext(context.Background())
}

func (o JsonInputSchemaMappingResponseOutput) ToJsonInputSchemaMappingResponsePtrOutputWithContext(ctx context.Context) JsonInputSchemaMappingResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JsonInputSchemaMappingResponse) *JsonInputSchemaMappingResponse {
		return &v
	}).(JsonInputSchemaMappingResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponseOutput) DataVersion() JsonFieldWithDefaultResponsePtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMappingResponse) *JsonFieldWithDefaultResponse { return v.DataVersion }).(JsonFieldWithDefaultResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponseOutput) EventTime() JsonFieldResponsePtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMappingResponse) *JsonFieldResponse { return v.EventTime }).(JsonFieldResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponseOutput) EventType() JsonFieldWithDefaultResponsePtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMappingResponse) *JsonFieldWithDefaultResponse { return v.EventType }).(JsonFieldWithDefaultResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponseOutput) Id() JsonFieldResponsePtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMappingResponse) *JsonFieldResponse { return v.Id }).(JsonFieldResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponseOutput) InputSchemaMappingType() pulumi.StringOutput {
	return o.ApplyT(func(v JsonInputSchemaMappingResponse) string { return v.InputSchemaMappingType }).(pulumi.StringOutput)
}

func (o JsonInputSchemaMappingResponseOutput) Subject() JsonFieldWithDefaultResponsePtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMappingResponse) *JsonFieldWithDefaultResponse { return v.Subject }).(JsonFieldWithDefaultResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponseOutput) Topic() JsonFieldResponsePtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMappingResponse) *JsonFieldResponse { return v.Topic }).(JsonFieldResponsePtrOutput)
}

type JsonInputSchemaMappingResponsePtrOutput struct{ *pulumi.OutputState }

func (JsonInputSchemaMappingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonInputSchemaMappingResponse)(nil)).Elem()
}

func (o JsonInputSchemaMappingResponsePtrOutput) ToJsonInputSchemaMappingResponsePtrOutput() JsonInputSchemaMappingResponsePtrOutput {
	return o
}

func (o JsonInputSchemaMappingResponsePtrOutput) ToJsonInputSchemaMappingResponsePtrOutputWithContext(ctx context.Context) JsonInputSchemaMappingResponsePtrOutput {
	return o
}

func (o JsonInputSchemaMappingResponsePtrOutput) Elem() JsonInputSchemaMappingResponseOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) JsonInputSchemaMappingResponse {
		if v != nil {
			return *v
		}
		var ret JsonInputSchemaMappingResponse
		return ret
	}).(JsonInputSchemaMappingResponseOutput)
}

func (o JsonInputSchemaMappingResponsePtrOutput) DataVersion() JsonFieldWithDefaultResponsePtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) *JsonFieldWithDefaultResponse {
		if v == nil {
			return nil
		}
		return v.DataVersion
	}).(JsonFieldWithDefaultResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponsePtrOutput) EventTime() JsonFieldResponsePtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) *JsonFieldResponse {
		if v == nil {
			return nil
		}
		return v.EventTime
	}).(JsonFieldResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponsePtrOutput) EventType() JsonFieldWithDefaultResponsePtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) *JsonFieldWithDefaultResponse {
		if v == nil {
			return nil
		}
		return v.EventType
	}).(JsonFieldWithDefaultResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponsePtrOutput) Id() JsonFieldResponsePtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) *JsonFieldResponse {
		if v == nil {
			return nil
		}
		return v.Id
	}).(JsonFieldResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponsePtrOutput) InputSchemaMappingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) *string {
		if v == nil {
			return nil
		}
		return &v.InputSchemaMappingType
	}).(pulumi.StringPtrOutput)
}

func (o JsonInputSchemaMappingResponsePtrOutput) Subject() JsonFieldWithDefaultResponsePtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) *JsonFieldWithDefaultResponse {
		if v == nil {
			return nil
		}
		return v.Subject
	}).(JsonFieldWithDefaultResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponsePtrOutput) Topic() JsonFieldResponsePtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) *JsonFieldResponse {
		if v == nil {
			return nil
		}
		return v.Topic
	}).(JsonFieldResponsePtrOutput)
}

type NumberGreaterThanAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}





type NumberGreaterThanAdvancedFilterInput interface {
	pulumi.Input

	ToNumberGreaterThanAdvancedFilterOutput() NumberGreaterThanAdvancedFilterOutput
	ToNumberGreaterThanAdvancedFilterOutputWithContext(context.Context) NumberGreaterThanAdvancedFilterOutput
}

type NumberGreaterThanAdvancedFilterArgs struct {
	Key          pulumi.StringPtrInput  `pulumi:"key"`
	OperatorType pulumi.StringInput     `pulumi:"operatorType"`
	Value        pulumi.Float64PtrInput `pulumi:"value"`
}

func (NumberGreaterThanAdvancedFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberGreaterThanAdvancedFilter)(nil)).Elem()
}

func (i NumberGreaterThanAdvancedFilterArgs) ToNumberGreaterThanAdvancedFilterOutput() NumberGreaterThanAdvancedFilterOutput {
	return i.ToNumberGreaterThanAdvancedFilterOutputWithContext(context.Background())
}

func (i NumberGreaterThanAdvancedFilterArgs) ToNumberGreaterThanAdvancedFilterOutputWithContext(ctx context.Context) NumberGreaterThanAdvancedFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NumberGreaterThanAdvancedFilterOutput)
}

type NumberGreaterThanAdvancedFilterOutput struct{ *pulumi.OutputState }

func (NumberGreaterThanAdvancedFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberGreaterThanAdvancedFilter)(nil)).Elem()
}

func (o NumberGreaterThanAdvancedFilterOutput) ToNumberGreaterThanAdvancedFilterOutput() NumberGreaterThanAdvancedFilterOutput {
	return o
}

func (o NumberGreaterThanAdvancedFilterOutput) ToNumberGreaterThanAdvancedFilterOutputWithContext(ctx context.Context) NumberGreaterThanAdvancedFilterOutput {
	return o
}

func (o NumberGreaterThanAdvancedFilterOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NumberGreaterThanAdvancedFilter) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o NumberGreaterThanAdvancedFilterOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v NumberGreaterThanAdvancedFilter) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o NumberGreaterThanAdvancedFilterOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v NumberGreaterThanAdvancedFilter) *float64 { return v.Value }).(pulumi.Float64PtrOutput)
}

type NumberGreaterThanAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}





type NumberGreaterThanAdvancedFilterResponseInput interface {
	pulumi.Input

	ToNumberGreaterThanAdvancedFilterResponseOutput() NumberGreaterThanAdvancedFilterResponseOutput
	ToNumberGreaterThanAdvancedFilterResponseOutputWithContext(context.Context) NumberGreaterThanAdvancedFilterResponseOutput
}

type NumberGreaterThanAdvancedFilterResponseArgs struct {
	Key          pulumi.StringPtrInput  `pulumi:"key"`
	OperatorType pulumi.StringInput     `pulumi:"operatorType"`
	Value        pulumi.Float64PtrInput `pulumi:"value"`
}

func (NumberGreaterThanAdvancedFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberGreaterThanAdvancedFilterResponse)(nil)).Elem()
}

func (i NumberGreaterThanAdvancedFilterResponseArgs) ToNumberGreaterThanAdvancedFilterResponseOutput() NumberGreaterThanAdvancedFilterResponseOutput {
	return i.ToNumberGreaterThanAdvancedFilterResponseOutputWithContext(context.Background())
}

func (i NumberGreaterThanAdvancedFilterResponseArgs) ToNumberGreaterThanAdvancedFilterResponseOutputWithContext(ctx context.Context) NumberGreaterThanAdvancedFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NumberGreaterThanAdvancedFilterResponseOutput)
}

type NumberGreaterThanAdvancedFilterResponseOutput struct{ *pulumi.OutputState }

func (NumberGreaterThanAdvancedFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberGreaterThanAdvancedFilterResponse)(nil)).Elem()
}

func (o NumberGreaterThanAdvancedFilterResponseOutput) ToNumberGreaterThanAdvancedFilterResponseOutput() NumberGreaterThanAdvancedFilterResponseOutput {
	return o
}

func (o NumberGreaterThanAdvancedFilterResponseOutput) ToNumberGreaterThanAdvancedFilterResponseOutputWithContext(ctx context.Context) NumberGreaterThanAdvancedFilterResponseOutput {
	return o
}

func (o NumberGreaterThanAdvancedFilterResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NumberGreaterThanAdvancedFilterResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o NumberGreaterThanAdvancedFilterResponseOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v NumberGreaterThanAdvancedFilterResponse) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o NumberGreaterThanAdvancedFilterResponseOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v NumberGreaterThanAdvancedFilterResponse) *float64 { return v.Value }).(pulumi.Float64PtrOutput)
}

type NumberGreaterThanOrEqualsAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}





type NumberGreaterThanOrEqualsAdvancedFilterInput interface {
	pulumi.Input

	ToNumberGreaterThanOrEqualsAdvancedFilterOutput() NumberGreaterThanOrEqualsAdvancedFilterOutput
	ToNumberGreaterThanOrEqualsAdvancedFilterOutputWithContext(context.Context) NumberGreaterThanOrEqualsAdvancedFilterOutput
}

type NumberGreaterThanOrEqualsAdvancedFilterArgs struct {
	Key          pulumi.StringPtrInput  `pulumi:"key"`
	OperatorType pulumi.StringInput     `pulumi:"operatorType"`
	Value        pulumi.Float64PtrInput `pulumi:"value"`
}

func (NumberGreaterThanOrEqualsAdvancedFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberGreaterThanOrEqualsAdvancedFilter)(nil)).Elem()
}

func (i NumberGreaterThanOrEqualsAdvancedFilterArgs) ToNumberGreaterThanOrEqualsAdvancedFilterOutput() NumberGreaterThanOrEqualsAdvancedFilterOutput {
	return i.ToNumberGreaterThanOrEqualsAdvancedFilterOutputWithContext(context.Background())
}

func (i NumberGreaterThanOrEqualsAdvancedFilterArgs) ToNumberGreaterThanOrEqualsAdvancedFilterOutputWithContext(ctx context.Context) NumberGreaterThanOrEqualsAdvancedFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NumberGreaterThanOrEqualsAdvancedFilterOutput)
}

type NumberGreaterThanOrEqualsAdvancedFilterOutput struct{ *pulumi.OutputState }

func (NumberGreaterThanOrEqualsAdvancedFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberGreaterThanOrEqualsAdvancedFilter)(nil)).Elem()
}

func (o NumberGreaterThanOrEqualsAdvancedFilterOutput) ToNumberGreaterThanOrEqualsAdvancedFilterOutput() NumberGreaterThanOrEqualsAdvancedFilterOutput {
	return o
}

func (o NumberGreaterThanOrEqualsAdvancedFilterOutput) ToNumberGreaterThanOrEqualsAdvancedFilterOutputWithContext(ctx context.Context) NumberGreaterThanOrEqualsAdvancedFilterOutput {
	return o
}

func (o NumberGreaterThanOrEqualsAdvancedFilterOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NumberGreaterThanOrEqualsAdvancedFilter) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o NumberGreaterThanOrEqualsAdvancedFilterOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v NumberGreaterThanOrEqualsAdvancedFilter) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o NumberGreaterThanOrEqualsAdvancedFilterOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v NumberGreaterThanOrEqualsAdvancedFilter) *float64 { return v.Value }).(pulumi.Float64PtrOutput)
}

type NumberGreaterThanOrEqualsAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}





type NumberGreaterThanOrEqualsAdvancedFilterResponseInput interface {
	pulumi.Input

	ToNumberGreaterThanOrEqualsAdvancedFilterResponseOutput() NumberGreaterThanOrEqualsAdvancedFilterResponseOutput
	ToNumberGreaterThanOrEqualsAdvancedFilterResponseOutputWithContext(context.Context) NumberGreaterThanOrEqualsAdvancedFilterResponseOutput
}

type NumberGreaterThanOrEqualsAdvancedFilterResponseArgs struct {
	Key          pulumi.StringPtrInput  `pulumi:"key"`
	OperatorType pulumi.StringInput     `pulumi:"operatorType"`
	Value        pulumi.Float64PtrInput `pulumi:"value"`
}

func (NumberGreaterThanOrEqualsAdvancedFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberGreaterThanOrEqualsAdvancedFilterResponse)(nil)).Elem()
}

func (i NumberGreaterThanOrEqualsAdvancedFilterResponseArgs) ToNumberGreaterThanOrEqualsAdvancedFilterResponseOutput() NumberGreaterThanOrEqualsAdvancedFilterResponseOutput {
	return i.ToNumberGreaterThanOrEqualsAdvancedFilterResponseOutputWithContext(context.Background())
}

func (i NumberGreaterThanOrEqualsAdvancedFilterResponseArgs) ToNumberGreaterThanOrEqualsAdvancedFilterResponseOutputWithContext(ctx context.Context) NumberGreaterThanOrEqualsAdvancedFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NumberGreaterThanOrEqualsAdvancedFilterResponseOutput)
}

type NumberGreaterThanOrEqualsAdvancedFilterResponseOutput struct{ *pulumi.OutputState }

func (NumberGreaterThanOrEqualsAdvancedFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberGreaterThanOrEqualsAdvancedFilterResponse)(nil)).Elem()
}

func (o NumberGreaterThanOrEqualsAdvancedFilterResponseOutput) ToNumberGreaterThanOrEqualsAdvancedFilterResponseOutput() NumberGreaterThanOrEqualsAdvancedFilterResponseOutput {
	return o
}

func (o NumberGreaterThanOrEqualsAdvancedFilterResponseOutput) ToNumberGreaterThanOrEqualsAdvancedFilterResponseOutputWithContext(ctx context.Context) NumberGreaterThanOrEqualsAdvancedFilterResponseOutput {
	return o
}

func (o NumberGreaterThanOrEqualsAdvancedFilterResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NumberGreaterThanOrEqualsAdvancedFilterResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o NumberGreaterThanOrEqualsAdvancedFilterResponseOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v NumberGreaterThanOrEqualsAdvancedFilterResponse) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o NumberGreaterThanOrEqualsAdvancedFilterResponseOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v NumberGreaterThanOrEqualsAdvancedFilterResponse) *float64 { return v.Value }).(pulumi.Float64PtrOutput)
}

type NumberInAdvancedFilter struct {
	Key          *string   `pulumi:"key"`
	OperatorType string    `pulumi:"operatorType"`
	Values       []float64 `pulumi:"values"`
}





type NumberInAdvancedFilterInput interface {
	pulumi.Input

	ToNumberInAdvancedFilterOutput() NumberInAdvancedFilterOutput
	ToNumberInAdvancedFilterOutputWithContext(context.Context) NumberInAdvancedFilterOutput
}

type NumberInAdvancedFilterArgs struct {
	Key          pulumi.StringPtrInput    `pulumi:"key"`
	OperatorType pulumi.StringInput       `pulumi:"operatorType"`
	Values       pulumi.Float64ArrayInput `pulumi:"values"`
}

func (NumberInAdvancedFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberInAdvancedFilter)(nil)).Elem()
}

func (i NumberInAdvancedFilterArgs) ToNumberInAdvancedFilterOutput() NumberInAdvancedFilterOutput {
	return i.ToNumberInAdvancedFilterOutputWithContext(context.Background())
}

func (i NumberInAdvancedFilterArgs) ToNumberInAdvancedFilterOutputWithContext(ctx context.Context) NumberInAdvancedFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NumberInAdvancedFilterOutput)
}

type NumberInAdvancedFilterOutput struct{ *pulumi.OutputState }

func (NumberInAdvancedFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberInAdvancedFilter)(nil)).Elem()
}

func (o NumberInAdvancedFilterOutput) ToNumberInAdvancedFilterOutput() NumberInAdvancedFilterOutput {
	return o
}

func (o NumberInAdvancedFilterOutput) ToNumberInAdvancedFilterOutputWithContext(ctx context.Context) NumberInAdvancedFilterOutput {
	return o
}

func (o NumberInAdvancedFilterOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NumberInAdvancedFilter) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o NumberInAdvancedFilterOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v NumberInAdvancedFilter) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o NumberInAdvancedFilterOutput) Values() pulumi.Float64ArrayOutput {
	return o.ApplyT(func(v NumberInAdvancedFilter) []float64 { return v.Values }).(pulumi.Float64ArrayOutput)
}

type NumberInAdvancedFilterResponse struct {
	Key          *string   `pulumi:"key"`
	OperatorType string    `pulumi:"operatorType"`
	Values       []float64 `pulumi:"values"`
}





type NumberInAdvancedFilterResponseInput interface {
	pulumi.Input

	ToNumberInAdvancedFilterResponseOutput() NumberInAdvancedFilterResponseOutput
	ToNumberInAdvancedFilterResponseOutputWithContext(context.Context) NumberInAdvancedFilterResponseOutput
}

type NumberInAdvancedFilterResponseArgs struct {
	Key          pulumi.StringPtrInput    `pulumi:"key"`
	OperatorType pulumi.StringInput       `pulumi:"operatorType"`
	Values       pulumi.Float64ArrayInput `pulumi:"values"`
}

func (NumberInAdvancedFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberInAdvancedFilterResponse)(nil)).Elem()
}

func (i NumberInAdvancedFilterResponseArgs) ToNumberInAdvancedFilterResponseOutput() NumberInAdvancedFilterResponseOutput {
	return i.ToNumberInAdvancedFilterResponseOutputWithContext(context.Background())
}

func (i NumberInAdvancedFilterResponseArgs) ToNumberInAdvancedFilterResponseOutputWithContext(ctx context.Context) NumberInAdvancedFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NumberInAdvancedFilterResponseOutput)
}

type NumberInAdvancedFilterResponseOutput struct{ *pulumi.OutputState }

func (NumberInAdvancedFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberInAdvancedFilterResponse)(nil)).Elem()
}

func (o NumberInAdvancedFilterResponseOutput) ToNumberInAdvancedFilterResponseOutput() NumberInAdvancedFilterResponseOutput {
	return o
}

func (o NumberInAdvancedFilterResponseOutput) ToNumberInAdvancedFilterResponseOutputWithContext(ctx context.Context) NumberInAdvancedFilterResponseOutput {
	return o
}

func (o NumberInAdvancedFilterResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NumberInAdvancedFilterResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o NumberInAdvancedFilterResponseOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v NumberInAdvancedFilterResponse) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o NumberInAdvancedFilterResponseOutput) Values() pulumi.Float64ArrayOutput {
	return o.ApplyT(func(v NumberInAdvancedFilterResponse) []float64 { return v.Values }).(pulumi.Float64ArrayOutput)
}

type NumberLessThanAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}





type NumberLessThanAdvancedFilterInput interface {
	pulumi.Input

	ToNumberLessThanAdvancedFilterOutput() NumberLessThanAdvancedFilterOutput
	ToNumberLessThanAdvancedFilterOutputWithContext(context.Context) NumberLessThanAdvancedFilterOutput
}

type NumberLessThanAdvancedFilterArgs struct {
	Key          pulumi.StringPtrInput  `pulumi:"key"`
	OperatorType pulumi.StringInput     `pulumi:"operatorType"`
	Value        pulumi.Float64PtrInput `pulumi:"value"`
}

func (NumberLessThanAdvancedFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberLessThanAdvancedFilter)(nil)).Elem()
}

func (i NumberLessThanAdvancedFilterArgs) ToNumberLessThanAdvancedFilterOutput() NumberLessThanAdvancedFilterOutput {
	return i.ToNumberLessThanAdvancedFilterOutputWithContext(context.Background())
}

func (i NumberLessThanAdvancedFilterArgs) ToNumberLessThanAdvancedFilterOutputWithContext(ctx context.Context) NumberLessThanAdvancedFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NumberLessThanAdvancedFilterOutput)
}

type NumberLessThanAdvancedFilterOutput struct{ *pulumi.OutputState }

func (NumberLessThanAdvancedFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberLessThanAdvancedFilter)(nil)).Elem()
}

func (o NumberLessThanAdvancedFilterOutput) ToNumberLessThanAdvancedFilterOutput() NumberLessThanAdvancedFilterOutput {
	return o
}

func (o NumberLessThanAdvancedFilterOutput) ToNumberLessThanAdvancedFilterOutputWithContext(ctx context.Context) NumberLessThanAdvancedFilterOutput {
	return o
}

func (o NumberLessThanAdvancedFilterOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NumberLessThanAdvancedFilter) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o NumberLessThanAdvancedFilterOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v NumberLessThanAdvancedFilter) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o NumberLessThanAdvancedFilterOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v NumberLessThanAdvancedFilter) *float64 { return v.Value }).(pulumi.Float64PtrOutput)
}

type NumberLessThanAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}





type NumberLessThanAdvancedFilterResponseInput interface {
	pulumi.Input

	ToNumberLessThanAdvancedFilterResponseOutput() NumberLessThanAdvancedFilterResponseOutput
	ToNumberLessThanAdvancedFilterResponseOutputWithContext(context.Context) NumberLessThanAdvancedFilterResponseOutput
}

type NumberLessThanAdvancedFilterResponseArgs struct {
	Key          pulumi.StringPtrInput  `pulumi:"key"`
	OperatorType pulumi.StringInput     `pulumi:"operatorType"`
	Value        pulumi.Float64PtrInput `pulumi:"value"`
}

func (NumberLessThanAdvancedFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberLessThanAdvancedFilterResponse)(nil)).Elem()
}

func (i NumberLessThanAdvancedFilterResponseArgs) ToNumberLessThanAdvancedFilterResponseOutput() NumberLessThanAdvancedFilterResponseOutput {
	return i.ToNumberLessThanAdvancedFilterResponseOutputWithContext(context.Background())
}

func (i NumberLessThanAdvancedFilterResponseArgs) ToNumberLessThanAdvancedFilterResponseOutputWithContext(ctx context.Context) NumberLessThanAdvancedFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NumberLessThanAdvancedFilterResponseOutput)
}

type NumberLessThanAdvancedFilterResponseOutput struct{ *pulumi.OutputState }

func (NumberLessThanAdvancedFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberLessThanAdvancedFilterResponse)(nil)).Elem()
}

func (o NumberLessThanAdvancedFilterResponseOutput) ToNumberLessThanAdvancedFilterResponseOutput() NumberLessThanAdvancedFilterResponseOutput {
	return o
}

func (o NumberLessThanAdvancedFilterResponseOutput) ToNumberLessThanAdvancedFilterResponseOutputWithContext(ctx context.Context) NumberLessThanAdvancedFilterResponseOutput {
	return o
}

func (o NumberLessThanAdvancedFilterResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NumberLessThanAdvancedFilterResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o NumberLessThanAdvancedFilterResponseOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v NumberLessThanAdvancedFilterResponse) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o NumberLessThanAdvancedFilterResponseOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v NumberLessThanAdvancedFilterResponse) *float64 { return v.Value }).(pulumi.Float64PtrOutput)
}

type NumberLessThanOrEqualsAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}





type NumberLessThanOrEqualsAdvancedFilterInput interface {
	pulumi.Input

	ToNumberLessThanOrEqualsAdvancedFilterOutput() NumberLessThanOrEqualsAdvancedFilterOutput
	ToNumberLessThanOrEqualsAdvancedFilterOutputWithContext(context.Context) NumberLessThanOrEqualsAdvancedFilterOutput
}

type NumberLessThanOrEqualsAdvancedFilterArgs struct {
	Key          pulumi.StringPtrInput  `pulumi:"key"`
	OperatorType pulumi.StringInput     `pulumi:"operatorType"`
	Value        pulumi.Float64PtrInput `pulumi:"value"`
}

func (NumberLessThanOrEqualsAdvancedFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberLessThanOrEqualsAdvancedFilter)(nil)).Elem()
}

func (i NumberLessThanOrEqualsAdvancedFilterArgs) ToNumberLessThanOrEqualsAdvancedFilterOutput() NumberLessThanOrEqualsAdvancedFilterOutput {
	return i.ToNumberLessThanOrEqualsAdvancedFilterOutputWithContext(context.Background())
}

func (i NumberLessThanOrEqualsAdvancedFilterArgs) ToNumberLessThanOrEqualsAdvancedFilterOutputWithContext(ctx context.Context) NumberLessThanOrEqualsAdvancedFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NumberLessThanOrEqualsAdvancedFilterOutput)
}

type NumberLessThanOrEqualsAdvancedFilterOutput struct{ *pulumi.OutputState }

func (NumberLessThanOrEqualsAdvancedFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberLessThanOrEqualsAdvancedFilter)(nil)).Elem()
}

func (o NumberLessThanOrEqualsAdvancedFilterOutput) ToNumberLessThanOrEqualsAdvancedFilterOutput() NumberLessThanOrEqualsAdvancedFilterOutput {
	return o
}

func (o NumberLessThanOrEqualsAdvancedFilterOutput) ToNumberLessThanOrEqualsAdvancedFilterOutputWithContext(ctx context.Context) NumberLessThanOrEqualsAdvancedFilterOutput {
	return o
}

func (o NumberLessThanOrEqualsAdvancedFilterOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NumberLessThanOrEqualsAdvancedFilter) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o NumberLessThanOrEqualsAdvancedFilterOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v NumberLessThanOrEqualsAdvancedFilter) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o NumberLessThanOrEqualsAdvancedFilterOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v NumberLessThanOrEqualsAdvancedFilter) *float64 { return v.Value }).(pulumi.Float64PtrOutput)
}

type NumberLessThanOrEqualsAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}





type NumberLessThanOrEqualsAdvancedFilterResponseInput interface {
	pulumi.Input

	ToNumberLessThanOrEqualsAdvancedFilterResponseOutput() NumberLessThanOrEqualsAdvancedFilterResponseOutput
	ToNumberLessThanOrEqualsAdvancedFilterResponseOutputWithContext(context.Context) NumberLessThanOrEqualsAdvancedFilterResponseOutput
}

type NumberLessThanOrEqualsAdvancedFilterResponseArgs struct {
	Key          pulumi.StringPtrInput  `pulumi:"key"`
	OperatorType pulumi.StringInput     `pulumi:"operatorType"`
	Value        pulumi.Float64PtrInput `pulumi:"value"`
}

func (NumberLessThanOrEqualsAdvancedFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberLessThanOrEqualsAdvancedFilterResponse)(nil)).Elem()
}

func (i NumberLessThanOrEqualsAdvancedFilterResponseArgs) ToNumberLessThanOrEqualsAdvancedFilterResponseOutput() NumberLessThanOrEqualsAdvancedFilterResponseOutput {
	return i.ToNumberLessThanOrEqualsAdvancedFilterResponseOutputWithContext(context.Background())
}

func (i NumberLessThanOrEqualsAdvancedFilterResponseArgs) ToNumberLessThanOrEqualsAdvancedFilterResponseOutputWithContext(ctx context.Context) NumberLessThanOrEqualsAdvancedFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NumberLessThanOrEqualsAdvancedFilterResponseOutput)
}

type NumberLessThanOrEqualsAdvancedFilterResponseOutput struct{ *pulumi.OutputState }

func (NumberLessThanOrEqualsAdvancedFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberLessThanOrEqualsAdvancedFilterResponse)(nil)).Elem()
}

func (o NumberLessThanOrEqualsAdvancedFilterResponseOutput) ToNumberLessThanOrEqualsAdvancedFilterResponseOutput() NumberLessThanOrEqualsAdvancedFilterResponseOutput {
	return o
}

func (o NumberLessThanOrEqualsAdvancedFilterResponseOutput) ToNumberLessThanOrEqualsAdvancedFilterResponseOutputWithContext(ctx context.Context) NumberLessThanOrEqualsAdvancedFilterResponseOutput {
	return o
}

func (o NumberLessThanOrEqualsAdvancedFilterResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NumberLessThanOrEqualsAdvancedFilterResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o NumberLessThanOrEqualsAdvancedFilterResponseOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v NumberLessThanOrEqualsAdvancedFilterResponse) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o NumberLessThanOrEqualsAdvancedFilterResponseOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v NumberLessThanOrEqualsAdvancedFilterResponse) *float64 { return v.Value }).(pulumi.Float64PtrOutput)
}

type NumberNotInAdvancedFilter struct {
	Key          *string   `pulumi:"key"`
	OperatorType string    `pulumi:"operatorType"`
	Values       []float64 `pulumi:"values"`
}





type NumberNotInAdvancedFilterInput interface {
	pulumi.Input

	ToNumberNotInAdvancedFilterOutput() NumberNotInAdvancedFilterOutput
	ToNumberNotInAdvancedFilterOutputWithContext(context.Context) NumberNotInAdvancedFilterOutput
}

type NumberNotInAdvancedFilterArgs struct {
	Key          pulumi.StringPtrInput    `pulumi:"key"`
	OperatorType pulumi.StringInput       `pulumi:"operatorType"`
	Values       pulumi.Float64ArrayInput `pulumi:"values"`
}

func (NumberNotInAdvancedFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberNotInAdvancedFilter)(nil)).Elem()
}

func (i NumberNotInAdvancedFilterArgs) ToNumberNotInAdvancedFilterOutput() NumberNotInAdvancedFilterOutput {
	return i.ToNumberNotInAdvancedFilterOutputWithContext(context.Background())
}

func (i NumberNotInAdvancedFilterArgs) ToNumberNotInAdvancedFilterOutputWithContext(ctx context.Context) NumberNotInAdvancedFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NumberNotInAdvancedFilterOutput)
}

type NumberNotInAdvancedFilterOutput struct{ *pulumi.OutputState }

func (NumberNotInAdvancedFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberNotInAdvancedFilter)(nil)).Elem()
}

func (o NumberNotInAdvancedFilterOutput) ToNumberNotInAdvancedFilterOutput() NumberNotInAdvancedFilterOutput {
	return o
}

func (o NumberNotInAdvancedFilterOutput) ToNumberNotInAdvancedFilterOutputWithContext(ctx context.Context) NumberNotInAdvancedFilterOutput {
	return o
}

func (o NumberNotInAdvancedFilterOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NumberNotInAdvancedFilter) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o NumberNotInAdvancedFilterOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v NumberNotInAdvancedFilter) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o NumberNotInAdvancedFilterOutput) Values() pulumi.Float64ArrayOutput {
	return o.ApplyT(func(v NumberNotInAdvancedFilter) []float64 { return v.Values }).(pulumi.Float64ArrayOutput)
}

type NumberNotInAdvancedFilterResponse struct {
	Key          *string   `pulumi:"key"`
	OperatorType string    `pulumi:"operatorType"`
	Values       []float64 `pulumi:"values"`
}





type NumberNotInAdvancedFilterResponseInput interface {
	pulumi.Input

	ToNumberNotInAdvancedFilterResponseOutput() NumberNotInAdvancedFilterResponseOutput
	ToNumberNotInAdvancedFilterResponseOutputWithContext(context.Context) NumberNotInAdvancedFilterResponseOutput
}

type NumberNotInAdvancedFilterResponseArgs struct {
	Key          pulumi.StringPtrInput    `pulumi:"key"`
	OperatorType pulumi.StringInput       `pulumi:"operatorType"`
	Values       pulumi.Float64ArrayInput `pulumi:"values"`
}

func (NumberNotInAdvancedFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberNotInAdvancedFilterResponse)(nil)).Elem()
}

func (i NumberNotInAdvancedFilterResponseArgs) ToNumberNotInAdvancedFilterResponseOutput() NumberNotInAdvancedFilterResponseOutput {
	return i.ToNumberNotInAdvancedFilterResponseOutputWithContext(context.Background())
}

func (i NumberNotInAdvancedFilterResponseArgs) ToNumberNotInAdvancedFilterResponseOutputWithContext(ctx context.Context) NumberNotInAdvancedFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NumberNotInAdvancedFilterResponseOutput)
}

type NumberNotInAdvancedFilterResponseOutput struct{ *pulumi.OutputState }

func (NumberNotInAdvancedFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NumberNotInAdvancedFilterResponse)(nil)).Elem()
}

func (o NumberNotInAdvancedFilterResponseOutput) ToNumberNotInAdvancedFilterResponseOutput() NumberNotInAdvancedFilterResponseOutput {
	return o
}

func (o NumberNotInAdvancedFilterResponseOutput) ToNumberNotInAdvancedFilterResponseOutputWithContext(ctx context.Context) NumberNotInAdvancedFilterResponseOutput {
	return o
}

func (o NumberNotInAdvancedFilterResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NumberNotInAdvancedFilterResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o NumberNotInAdvancedFilterResponseOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v NumberNotInAdvancedFilterResponse) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o NumberNotInAdvancedFilterResponseOutput) Values() pulumi.Float64ArrayOutput {
	return o.ApplyT(func(v NumberNotInAdvancedFilterResponse) []float64 { return v.Values }).(pulumi.Float64ArrayOutput)
}

type PrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(context.Context) PrivateEndpointOutput
}

type PrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput).ToPrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput
	ToPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointPtrOutput
}

type privateEndpointPtrType PrivateEndpointArgs

func PrivateEndpointPtr(v *PrivateEndpointArgs) PrivateEndpointPtrInput {
	return (*privateEndpointPtrType)(v)
}

func (*privateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPtrOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpoint) *PrivateEndpoint {
		return &v
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) Elem() PrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpoint
		return ret
	}).(PrivateEndpointOutput)
}

func (o PrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	GroupIds                          []string                 `pulumi:"groupIds"`
	Id                                string                   `pulumi:"id"`
	Name                              string                   `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *ConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 *string                  `pulumi:"provisioningState"`
	Type                              string                   `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	GroupIds                          pulumi.StringArrayInput         `pulumi:"groupIds"`
	Id                                pulumi.StringInput              `pulumi:"id"`
	Name                              pulumi.StringInput              `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrInput `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState ConnectionStateResponsePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringPtrInput           `pulumi:"provisioningState"`
	Type                              pulumi.StringInput              `pulumi:"type"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}





type PrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput
	ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionResponseArrayOutput
}

type PrivateEndpointConnectionResponseArray []PrivateEndpointConnectionResponseInput

func (PrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return i.ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseArrayOutput)
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() ConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *ConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(ConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type RetryPolicy struct {
	EventTimeToLiveInMinutes *int `pulumi:"eventTimeToLiveInMinutes"`
	MaxDeliveryAttempts      *int `pulumi:"maxDeliveryAttempts"`
}





type RetryPolicyInput interface {
	pulumi.Input

	ToRetryPolicyOutput() RetryPolicyOutput
	ToRetryPolicyOutputWithContext(context.Context) RetryPolicyOutput
}

type RetryPolicyArgs struct {
	EventTimeToLiveInMinutes pulumi.IntPtrInput `pulumi:"eventTimeToLiveInMinutes"`
	MaxDeliveryAttempts      pulumi.IntPtrInput `pulumi:"maxDeliveryAttempts"`
}

func (RetryPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetryPolicy)(nil)).Elem()
}

func (i RetryPolicyArgs) ToRetryPolicyOutput() RetryPolicyOutput {
	return i.ToRetryPolicyOutputWithContext(context.Background())
}

func (i RetryPolicyArgs) ToRetryPolicyOutputWithContext(ctx context.Context) RetryPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyOutput)
}

func (i RetryPolicyArgs) ToRetryPolicyPtrOutput() RetryPolicyPtrOutput {
	return i.ToRetryPolicyPtrOutputWithContext(context.Background())
}

func (i RetryPolicyArgs) ToRetryPolicyPtrOutputWithContext(ctx context.Context) RetryPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyOutput).ToRetryPolicyPtrOutputWithContext(ctx)
}









type RetryPolicyPtrInput interface {
	pulumi.Input

	ToRetryPolicyPtrOutput() RetryPolicyPtrOutput
	ToRetryPolicyPtrOutputWithContext(context.Context) RetryPolicyPtrOutput
}

type retryPolicyPtrType RetryPolicyArgs

func RetryPolicyPtr(v *RetryPolicyArgs) RetryPolicyPtrInput {
	return (*retryPolicyPtrType)(v)
}

func (*retryPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetryPolicy)(nil)).Elem()
}

func (i *retryPolicyPtrType) ToRetryPolicyPtrOutput() RetryPolicyPtrOutput {
	return i.ToRetryPolicyPtrOutputWithContext(context.Background())
}

func (i *retryPolicyPtrType) ToRetryPolicyPtrOutputWithContext(ctx context.Context) RetryPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyPtrOutput)
}

type RetryPolicyOutput struct{ *pulumi.OutputState }

func (RetryPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetryPolicy)(nil)).Elem()
}

func (o RetryPolicyOutput) ToRetryPolicyOutput() RetryPolicyOutput {
	return o
}

func (o RetryPolicyOutput) ToRetryPolicyOutputWithContext(ctx context.Context) RetryPolicyOutput {
	return o
}

func (o RetryPolicyOutput) ToRetryPolicyPtrOutput() RetryPolicyPtrOutput {
	return o.ToRetryPolicyPtrOutputWithContext(context.Background())
}

func (o RetryPolicyOutput) ToRetryPolicyPtrOutputWithContext(ctx context.Context) RetryPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetryPolicy) *RetryPolicy {
		return &v
	}).(RetryPolicyPtrOutput)
}

func (o RetryPolicyOutput) EventTimeToLiveInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetryPolicy) *int { return v.EventTimeToLiveInMinutes }).(pulumi.IntPtrOutput)
}

func (o RetryPolicyOutput) MaxDeliveryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetryPolicy) *int { return v.MaxDeliveryAttempts }).(pulumi.IntPtrOutput)
}

type RetryPolicyPtrOutput struct{ *pulumi.OutputState }

func (RetryPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetryPolicy)(nil)).Elem()
}

func (o RetryPolicyPtrOutput) ToRetryPolicyPtrOutput() RetryPolicyPtrOutput {
	return o
}

func (o RetryPolicyPtrOutput) ToRetryPolicyPtrOutputWithContext(ctx context.Context) RetryPolicyPtrOutput {
	return o
}

func (o RetryPolicyPtrOutput) Elem() RetryPolicyOutput {
	return o.ApplyT(func(v *RetryPolicy) RetryPolicy {
		if v != nil {
			return *v
		}
		var ret RetryPolicy
		return ret
	}).(RetryPolicyOutput)
}

func (o RetryPolicyPtrOutput) EventTimeToLiveInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetryPolicy) *int {
		if v == nil {
			return nil
		}
		return v.EventTimeToLiveInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o RetryPolicyPtrOutput) MaxDeliveryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetryPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxDeliveryAttempts
	}).(pulumi.IntPtrOutput)
}

type RetryPolicyResponse struct {
	EventTimeToLiveInMinutes *int `pulumi:"eventTimeToLiveInMinutes"`
	MaxDeliveryAttempts      *int `pulumi:"maxDeliveryAttempts"`
}





type RetryPolicyResponseInput interface {
	pulumi.Input

	ToRetryPolicyResponseOutput() RetryPolicyResponseOutput
	ToRetryPolicyResponseOutputWithContext(context.Context) RetryPolicyResponseOutput
}

type RetryPolicyResponseArgs struct {
	EventTimeToLiveInMinutes pulumi.IntPtrInput `pulumi:"eventTimeToLiveInMinutes"`
	MaxDeliveryAttempts      pulumi.IntPtrInput `pulumi:"maxDeliveryAttempts"`
}

func (RetryPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetryPolicyResponse)(nil)).Elem()
}

func (i RetryPolicyResponseArgs) ToRetryPolicyResponseOutput() RetryPolicyResponseOutput {
	return i.ToRetryPolicyResponseOutputWithContext(context.Background())
}

func (i RetryPolicyResponseArgs) ToRetryPolicyResponseOutputWithContext(ctx context.Context) RetryPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyResponseOutput)
}

func (i RetryPolicyResponseArgs) ToRetryPolicyResponsePtrOutput() RetryPolicyResponsePtrOutput {
	return i.ToRetryPolicyResponsePtrOutputWithContext(context.Background())
}

func (i RetryPolicyResponseArgs) ToRetryPolicyResponsePtrOutputWithContext(ctx context.Context) RetryPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyResponseOutput).ToRetryPolicyResponsePtrOutputWithContext(ctx)
}









type RetryPolicyResponsePtrInput interface {
	pulumi.Input

	ToRetryPolicyResponsePtrOutput() RetryPolicyResponsePtrOutput
	ToRetryPolicyResponsePtrOutputWithContext(context.Context) RetryPolicyResponsePtrOutput
}

type retryPolicyResponsePtrType RetryPolicyResponseArgs

func RetryPolicyResponsePtr(v *RetryPolicyResponseArgs) RetryPolicyResponsePtrInput {
	return (*retryPolicyResponsePtrType)(v)
}

func (*retryPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetryPolicyResponse)(nil)).Elem()
}

func (i *retryPolicyResponsePtrType) ToRetryPolicyResponsePtrOutput() RetryPolicyResponsePtrOutput {
	return i.ToRetryPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *retryPolicyResponsePtrType) ToRetryPolicyResponsePtrOutputWithContext(ctx context.Context) RetryPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyResponsePtrOutput)
}

type RetryPolicyResponseOutput struct{ *pulumi.OutputState }

func (RetryPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetryPolicyResponse)(nil)).Elem()
}

func (o RetryPolicyResponseOutput) ToRetryPolicyResponseOutput() RetryPolicyResponseOutput {
	return o
}

func (o RetryPolicyResponseOutput) ToRetryPolicyResponseOutputWithContext(ctx context.Context) RetryPolicyResponseOutput {
	return o
}

func (o RetryPolicyResponseOutput) ToRetryPolicyResponsePtrOutput() RetryPolicyResponsePtrOutput {
	return o.ToRetryPolicyResponsePtrOutputWithContext(context.Background())
}

func (o RetryPolicyResponseOutput) ToRetryPolicyResponsePtrOutputWithContext(ctx context.Context) RetryPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetryPolicyResponse) *RetryPolicyResponse {
		return &v
	}).(RetryPolicyResponsePtrOutput)
}

func (o RetryPolicyResponseOutput) EventTimeToLiveInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetryPolicyResponse) *int { return v.EventTimeToLiveInMinutes }).(pulumi.IntPtrOutput)
}

func (o RetryPolicyResponseOutput) MaxDeliveryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetryPolicyResponse) *int { return v.MaxDeliveryAttempts }).(pulumi.IntPtrOutput)
}

type RetryPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (RetryPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetryPolicyResponse)(nil)).Elem()
}

func (o RetryPolicyResponsePtrOutput) ToRetryPolicyResponsePtrOutput() RetryPolicyResponsePtrOutput {
	return o
}

func (o RetryPolicyResponsePtrOutput) ToRetryPolicyResponsePtrOutputWithContext(ctx context.Context) RetryPolicyResponsePtrOutput {
	return o
}

func (o RetryPolicyResponsePtrOutput) Elem() RetryPolicyResponseOutput {
	return o.ApplyT(func(v *RetryPolicyResponse) RetryPolicyResponse {
		if v != nil {
			return *v
		}
		var ret RetryPolicyResponse
		return ret
	}).(RetryPolicyResponseOutput)
}

func (o RetryPolicyResponsePtrOutput) EventTimeToLiveInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetryPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.EventTimeToLiveInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o RetryPolicyResponsePtrOutput) MaxDeliveryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetryPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxDeliveryAttempts
	}).(pulumi.IntPtrOutput)
}

type ServiceBusQueueEventSubscriptionDestination struct {
	DeliveryAttributeMappings []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType              string        `pulumi:"endpointType"`
	ResourceId                *string       `pulumi:"resourceId"`
}





type ServiceBusQueueEventSubscriptionDestinationInput interface {
	pulumi.Input

	ToServiceBusQueueEventSubscriptionDestinationOutput() ServiceBusQueueEventSubscriptionDestinationOutput
	ToServiceBusQueueEventSubscriptionDestinationOutputWithContext(context.Context) ServiceBusQueueEventSubscriptionDestinationOutput
}

type ServiceBusQueueEventSubscriptionDestinationArgs struct {
	DeliveryAttributeMappings pulumi.ArrayInput     `pulumi:"deliveryAttributeMappings"`
	EndpointType              pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId                pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (ServiceBusQueueEventSubscriptionDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusQueueEventSubscriptionDestination)(nil)).Elem()
}

func (i ServiceBusQueueEventSubscriptionDestinationArgs) ToServiceBusQueueEventSubscriptionDestinationOutput() ServiceBusQueueEventSubscriptionDestinationOutput {
	return i.ToServiceBusQueueEventSubscriptionDestinationOutputWithContext(context.Background())
}

func (i ServiceBusQueueEventSubscriptionDestinationArgs) ToServiceBusQueueEventSubscriptionDestinationOutputWithContext(ctx context.Context) ServiceBusQueueEventSubscriptionDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusQueueEventSubscriptionDestinationOutput)
}

type ServiceBusQueueEventSubscriptionDestinationOutput struct{ *pulumi.OutputState }

func (ServiceBusQueueEventSubscriptionDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusQueueEventSubscriptionDestination)(nil)).Elem()
}

func (o ServiceBusQueueEventSubscriptionDestinationOutput) ToServiceBusQueueEventSubscriptionDestinationOutput() ServiceBusQueueEventSubscriptionDestinationOutput {
	return o
}

func (o ServiceBusQueueEventSubscriptionDestinationOutput) ToServiceBusQueueEventSubscriptionDestinationOutputWithContext(ctx context.Context) ServiceBusQueueEventSubscriptionDestinationOutput {
	return o
}

func (o ServiceBusQueueEventSubscriptionDestinationOutput) DeliveryAttributeMappings() pulumi.ArrayOutput {
	return o.ApplyT(func(v ServiceBusQueueEventSubscriptionDestination) []interface{} { return v.DeliveryAttributeMappings }).(pulumi.ArrayOutput)
}

func (o ServiceBusQueueEventSubscriptionDestinationOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBusQueueEventSubscriptionDestination) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o ServiceBusQueueEventSubscriptionDestinationOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueEventSubscriptionDestination) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ServiceBusQueueEventSubscriptionDestinationResponse struct {
	DeliveryAttributeMappings []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType              string        `pulumi:"endpointType"`
	ResourceId                *string       `pulumi:"resourceId"`
}





type ServiceBusQueueEventSubscriptionDestinationResponseInput interface {
	pulumi.Input

	ToServiceBusQueueEventSubscriptionDestinationResponseOutput() ServiceBusQueueEventSubscriptionDestinationResponseOutput
	ToServiceBusQueueEventSubscriptionDestinationResponseOutputWithContext(context.Context) ServiceBusQueueEventSubscriptionDestinationResponseOutput
}

type ServiceBusQueueEventSubscriptionDestinationResponseArgs struct {
	DeliveryAttributeMappings pulumi.ArrayInput     `pulumi:"deliveryAttributeMappings"`
	EndpointType              pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId                pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (ServiceBusQueueEventSubscriptionDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusQueueEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (i ServiceBusQueueEventSubscriptionDestinationResponseArgs) ToServiceBusQueueEventSubscriptionDestinationResponseOutput() ServiceBusQueueEventSubscriptionDestinationResponseOutput {
	return i.ToServiceBusQueueEventSubscriptionDestinationResponseOutputWithContext(context.Background())
}

func (i ServiceBusQueueEventSubscriptionDestinationResponseArgs) ToServiceBusQueueEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) ServiceBusQueueEventSubscriptionDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusQueueEventSubscriptionDestinationResponseOutput)
}

type ServiceBusQueueEventSubscriptionDestinationResponseOutput struct{ *pulumi.OutputState }

func (ServiceBusQueueEventSubscriptionDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusQueueEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (o ServiceBusQueueEventSubscriptionDestinationResponseOutput) ToServiceBusQueueEventSubscriptionDestinationResponseOutput() ServiceBusQueueEventSubscriptionDestinationResponseOutput {
	return o
}

func (o ServiceBusQueueEventSubscriptionDestinationResponseOutput) ToServiceBusQueueEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) ServiceBusQueueEventSubscriptionDestinationResponseOutput {
	return o
}

func (o ServiceBusQueueEventSubscriptionDestinationResponseOutput) DeliveryAttributeMappings() pulumi.ArrayOutput {
	return o.ApplyT(func(v ServiceBusQueueEventSubscriptionDestinationResponse) []interface{} {
		return v.DeliveryAttributeMappings
	}).(pulumi.ArrayOutput)
}

func (o ServiceBusQueueEventSubscriptionDestinationResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBusQueueEventSubscriptionDestinationResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o ServiceBusQueueEventSubscriptionDestinationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueEventSubscriptionDestinationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ServiceBusTopicEventSubscriptionDestination struct {
	DeliveryAttributeMappings []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType              string        `pulumi:"endpointType"`
	ResourceId                *string       `pulumi:"resourceId"`
}





type ServiceBusTopicEventSubscriptionDestinationInput interface {
	pulumi.Input

	ToServiceBusTopicEventSubscriptionDestinationOutput() ServiceBusTopicEventSubscriptionDestinationOutput
	ToServiceBusTopicEventSubscriptionDestinationOutputWithContext(context.Context) ServiceBusTopicEventSubscriptionDestinationOutput
}

type ServiceBusTopicEventSubscriptionDestinationArgs struct {
	DeliveryAttributeMappings pulumi.ArrayInput     `pulumi:"deliveryAttributeMappings"`
	EndpointType              pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId                pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (ServiceBusTopicEventSubscriptionDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusTopicEventSubscriptionDestination)(nil)).Elem()
}

func (i ServiceBusTopicEventSubscriptionDestinationArgs) ToServiceBusTopicEventSubscriptionDestinationOutput() ServiceBusTopicEventSubscriptionDestinationOutput {
	return i.ToServiceBusTopicEventSubscriptionDestinationOutputWithContext(context.Background())
}

func (i ServiceBusTopicEventSubscriptionDestinationArgs) ToServiceBusTopicEventSubscriptionDestinationOutputWithContext(ctx context.Context) ServiceBusTopicEventSubscriptionDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusTopicEventSubscriptionDestinationOutput)
}

type ServiceBusTopicEventSubscriptionDestinationOutput struct{ *pulumi.OutputState }

func (ServiceBusTopicEventSubscriptionDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusTopicEventSubscriptionDestination)(nil)).Elem()
}

func (o ServiceBusTopicEventSubscriptionDestinationOutput) ToServiceBusTopicEventSubscriptionDestinationOutput() ServiceBusTopicEventSubscriptionDestinationOutput {
	return o
}

func (o ServiceBusTopicEventSubscriptionDestinationOutput) ToServiceBusTopicEventSubscriptionDestinationOutputWithContext(ctx context.Context) ServiceBusTopicEventSubscriptionDestinationOutput {
	return o
}

func (o ServiceBusTopicEventSubscriptionDestinationOutput) DeliveryAttributeMappings() pulumi.ArrayOutput {
	return o.ApplyT(func(v ServiceBusTopicEventSubscriptionDestination) []interface{} { return v.DeliveryAttributeMappings }).(pulumi.ArrayOutput)
}

func (o ServiceBusTopicEventSubscriptionDestinationOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBusTopicEventSubscriptionDestination) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o ServiceBusTopicEventSubscriptionDestinationOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicEventSubscriptionDestination) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ServiceBusTopicEventSubscriptionDestinationResponse struct {
	DeliveryAttributeMappings []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType              string        `pulumi:"endpointType"`
	ResourceId                *string       `pulumi:"resourceId"`
}





type ServiceBusTopicEventSubscriptionDestinationResponseInput interface {
	pulumi.Input

	ToServiceBusTopicEventSubscriptionDestinationResponseOutput() ServiceBusTopicEventSubscriptionDestinationResponseOutput
	ToServiceBusTopicEventSubscriptionDestinationResponseOutputWithContext(context.Context) ServiceBusTopicEventSubscriptionDestinationResponseOutput
}

type ServiceBusTopicEventSubscriptionDestinationResponseArgs struct {
	DeliveryAttributeMappings pulumi.ArrayInput     `pulumi:"deliveryAttributeMappings"`
	EndpointType              pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId                pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (ServiceBusTopicEventSubscriptionDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusTopicEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (i ServiceBusTopicEventSubscriptionDestinationResponseArgs) ToServiceBusTopicEventSubscriptionDestinationResponseOutput() ServiceBusTopicEventSubscriptionDestinationResponseOutput {
	return i.ToServiceBusTopicEventSubscriptionDestinationResponseOutputWithContext(context.Background())
}

func (i ServiceBusTopicEventSubscriptionDestinationResponseArgs) ToServiceBusTopicEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) ServiceBusTopicEventSubscriptionDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusTopicEventSubscriptionDestinationResponseOutput)
}

type ServiceBusTopicEventSubscriptionDestinationResponseOutput struct{ *pulumi.OutputState }

func (ServiceBusTopicEventSubscriptionDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusTopicEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (o ServiceBusTopicEventSubscriptionDestinationResponseOutput) ToServiceBusTopicEventSubscriptionDestinationResponseOutput() ServiceBusTopicEventSubscriptionDestinationResponseOutput {
	return o
}

func (o ServiceBusTopicEventSubscriptionDestinationResponseOutput) ToServiceBusTopicEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) ServiceBusTopicEventSubscriptionDestinationResponseOutput {
	return o
}

func (o ServiceBusTopicEventSubscriptionDestinationResponseOutput) DeliveryAttributeMappings() pulumi.ArrayOutput {
	return o.ApplyT(func(v ServiceBusTopicEventSubscriptionDestinationResponse) []interface{} {
		return v.DeliveryAttributeMappings
	}).(pulumi.ArrayOutput)
}

func (o ServiceBusTopicEventSubscriptionDestinationResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBusTopicEventSubscriptionDestinationResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o ServiceBusTopicEventSubscriptionDestinationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicEventSubscriptionDestinationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type StaticDeliveryAttributeMapping struct {
	IsSecret *bool   `pulumi:"isSecret"`
	Name     *string `pulumi:"name"`
	Type     string  `pulumi:"type"`
	Value    *string `pulumi:"value"`
}





type StaticDeliveryAttributeMappingInput interface {
	pulumi.Input

	ToStaticDeliveryAttributeMappingOutput() StaticDeliveryAttributeMappingOutput
	ToStaticDeliveryAttributeMappingOutputWithContext(context.Context) StaticDeliveryAttributeMappingOutput
}

type StaticDeliveryAttributeMappingArgs struct {
	IsSecret pulumi.BoolPtrInput   `pulumi:"isSecret"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Type     pulumi.StringInput    `pulumi:"type"`
	Value    pulumi.StringPtrInput `pulumi:"value"`
}

func (StaticDeliveryAttributeMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticDeliveryAttributeMapping)(nil)).Elem()
}

func (i StaticDeliveryAttributeMappingArgs) ToStaticDeliveryAttributeMappingOutput() StaticDeliveryAttributeMappingOutput {
	return i.ToStaticDeliveryAttributeMappingOutputWithContext(context.Background())
}

func (i StaticDeliveryAttributeMappingArgs) ToStaticDeliveryAttributeMappingOutputWithContext(ctx context.Context) StaticDeliveryAttributeMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticDeliveryAttributeMappingOutput)
}

type StaticDeliveryAttributeMappingOutput struct{ *pulumi.OutputState }

func (StaticDeliveryAttributeMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticDeliveryAttributeMapping)(nil)).Elem()
}

func (o StaticDeliveryAttributeMappingOutput) ToStaticDeliveryAttributeMappingOutput() StaticDeliveryAttributeMappingOutput {
	return o
}

func (o StaticDeliveryAttributeMappingOutput) ToStaticDeliveryAttributeMappingOutputWithContext(ctx context.Context) StaticDeliveryAttributeMappingOutput {
	return o
}

func (o StaticDeliveryAttributeMappingOutput) GetIsSecret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StaticDeliveryAttributeMapping) *bool { return v.IsSecret }).(pulumi.BoolPtrOutput)
}

func (o StaticDeliveryAttributeMappingOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticDeliveryAttributeMapping) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StaticDeliveryAttributeMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StaticDeliveryAttributeMapping) string { return v.Type }).(pulumi.StringOutput)
}

func (o StaticDeliveryAttributeMappingOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticDeliveryAttributeMapping) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type StaticDeliveryAttributeMappingResponse struct {
	IsSecret *bool   `pulumi:"isSecret"`
	Name     *string `pulumi:"name"`
	Type     string  `pulumi:"type"`
	Value    *string `pulumi:"value"`
}





type StaticDeliveryAttributeMappingResponseInput interface {
	pulumi.Input

	ToStaticDeliveryAttributeMappingResponseOutput() StaticDeliveryAttributeMappingResponseOutput
	ToStaticDeliveryAttributeMappingResponseOutputWithContext(context.Context) StaticDeliveryAttributeMappingResponseOutput
}

type StaticDeliveryAttributeMappingResponseArgs struct {
	IsSecret pulumi.BoolPtrInput   `pulumi:"isSecret"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Type     pulumi.StringInput    `pulumi:"type"`
	Value    pulumi.StringPtrInput `pulumi:"value"`
}

func (StaticDeliveryAttributeMappingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticDeliveryAttributeMappingResponse)(nil)).Elem()
}

func (i StaticDeliveryAttributeMappingResponseArgs) ToStaticDeliveryAttributeMappingResponseOutput() StaticDeliveryAttributeMappingResponseOutput {
	return i.ToStaticDeliveryAttributeMappingResponseOutputWithContext(context.Background())
}

func (i StaticDeliveryAttributeMappingResponseArgs) ToStaticDeliveryAttributeMappingResponseOutputWithContext(ctx context.Context) StaticDeliveryAttributeMappingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticDeliveryAttributeMappingResponseOutput)
}

type StaticDeliveryAttributeMappingResponseOutput struct{ *pulumi.OutputState }

func (StaticDeliveryAttributeMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticDeliveryAttributeMappingResponse)(nil)).Elem()
}

func (o StaticDeliveryAttributeMappingResponseOutput) ToStaticDeliveryAttributeMappingResponseOutput() StaticDeliveryAttributeMappingResponseOutput {
	return o
}

func (o StaticDeliveryAttributeMappingResponseOutput) ToStaticDeliveryAttributeMappingResponseOutputWithContext(ctx context.Context) StaticDeliveryAttributeMappingResponseOutput {
	return o
}

func (o StaticDeliveryAttributeMappingResponseOutput) GetIsSecret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StaticDeliveryAttributeMappingResponse) *bool { return v.IsSecret }).(pulumi.BoolPtrOutput)
}

func (o StaticDeliveryAttributeMappingResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticDeliveryAttributeMappingResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StaticDeliveryAttributeMappingResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StaticDeliveryAttributeMappingResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o StaticDeliveryAttributeMappingResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticDeliveryAttributeMappingResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type StorageBlobDeadLetterDestination struct {
	BlobContainerName *string `pulumi:"blobContainerName"`
	EndpointType      string  `pulumi:"endpointType"`
	ResourceId        *string `pulumi:"resourceId"`
}





type StorageBlobDeadLetterDestinationInput interface {
	pulumi.Input

	ToStorageBlobDeadLetterDestinationOutput() StorageBlobDeadLetterDestinationOutput
	ToStorageBlobDeadLetterDestinationOutputWithContext(context.Context) StorageBlobDeadLetterDestinationOutput
}

type StorageBlobDeadLetterDestinationArgs struct {
	BlobContainerName pulumi.StringPtrInput `pulumi:"blobContainerName"`
	EndpointType      pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId        pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (StorageBlobDeadLetterDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageBlobDeadLetterDestination)(nil)).Elem()
}

func (i StorageBlobDeadLetterDestinationArgs) ToStorageBlobDeadLetterDestinationOutput() StorageBlobDeadLetterDestinationOutput {
	return i.ToStorageBlobDeadLetterDestinationOutputWithContext(context.Background())
}

func (i StorageBlobDeadLetterDestinationArgs) ToStorageBlobDeadLetterDestinationOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBlobDeadLetterDestinationOutput)
}

func (i StorageBlobDeadLetterDestinationArgs) ToStorageBlobDeadLetterDestinationPtrOutput() StorageBlobDeadLetterDestinationPtrOutput {
	return i.ToStorageBlobDeadLetterDestinationPtrOutputWithContext(context.Background())
}

func (i StorageBlobDeadLetterDestinationArgs) ToStorageBlobDeadLetterDestinationPtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBlobDeadLetterDestinationOutput).ToStorageBlobDeadLetterDestinationPtrOutputWithContext(ctx)
}









type StorageBlobDeadLetterDestinationPtrInput interface {
	pulumi.Input

	ToStorageBlobDeadLetterDestinationPtrOutput() StorageBlobDeadLetterDestinationPtrOutput
	ToStorageBlobDeadLetterDestinationPtrOutputWithContext(context.Context) StorageBlobDeadLetterDestinationPtrOutput
}

type storageBlobDeadLetterDestinationPtrType StorageBlobDeadLetterDestinationArgs

func StorageBlobDeadLetterDestinationPtr(v *StorageBlobDeadLetterDestinationArgs) StorageBlobDeadLetterDestinationPtrInput {
	return (*storageBlobDeadLetterDestinationPtrType)(v)
}

func (*storageBlobDeadLetterDestinationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBlobDeadLetterDestination)(nil)).Elem()
}

func (i *storageBlobDeadLetterDestinationPtrType) ToStorageBlobDeadLetterDestinationPtrOutput() StorageBlobDeadLetterDestinationPtrOutput {
	return i.ToStorageBlobDeadLetterDestinationPtrOutputWithContext(context.Background())
}

func (i *storageBlobDeadLetterDestinationPtrType) ToStorageBlobDeadLetterDestinationPtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBlobDeadLetterDestinationPtrOutput)
}

type StorageBlobDeadLetterDestinationOutput struct{ *pulumi.OutputState }

func (StorageBlobDeadLetterDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageBlobDeadLetterDestination)(nil)).Elem()
}

func (o StorageBlobDeadLetterDestinationOutput) ToStorageBlobDeadLetterDestinationOutput() StorageBlobDeadLetterDestinationOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationOutput) ToStorageBlobDeadLetterDestinationOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationOutput) ToStorageBlobDeadLetterDestinationPtrOutput() StorageBlobDeadLetterDestinationPtrOutput {
	return o.ToStorageBlobDeadLetterDestinationPtrOutputWithContext(context.Background())
}

func (o StorageBlobDeadLetterDestinationOutput) ToStorageBlobDeadLetterDestinationPtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageBlobDeadLetterDestination) *StorageBlobDeadLetterDestination {
		return &v
	}).(StorageBlobDeadLetterDestinationPtrOutput)
}

func (o StorageBlobDeadLetterDestinationOutput) BlobContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageBlobDeadLetterDestination) *string { return v.BlobContainerName }).(pulumi.StringPtrOutput)
}

func (o StorageBlobDeadLetterDestinationOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v StorageBlobDeadLetterDestination) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o StorageBlobDeadLetterDestinationOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageBlobDeadLetterDestination) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type StorageBlobDeadLetterDestinationPtrOutput struct{ *pulumi.OutputState }

func (StorageBlobDeadLetterDestinationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBlobDeadLetterDestination)(nil)).Elem()
}

func (o StorageBlobDeadLetterDestinationPtrOutput) ToStorageBlobDeadLetterDestinationPtrOutput() StorageBlobDeadLetterDestinationPtrOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationPtrOutput) ToStorageBlobDeadLetterDestinationPtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationPtrOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationPtrOutput) Elem() StorageBlobDeadLetterDestinationOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestination) StorageBlobDeadLetterDestination {
		if v != nil {
			return *v
		}
		var ret StorageBlobDeadLetterDestination
		return ret
	}).(StorageBlobDeadLetterDestinationOutput)
}

func (o StorageBlobDeadLetterDestinationPtrOutput) BlobContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestination) *string {
		if v == nil {
			return nil
		}
		return v.BlobContainerName
	}).(pulumi.StringPtrOutput)
}

func (o StorageBlobDeadLetterDestinationPtrOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestination) *string {
		if v == nil {
			return nil
		}
		return &v.EndpointType
	}).(pulumi.StringPtrOutput)
}

func (o StorageBlobDeadLetterDestinationPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestination) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type StorageBlobDeadLetterDestinationResponse struct {
	BlobContainerName *string `pulumi:"blobContainerName"`
	EndpointType      string  `pulumi:"endpointType"`
	ResourceId        *string `pulumi:"resourceId"`
}





type StorageBlobDeadLetterDestinationResponseInput interface {
	pulumi.Input

	ToStorageBlobDeadLetterDestinationResponseOutput() StorageBlobDeadLetterDestinationResponseOutput
	ToStorageBlobDeadLetterDestinationResponseOutputWithContext(context.Context) StorageBlobDeadLetterDestinationResponseOutput
}

type StorageBlobDeadLetterDestinationResponseArgs struct {
	BlobContainerName pulumi.StringPtrInput `pulumi:"blobContainerName"`
	EndpointType      pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId        pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (StorageBlobDeadLetterDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageBlobDeadLetterDestinationResponse)(nil)).Elem()
}

func (i StorageBlobDeadLetterDestinationResponseArgs) ToStorageBlobDeadLetterDestinationResponseOutput() StorageBlobDeadLetterDestinationResponseOutput {
	return i.ToStorageBlobDeadLetterDestinationResponseOutputWithContext(context.Background())
}

func (i StorageBlobDeadLetterDestinationResponseArgs) ToStorageBlobDeadLetterDestinationResponseOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBlobDeadLetterDestinationResponseOutput)
}

func (i StorageBlobDeadLetterDestinationResponseArgs) ToStorageBlobDeadLetterDestinationResponsePtrOutput() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return i.ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(context.Background())
}

func (i StorageBlobDeadLetterDestinationResponseArgs) ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBlobDeadLetterDestinationResponseOutput).ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(ctx)
}









type StorageBlobDeadLetterDestinationResponsePtrInput interface {
	pulumi.Input

	ToStorageBlobDeadLetterDestinationResponsePtrOutput() StorageBlobDeadLetterDestinationResponsePtrOutput
	ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(context.Context) StorageBlobDeadLetterDestinationResponsePtrOutput
}

type storageBlobDeadLetterDestinationResponsePtrType StorageBlobDeadLetterDestinationResponseArgs

func StorageBlobDeadLetterDestinationResponsePtr(v *StorageBlobDeadLetterDestinationResponseArgs) StorageBlobDeadLetterDestinationResponsePtrInput {
	return (*storageBlobDeadLetterDestinationResponsePtrType)(v)
}

func (*storageBlobDeadLetterDestinationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBlobDeadLetterDestinationResponse)(nil)).Elem()
}

func (i *storageBlobDeadLetterDestinationResponsePtrType) ToStorageBlobDeadLetterDestinationResponsePtrOutput() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return i.ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(context.Background())
}

func (i *storageBlobDeadLetterDestinationResponsePtrType) ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

type StorageBlobDeadLetterDestinationResponseOutput struct{ *pulumi.OutputState }

func (StorageBlobDeadLetterDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageBlobDeadLetterDestinationResponse)(nil)).Elem()
}

func (o StorageBlobDeadLetterDestinationResponseOutput) ToStorageBlobDeadLetterDestinationResponseOutput() StorageBlobDeadLetterDestinationResponseOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationResponseOutput) ToStorageBlobDeadLetterDestinationResponseOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationResponseOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationResponseOutput) ToStorageBlobDeadLetterDestinationResponsePtrOutput() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(context.Background())
}

func (o StorageBlobDeadLetterDestinationResponseOutput) ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageBlobDeadLetterDestinationResponse) *StorageBlobDeadLetterDestinationResponse {
		return &v
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

func (o StorageBlobDeadLetterDestinationResponseOutput) BlobContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageBlobDeadLetterDestinationResponse) *string { return v.BlobContainerName }).(pulumi.StringPtrOutput)
}

func (o StorageBlobDeadLetterDestinationResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v StorageBlobDeadLetterDestinationResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o StorageBlobDeadLetterDestinationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageBlobDeadLetterDestinationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type StorageBlobDeadLetterDestinationResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageBlobDeadLetterDestinationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBlobDeadLetterDestinationResponse)(nil)).Elem()
}

func (o StorageBlobDeadLetterDestinationResponsePtrOutput) ToStorageBlobDeadLetterDestinationResponsePtrOutput() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationResponsePtrOutput) ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationResponsePtrOutput) Elem() StorageBlobDeadLetterDestinationResponseOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestinationResponse) StorageBlobDeadLetterDestinationResponse {
		if v != nil {
			return *v
		}
		var ret StorageBlobDeadLetterDestinationResponse
		return ret
	}).(StorageBlobDeadLetterDestinationResponseOutput)
}

func (o StorageBlobDeadLetterDestinationResponsePtrOutput) BlobContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.BlobContainerName
	}).(pulumi.StringPtrOutput)
}

func (o StorageBlobDeadLetterDestinationResponsePtrOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EndpointType
	}).(pulumi.StringPtrOutput)
}

func (o StorageBlobDeadLetterDestinationResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type StorageQueueEventSubscriptionDestination struct {
	EndpointType                    string   `pulumi:"endpointType"`
	QueueMessageTimeToLiveInSeconds *float64 `pulumi:"queueMessageTimeToLiveInSeconds"`
	QueueName                       *string  `pulumi:"queueName"`
	ResourceId                      *string  `pulumi:"resourceId"`
}





type StorageQueueEventSubscriptionDestinationInput interface {
	pulumi.Input

	ToStorageQueueEventSubscriptionDestinationOutput() StorageQueueEventSubscriptionDestinationOutput
	ToStorageQueueEventSubscriptionDestinationOutputWithContext(context.Context) StorageQueueEventSubscriptionDestinationOutput
}

type StorageQueueEventSubscriptionDestinationArgs struct {
	EndpointType                    pulumi.StringInput     `pulumi:"endpointType"`
	QueueMessageTimeToLiveInSeconds pulumi.Float64PtrInput `pulumi:"queueMessageTimeToLiveInSeconds"`
	QueueName                       pulumi.StringPtrInput  `pulumi:"queueName"`
	ResourceId                      pulumi.StringPtrInput  `pulumi:"resourceId"`
}

func (StorageQueueEventSubscriptionDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQueueEventSubscriptionDestination)(nil)).Elem()
}

func (i StorageQueueEventSubscriptionDestinationArgs) ToStorageQueueEventSubscriptionDestinationOutput() StorageQueueEventSubscriptionDestinationOutput {
	return i.ToStorageQueueEventSubscriptionDestinationOutputWithContext(context.Background())
}

func (i StorageQueueEventSubscriptionDestinationArgs) ToStorageQueueEventSubscriptionDestinationOutputWithContext(ctx context.Context) StorageQueueEventSubscriptionDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageQueueEventSubscriptionDestinationOutput)
}

type StorageQueueEventSubscriptionDestinationOutput struct{ *pulumi.OutputState }

func (StorageQueueEventSubscriptionDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQueueEventSubscriptionDestination)(nil)).Elem()
}

func (o StorageQueueEventSubscriptionDestinationOutput) ToStorageQueueEventSubscriptionDestinationOutput() StorageQueueEventSubscriptionDestinationOutput {
	return o
}

func (o StorageQueueEventSubscriptionDestinationOutput) ToStorageQueueEventSubscriptionDestinationOutputWithContext(ctx context.Context) StorageQueueEventSubscriptionDestinationOutput {
	return o
}

func (o StorageQueueEventSubscriptionDestinationOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v StorageQueueEventSubscriptionDestination) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o StorageQueueEventSubscriptionDestinationOutput) QueueMessageTimeToLiveInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v StorageQueueEventSubscriptionDestination) *float64 { return v.QueueMessageTimeToLiveInSeconds }).(pulumi.Float64PtrOutput)
}

func (o StorageQueueEventSubscriptionDestinationOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueEventSubscriptionDestination) *string { return v.QueueName }).(pulumi.StringPtrOutput)
}

func (o StorageQueueEventSubscriptionDestinationOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueEventSubscriptionDestination) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type StorageQueueEventSubscriptionDestinationResponse struct {
	EndpointType                    string   `pulumi:"endpointType"`
	QueueMessageTimeToLiveInSeconds *float64 `pulumi:"queueMessageTimeToLiveInSeconds"`
	QueueName                       *string  `pulumi:"queueName"`
	ResourceId                      *string  `pulumi:"resourceId"`
}





type StorageQueueEventSubscriptionDestinationResponseInput interface {
	pulumi.Input

	ToStorageQueueEventSubscriptionDestinationResponseOutput() StorageQueueEventSubscriptionDestinationResponseOutput
	ToStorageQueueEventSubscriptionDestinationResponseOutputWithContext(context.Context) StorageQueueEventSubscriptionDestinationResponseOutput
}

type StorageQueueEventSubscriptionDestinationResponseArgs struct {
	EndpointType                    pulumi.StringInput     `pulumi:"endpointType"`
	QueueMessageTimeToLiveInSeconds pulumi.Float64PtrInput `pulumi:"queueMessageTimeToLiveInSeconds"`
	QueueName                       pulumi.StringPtrInput  `pulumi:"queueName"`
	ResourceId                      pulumi.StringPtrInput  `pulumi:"resourceId"`
}

func (StorageQueueEventSubscriptionDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQueueEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (i StorageQueueEventSubscriptionDestinationResponseArgs) ToStorageQueueEventSubscriptionDestinationResponseOutput() StorageQueueEventSubscriptionDestinationResponseOutput {
	return i.ToStorageQueueEventSubscriptionDestinationResponseOutputWithContext(context.Background())
}

func (i StorageQueueEventSubscriptionDestinationResponseArgs) ToStorageQueueEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) StorageQueueEventSubscriptionDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageQueueEventSubscriptionDestinationResponseOutput)
}

type StorageQueueEventSubscriptionDestinationResponseOutput struct{ *pulumi.OutputState }

func (StorageQueueEventSubscriptionDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQueueEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (o StorageQueueEventSubscriptionDestinationResponseOutput) ToStorageQueueEventSubscriptionDestinationResponseOutput() StorageQueueEventSubscriptionDestinationResponseOutput {
	return o
}

func (o StorageQueueEventSubscriptionDestinationResponseOutput) ToStorageQueueEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) StorageQueueEventSubscriptionDestinationResponseOutput {
	return o
}

func (o StorageQueueEventSubscriptionDestinationResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v StorageQueueEventSubscriptionDestinationResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o StorageQueueEventSubscriptionDestinationResponseOutput) QueueMessageTimeToLiveInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v StorageQueueEventSubscriptionDestinationResponse) *float64 {
		return v.QueueMessageTimeToLiveInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o StorageQueueEventSubscriptionDestinationResponseOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueEventSubscriptionDestinationResponse) *string { return v.QueueName }).(pulumi.StringPtrOutput)
}

func (o StorageQueueEventSubscriptionDestinationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueEventSubscriptionDestinationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type StringBeginsWithAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}





type StringBeginsWithAdvancedFilterInput interface {
	pulumi.Input

	ToStringBeginsWithAdvancedFilterOutput() StringBeginsWithAdvancedFilterOutput
	ToStringBeginsWithAdvancedFilterOutputWithContext(context.Context) StringBeginsWithAdvancedFilterOutput
}

type StringBeginsWithAdvancedFilterArgs struct {
	Key          pulumi.StringPtrInput   `pulumi:"key"`
	OperatorType pulumi.StringInput      `pulumi:"operatorType"`
	Values       pulumi.StringArrayInput `pulumi:"values"`
}

func (StringBeginsWithAdvancedFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StringBeginsWithAdvancedFilter)(nil)).Elem()
}

func (i StringBeginsWithAdvancedFilterArgs) ToStringBeginsWithAdvancedFilterOutput() StringBeginsWithAdvancedFilterOutput {
	return i.ToStringBeginsWithAdvancedFilterOutputWithContext(context.Background())
}

func (i StringBeginsWithAdvancedFilterArgs) ToStringBeginsWithAdvancedFilterOutputWithContext(ctx context.Context) StringBeginsWithAdvancedFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StringBeginsWithAdvancedFilterOutput)
}

type StringBeginsWithAdvancedFilterOutput struct{ *pulumi.OutputState }

func (StringBeginsWithAdvancedFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StringBeginsWithAdvancedFilter)(nil)).Elem()
}

func (o StringBeginsWithAdvancedFilterOutput) ToStringBeginsWithAdvancedFilterOutput() StringBeginsWithAdvancedFilterOutput {
	return o
}

func (o StringBeginsWithAdvancedFilterOutput) ToStringBeginsWithAdvancedFilterOutputWithContext(ctx context.Context) StringBeginsWithAdvancedFilterOutput {
	return o
}

func (o StringBeginsWithAdvancedFilterOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StringBeginsWithAdvancedFilter) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o StringBeginsWithAdvancedFilterOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v StringBeginsWithAdvancedFilter) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o StringBeginsWithAdvancedFilterOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StringBeginsWithAdvancedFilter) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type StringBeginsWithAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}





type StringBeginsWithAdvancedFilterResponseInput interface {
	pulumi.Input

	ToStringBeginsWithAdvancedFilterResponseOutput() StringBeginsWithAdvancedFilterResponseOutput
	ToStringBeginsWithAdvancedFilterResponseOutputWithContext(context.Context) StringBeginsWithAdvancedFilterResponseOutput
}

type StringBeginsWithAdvancedFilterResponseArgs struct {
	Key          pulumi.StringPtrInput   `pulumi:"key"`
	OperatorType pulumi.StringInput      `pulumi:"operatorType"`
	Values       pulumi.StringArrayInput `pulumi:"values"`
}

func (StringBeginsWithAdvancedFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StringBeginsWithAdvancedFilterResponse)(nil)).Elem()
}

func (i StringBeginsWithAdvancedFilterResponseArgs) ToStringBeginsWithAdvancedFilterResponseOutput() StringBeginsWithAdvancedFilterResponseOutput {
	return i.ToStringBeginsWithAdvancedFilterResponseOutputWithContext(context.Background())
}

func (i StringBeginsWithAdvancedFilterResponseArgs) ToStringBeginsWithAdvancedFilterResponseOutputWithContext(ctx context.Context) StringBeginsWithAdvancedFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StringBeginsWithAdvancedFilterResponseOutput)
}

type StringBeginsWithAdvancedFilterResponseOutput struct{ *pulumi.OutputState }

func (StringBeginsWithAdvancedFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StringBeginsWithAdvancedFilterResponse)(nil)).Elem()
}

func (o StringBeginsWithAdvancedFilterResponseOutput) ToStringBeginsWithAdvancedFilterResponseOutput() StringBeginsWithAdvancedFilterResponseOutput {
	return o
}

func (o StringBeginsWithAdvancedFilterResponseOutput) ToStringBeginsWithAdvancedFilterResponseOutputWithContext(ctx context.Context) StringBeginsWithAdvancedFilterResponseOutput {
	return o
}

func (o StringBeginsWithAdvancedFilterResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StringBeginsWithAdvancedFilterResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o StringBeginsWithAdvancedFilterResponseOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v StringBeginsWithAdvancedFilterResponse) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o StringBeginsWithAdvancedFilterResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StringBeginsWithAdvancedFilterResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type StringContainsAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}





type StringContainsAdvancedFilterInput interface {
	pulumi.Input

	ToStringContainsAdvancedFilterOutput() StringContainsAdvancedFilterOutput
	ToStringContainsAdvancedFilterOutputWithContext(context.Context) StringContainsAdvancedFilterOutput
}

type StringContainsAdvancedFilterArgs struct {
	Key          pulumi.StringPtrInput   `pulumi:"key"`
	OperatorType pulumi.StringInput      `pulumi:"operatorType"`
	Values       pulumi.StringArrayInput `pulumi:"values"`
}

func (StringContainsAdvancedFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StringContainsAdvancedFilter)(nil)).Elem()
}

func (i StringContainsAdvancedFilterArgs) ToStringContainsAdvancedFilterOutput() StringContainsAdvancedFilterOutput {
	return i.ToStringContainsAdvancedFilterOutputWithContext(context.Background())
}

func (i StringContainsAdvancedFilterArgs) ToStringContainsAdvancedFilterOutputWithContext(ctx context.Context) StringContainsAdvancedFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StringContainsAdvancedFilterOutput)
}

type StringContainsAdvancedFilterOutput struct{ *pulumi.OutputState }

func (StringContainsAdvancedFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StringContainsAdvancedFilter)(nil)).Elem()
}

func (o StringContainsAdvancedFilterOutput) ToStringContainsAdvancedFilterOutput() StringContainsAdvancedFilterOutput {
	return o
}

func (o StringContainsAdvancedFilterOutput) ToStringContainsAdvancedFilterOutputWithContext(ctx context.Context) StringContainsAdvancedFilterOutput {
	return o
}

func (o StringContainsAdvancedFilterOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StringContainsAdvancedFilter) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o StringContainsAdvancedFilterOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v StringContainsAdvancedFilter) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o StringContainsAdvancedFilterOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StringContainsAdvancedFilter) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type StringContainsAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}





type StringContainsAdvancedFilterResponseInput interface {
	pulumi.Input

	ToStringContainsAdvancedFilterResponseOutput() StringContainsAdvancedFilterResponseOutput
	ToStringContainsAdvancedFilterResponseOutputWithContext(context.Context) StringContainsAdvancedFilterResponseOutput
}

type StringContainsAdvancedFilterResponseArgs struct {
	Key          pulumi.StringPtrInput   `pulumi:"key"`
	OperatorType pulumi.StringInput      `pulumi:"operatorType"`
	Values       pulumi.StringArrayInput `pulumi:"values"`
}

func (StringContainsAdvancedFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StringContainsAdvancedFilterResponse)(nil)).Elem()
}

func (i StringContainsAdvancedFilterResponseArgs) ToStringContainsAdvancedFilterResponseOutput() StringContainsAdvancedFilterResponseOutput {
	return i.ToStringContainsAdvancedFilterResponseOutputWithContext(context.Background())
}

func (i StringContainsAdvancedFilterResponseArgs) ToStringContainsAdvancedFilterResponseOutputWithContext(ctx context.Context) StringContainsAdvancedFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StringContainsAdvancedFilterResponseOutput)
}

type StringContainsAdvancedFilterResponseOutput struct{ *pulumi.OutputState }

func (StringContainsAdvancedFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StringContainsAdvancedFilterResponse)(nil)).Elem()
}

func (o StringContainsAdvancedFilterResponseOutput) ToStringContainsAdvancedFilterResponseOutput() StringContainsAdvancedFilterResponseOutput {
	return o
}

func (o StringContainsAdvancedFilterResponseOutput) ToStringContainsAdvancedFilterResponseOutputWithContext(ctx context.Context) StringContainsAdvancedFilterResponseOutput {
	return o
}

func (o StringContainsAdvancedFilterResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StringContainsAdvancedFilterResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o StringContainsAdvancedFilterResponseOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v StringContainsAdvancedFilterResponse) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o StringContainsAdvancedFilterResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StringContainsAdvancedFilterResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type StringEndsWithAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}





type StringEndsWithAdvancedFilterInput interface {
	pulumi.Input

	ToStringEndsWithAdvancedFilterOutput() StringEndsWithAdvancedFilterOutput
	ToStringEndsWithAdvancedFilterOutputWithContext(context.Context) StringEndsWithAdvancedFilterOutput
}

type StringEndsWithAdvancedFilterArgs struct {
	Key          pulumi.StringPtrInput   `pulumi:"key"`
	OperatorType pulumi.StringInput      `pulumi:"operatorType"`
	Values       pulumi.StringArrayInput `pulumi:"values"`
}

func (StringEndsWithAdvancedFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StringEndsWithAdvancedFilter)(nil)).Elem()
}

func (i StringEndsWithAdvancedFilterArgs) ToStringEndsWithAdvancedFilterOutput() StringEndsWithAdvancedFilterOutput {
	return i.ToStringEndsWithAdvancedFilterOutputWithContext(context.Background())
}

func (i StringEndsWithAdvancedFilterArgs) ToStringEndsWithAdvancedFilterOutputWithContext(ctx context.Context) StringEndsWithAdvancedFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StringEndsWithAdvancedFilterOutput)
}

type StringEndsWithAdvancedFilterOutput struct{ *pulumi.OutputState }

func (StringEndsWithAdvancedFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StringEndsWithAdvancedFilter)(nil)).Elem()
}

func (o StringEndsWithAdvancedFilterOutput) ToStringEndsWithAdvancedFilterOutput() StringEndsWithAdvancedFilterOutput {
	return o
}

func (o StringEndsWithAdvancedFilterOutput) ToStringEndsWithAdvancedFilterOutputWithContext(ctx context.Context) StringEndsWithAdvancedFilterOutput {
	return o
}

func (o StringEndsWithAdvancedFilterOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StringEndsWithAdvancedFilter) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o StringEndsWithAdvancedFilterOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v StringEndsWithAdvancedFilter) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o StringEndsWithAdvancedFilterOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StringEndsWithAdvancedFilter) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type StringEndsWithAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}





type StringEndsWithAdvancedFilterResponseInput interface {
	pulumi.Input

	ToStringEndsWithAdvancedFilterResponseOutput() StringEndsWithAdvancedFilterResponseOutput
	ToStringEndsWithAdvancedFilterResponseOutputWithContext(context.Context) StringEndsWithAdvancedFilterResponseOutput
}

type StringEndsWithAdvancedFilterResponseArgs struct {
	Key          pulumi.StringPtrInput   `pulumi:"key"`
	OperatorType pulumi.StringInput      `pulumi:"operatorType"`
	Values       pulumi.StringArrayInput `pulumi:"values"`
}

func (StringEndsWithAdvancedFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StringEndsWithAdvancedFilterResponse)(nil)).Elem()
}

func (i StringEndsWithAdvancedFilterResponseArgs) ToStringEndsWithAdvancedFilterResponseOutput() StringEndsWithAdvancedFilterResponseOutput {
	return i.ToStringEndsWithAdvancedFilterResponseOutputWithContext(context.Background())
}

func (i StringEndsWithAdvancedFilterResponseArgs) ToStringEndsWithAdvancedFilterResponseOutputWithContext(ctx context.Context) StringEndsWithAdvancedFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StringEndsWithAdvancedFilterResponseOutput)
}

type StringEndsWithAdvancedFilterResponseOutput struct{ *pulumi.OutputState }

func (StringEndsWithAdvancedFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StringEndsWithAdvancedFilterResponse)(nil)).Elem()
}

func (o StringEndsWithAdvancedFilterResponseOutput) ToStringEndsWithAdvancedFilterResponseOutput() StringEndsWithAdvancedFilterResponseOutput {
	return o
}

func (o StringEndsWithAdvancedFilterResponseOutput) ToStringEndsWithAdvancedFilterResponseOutputWithContext(ctx context.Context) StringEndsWithAdvancedFilterResponseOutput {
	return o
}

func (o StringEndsWithAdvancedFilterResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StringEndsWithAdvancedFilterResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o StringEndsWithAdvancedFilterResponseOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v StringEndsWithAdvancedFilterResponse) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o StringEndsWithAdvancedFilterResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StringEndsWithAdvancedFilterResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type StringInAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}





type StringInAdvancedFilterInput interface {
	pulumi.Input

	ToStringInAdvancedFilterOutput() StringInAdvancedFilterOutput
	ToStringInAdvancedFilterOutputWithContext(context.Context) StringInAdvancedFilterOutput
}

type StringInAdvancedFilterArgs struct {
	Key          pulumi.StringPtrInput   `pulumi:"key"`
	OperatorType pulumi.StringInput      `pulumi:"operatorType"`
	Values       pulumi.StringArrayInput `pulumi:"values"`
}

func (StringInAdvancedFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StringInAdvancedFilter)(nil)).Elem()
}

func (i StringInAdvancedFilterArgs) ToStringInAdvancedFilterOutput() StringInAdvancedFilterOutput {
	return i.ToStringInAdvancedFilterOutputWithContext(context.Background())
}

func (i StringInAdvancedFilterArgs) ToStringInAdvancedFilterOutputWithContext(ctx context.Context) StringInAdvancedFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StringInAdvancedFilterOutput)
}

type StringInAdvancedFilterOutput struct{ *pulumi.OutputState }

func (StringInAdvancedFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StringInAdvancedFilter)(nil)).Elem()
}

func (o StringInAdvancedFilterOutput) ToStringInAdvancedFilterOutput() StringInAdvancedFilterOutput {
	return o
}

func (o StringInAdvancedFilterOutput) ToStringInAdvancedFilterOutputWithContext(ctx context.Context) StringInAdvancedFilterOutput {
	return o
}

func (o StringInAdvancedFilterOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StringInAdvancedFilter) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o StringInAdvancedFilterOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v StringInAdvancedFilter) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o StringInAdvancedFilterOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StringInAdvancedFilter) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type StringInAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}





type StringInAdvancedFilterResponseInput interface {
	pulumi.Input

	ToStringInAdvancedFilterResponseOutput() StringInAdvancedFilterResponseOutput
	ToStringInAdvancedFilterResponseOutputWithContext(context.Context) StringInAdvancedFilterResponseOutput
}

type StringInAdvancedFilterResponseArgs struct {
	Key          pulumi.StringPtrInput   `pulumi:"key"`
	OperatorType pulumi.StringInput      `pulumi:"operatorType"`
	Values       pulumi.StringArrayInput `pulumi:"values"`
}

func (StringInAdvancedFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StringInAdvancedFilterResponse)(nil)).Elem()
}

func (i StringInAdvancedFilterResponseArgs) ToStringInAdvancedFilterResponseOutput() StringInAdvancedFilterResponseOutput {
	return i.ToStringInAdvancedFilterResponseOutputWithContext(context.Background())
}

func (i StringInAdvancedFilterResponseArgs) ToStringInAdvancedFilterResponseOutputWithContext(ctx context.Context) StringInAdvancedFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StringInAdvancedFilterResponseOutput)
}

type StringInAdvancedFilterResponseOutput struct{ *pulumi.OutputState }

func (StringInAdvancedFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StringInAdvancedFilterResponse)(nil)).Elem()
}

func (o StringInAdvancedFilterResponseOutput) ToStringInAdvancedFilterResponseOutput() StringInAdvancedFilterResponseOutput {
	return o
}

func (o StringInAdvancedFilterResponseOutput) ToStringInAdvancedFilterResponseOutputWithContext(ctx context.Context) StringInAdvancedFilterResponseOutput {
	return o
}

func (o StringInAdvancedFilterResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StringInAdvancedFilterResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o StringInAdvancedFilterResponseOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v StringInAdvancedFilterResponse) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o StringInAdvancedFilterResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StringInAdvancedFilterResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type StringNotInAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}





type StringNotInAdvancedFilterInput interface {
	pulumi.Input

	ToStringNotInAdvancedFilterOutput() StringNotInAdvancedFilterOutput
	ToStringNotInAdvancedFilterOutputWithContext(context.Context) StringNotInAdvancedFilterOutput
}

type StringNotInAdvancedFilterArgs struct {
	Key          pulumi.StringPtrInput   `pulumi:"key"`
	OperatorType pulumi.StringInput      `pulumi:"operatorType"`
	Values       pulumi.StringArrayInput `pulumi:"values"`
}

func (StringNotInAdvancedFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StringNotInAdvancedFilter)(nil)).Elem()
}

func (i StringNotInAdvancedFilterArgs) ToStringNotInAdvancedFilterOutput() StringNotInAdvancedFilterOutput {
	return i.ToStringNotInAdvancedFilterOutputWithContext(context.Background())
}

func (i StringNotInAdvancedFilterArgs) ToStringNotInAdvancedFilterOutputWithContext(ctx context.Context) StringNotInAdvancedFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StringNotInAdvancedFilterOutput)
}

type StringNotInAdvancedFilterOutput struct{ *pulumi.OutputState }

func (StringNotInAdvancedFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StringNotInAdvancedFilter)(nil)).Elem()
}

func (o StringNotInAdvancedFilterOutput) ToStringNotInAdvancedFilterOutput() StringNotInAdvancedFilterOutput {
	return o
}

func (o StringNotInAdvancedFilterOutput) ToStringNotInAdvancedFilterOutputWithContext(ctx context.Context) StringNotInAdvancedFilterOutput {
	return o
}

func (o StringNotInAdvancedFilterOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StringNotInAdvancedFilter) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o StringNotInAdvancedFilterOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v StringNotInAdvancedFilter) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o StringNotInAdvancedFilterOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StringNotInAdvancedFilter) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type StringNotInAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}





type StringNotInAdvancedFilterResponseInput interface {
	pulumi.Input

	ToStringNotInAdvancedFilterResponseOutput() StringNotInAdvancedFilterResponseOutput
	ToStringNotInAdvancedFilterResponseOutputWithContext(context.Context) StringNotInAdvancedFilterResponseOutput
}

type StringNotInAdvancedFilterResponseArgs struct {
	Key          pulumi.StringPtrInput   `pulumi:"key"`
	OperatorType pulumi.StringInput      `pulumi:"operatorType"`
	Values       pulumi.StringArrayInput `pulumi:"values"`
}

func (StringNotInAdvancedFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StringNotInAdvancedFilterResponse)(nil)).Elem()
}

func (i StringNotInAdvancedFilterResponseArgs) ToStringNotInAdvancedFilterResponseOutput() StringNotInAdvancedFilterResponseOutput {
	return i.ToStringNotInAdvancedFilterResponseOutputWithContext(context.Background())
}

func (i StringNotInAdvancedFilterResponseArgs) ToStringNotInAdvancedFilterResponseOutputWithContext(ctx context.Context) StringNotInAdvancedFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StringNotInAdvancedFilterResponseOutput)
}

type StringNotInAdvancedFilterResponseOutput struct{ *pulumi.OutputState }

func (StringNotInAdvancedFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StringNotInAdvancedFilterResponse)(nil)).Elem()
}

func (o StringNotInAdvancedFilterResponseOutput) ToStringNotInAdvancedFilterResponseOutput() StringNotInAdvancedFilterResponseOutput {
	return o
}

func (o StringNotInAdvancedFilterResponseOutput) ToStringNotInAdvancedFilterResponseOutputWithContext(ctx context.Context) StringNotInAdvancedFilterResponseOutput {
	return o
}

func (o StringNotInAdvancedFilterResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StringNotInAdvancedFilterResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o StringNotInAdvancedFilterResponseOutput) OperatorType() pulumi.StringOutput {
	return o.ApplyT(func(v StringNotInAdvancedFilterResponse) string { return v.OperatorType }).(pulumi.StringOutput)
}

func (o StringNotInAdvancedFilterResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StringNotInAdvancedFilterResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
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

type UserIdentityProperties struct {
	ClientId    *string `pulumi:"clientId"`
	PrincipalId *string `pulumi:"principalId"`
}





type UserIdentityPropertiesInput interface {
	pulumi.Input

	ToUserIdentityPropertiesOutput() UserIdentityPropertiesOutput
	ToUserIdentityPropertiesOutputWithContext(context.Context) UserIdentityPropertiesOutput
}

type UserIdentityPropertiesArgs struct {
	ClientId    pulumi.StringPtrInput `pulumi:"clientId"`
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
}

func (UserIdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityProperties)(nil)).Elem()
}

func (i UserIdentityPropertiesArgs) ToUserIdentityPropertiesOutput() UserIdentityPropertiesOutput {
	return i.ToUserIdentityPropertiesOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesArgs) ToUserIdentityPropertiesOutputWithContext(ctx context.Context) UserIdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesOutput)
}





type UserIdentityPropertiesMapInput interface {
	pulumi.Input

	ToUserIdentityPropertiesMapOutput() UserIdentityPropertiesMapOutput
	ToUserIdentityPropertiesMapOutputWithContext(context.Context) UserIdentityPropertiesMapOutput
}

type UserIdentityPropertiesMap map[string]UserIdentityPropertiesInput

func (UserIdentityPropertiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityProperties)(nil)).Elem()
}

func (i UserIdentityPropertiesMap) ToUserIdentityPropertiesMapOutput() UserIdentityPropertiesMapOutput {
	return i.ToUserIdentityPropertiesMapOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesMap) ToUserIdentityPropertiesMapOutputWithContext(ctx context.Context) UserIdentityPropertiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesMapOutput)
}

type UserIdentityPropertiesOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityProperties)(nil)).Elem()
}

func (o UserIdentityPropertiesOutput) ToUserIdentityPropertiesOutput() UserIdentityPropertiesOutput {
	return o
}

func (o UserIdentityPropertiesOutput) ToUserIdentityPropertiesOutputWithContext(ctx context.Context) UserIdentityPropertiesOutput {
	return o
}

func (o UserIdentityPropertiesOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityProperties) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserIdentityPropertiesOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityProperties) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type UserIdentityPropertiesMapOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityProperties)(nil)).Elem()
}

func (o UserIdentityPropertiesMapOutput) ToUserIdentityPropertiesMapOutput() UserIdentityPropertiesMapOutput {
	return o
}

func (o UserIdentityPropertiesMapOutput) ToUserIdentityPropertiesMapOutputWithContext(ctx context.Context) UserIdentityPropertiesMapOutput {
	return o
}

func (o UserIdentityPropertiesMapOutput) MapIndex(k pulumi.StringInput) UserIdentityPropertiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserIdentityProperties {
		return vs[0].(map[string]UserIdentityProperties)[vs[1].(string)]
	}).(UserIdentityPropertiesOutput)
}

type UserIdentityPropertiesResponse struct {
	ClientId    *string `pulumi:"clientId"`
	PrincipalId *string `pulumi:"principalId"`
}





type UserIdentityPropertiesResponseInput interface {
	pulumi.Input

	ToUserIdentityPropertiesResponseOutput() UserIdentityPropertiesResponseOutput
	ToUserIdentityPropertiesResponseOutputWithContext(context.Context) UserIdentityPropertiesResponseOutput
}

type UserIdentityPropertiesResponseArgs struct {
	ClientId    pulumi.StringPtrInput `pulumi:"clientId"`
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
}

func (UserIdentityPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityPropertiesResponse)(nil)).Elem()
}

func (i UserIdentityPropertiesResponseArgs) ToUserIdentityPropertiesResponseOutput() UserIdentityPropertiesResponseOutput {
	return i.ToUserIdentityPropertiesResponseOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesResponseArgs) ToUserIdentityPropertiesResponseOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesResponseOutput)
}





type UserIdentityPropertiesResponseMapInput interface {
	pulumi.Input

	ToUserIdentityPropertiesResponseMapOutput() UserIdentityPropertiesResponseMapOutput
	ToUserIdentityPropertiesResponseMapOutputWithContext(context.Context) UserIdentityPropertiesResponseMapOutput
}

type UserIdentityPropertiesResponseMap map[string]UserIdentityPropertiesResponseInput

func (UserIdentityPropertiesResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityPropertiesResponse)(nil)).Elem()
}

func (i UserIdentityPropertiesResponseMap) ToUserIdentityPropertiesResponseMapOutput() UserIdentityPropertiesResponseMapOutput {
	return i.ToUserIdentityPropertiesResponseMapOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesResponseMap) ToUserIdentityPropertiesResponseMapOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesResponseMapOutput)
}

type UserIdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserIdentityPropertiesResponseOutput) ToUserIdentityPropertiesResponseOutput() UserIdentityPropertiesResponseOutput {
	return o
}

func (o UserIdentityPropertiesResponseOutput) ToUserIdentityPropertiesResponseOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseOutput {
	return o
}

func (o UserIdentityPropertiesResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityPropertiesResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserIdentityPropertiesResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityPropertiesResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type UserIdentityPropertiesResponseMapOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserIdentityPropertiesResponseMapOutput) ToUserIdentityPropertiesResponseMapOutput() UserIdentityPropertiesResponseMapOutput {
	return o
}

func (o UserIdentityPropertiesResponseMapOutput) ToUserIdentityPropertiesResponseMapOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseMapOutput {
	return o
}

func (o UserIdentityPropertiesResponseMapOutput) MapIndex(k pulumi.StringInput) UserIdentityPropertiesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserIdentityPropertiesResponse {
		return vs[0].(map[string]UserIdentityPropertiesResponse)[vs[1].(string)]
	}).(UserIdentityPropertiesResponseOutput)
}

type WebHookEventSubscriptionDestination struct {
	AzureActiveDirectoryApplicationIdOrUri *string       `pulumi:"azureActiveDirectoryApplicationIdOrUri"`
	AzureActiveDirectoryTenantId           *string       `pulumi:"azureActiveDirectoryTenantId"`
	DeliveryAttributeMappings              []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType                           string        `pulumi:"endpointType"`
	EndpointUrl                            *string       `pulumi:"endpointUrl"`
	MaxEventsPerBatch                      *int          `pulumi:"maxEventsPerBatch"`
	PreferredBatchSizeInKilobytes          *int          `pulumi:"preferredBatchSizeInKilobytes"`
}





type WebHookEventSubscriptionDestinationInput interface {
	pulumi.Input

	ToWebHookEventSubscriptionDestinationOutput() WebHookEventSubscriptionDestinationOutput
	ToWebHookEventSubscriptionDestinationOutputWithContext(context.Context) WebHookEventSubscriptionDestinationOutput
}

type WebHookEventSubscriptionDestinationArgs struct {
	AzureActiveDirectoryApplicationIdOrUri pulumi.StringPtrInput `pulumi:"azureActiveDirectoryApplicationIdOrUri"`
	AzureActiveDirectoryTenantId           pulumi.StringPtrInput `pulumi:"azureActiveDirectoryTenantId"`
	DeliveryAttributeMappings              pulumi.ArrayInput     `pulumi:"deliveryAttributeMappings"`
	EndpointType                           pulumi.StringInput    `pulumi:"endpointType"`
	EndpointUrl                            pulumi.StringPtrInput `pulumi:"endpointUrl"`
	MaxEventsPerBatch                      pulumi.IntPtrInput    `pulumi:"maxEventsPerBatch"`
	PreferredBatchSizeInKilobytes          pulumi.IntPtrInput    `pulumi:"preferredBatchSizeInKilobytes"`
}

func (WebHookEventSubscriptionDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebHookEventSubscriptionDestination)(nil)).Elem()
}

func (i WebHookEventSubscriptionDestinationArgs) ToWebHookEventSubscriptionDestinationOutput() WebHookEventSubscriptionDestinationOutput {
	return i.ToWebHookEventSubscriptionDestinationOutputWithContext(context.Background())
}

func (i WebHookEventSubscriptionDestinationArgs) ToWebHookEventSubscriptionDestinationOutputWithContext(ctx context.Context) WebHookEventSubscriptionDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebHookEventSubscriptionDestinationOutput)
}

type WebHookEventSubscriptionDestinationOutput struct{ *pulumi.OutputState }

func (WebHookEventSubscriptionDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebHookEventSubscriptionDestination)(nil)).Elem()
}

func (o WebHookEventSubscriptionDestinationOutput) ToWebHookEventSubscriptionDestinationOutput() WebHookEventSubscriptionDestinationOutput {
	return o
}

func (o WebHookEventSubscriptionDestinationOutput) ToWebHookEventSubscriptionDestinationOutputWithContext(ctx context.Context) WebHookEventSubscriptionDestinationOutput {
	return o
}

func (o WebHookEventSubscriptionDestinationOutput) AzureActiveDirectoryApplicationIdOrUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestination) *string { return v.AzureActiveDirectoryApplicationIdOrUri }).(pulumi.StringPtrOutput)
}

func (o WebHookEventSubscriptionDestinationOutput) AzureActiveDirectoryTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestination) *string { return v.AzureActiveDirectoryTenantId }).(pulumi.StringPtrOutput)
}

func (o WebHookEventSubscriptionDestinationOutput) DeliveryAttributeMappings() pulumi.ArrayOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestination) []interface{} { return v.DeliveryAttributeMappings }).(pulumi.ArrayOutput)
}

func (o WebHookEventSubscriptionDestinationOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestination) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o WebHookEventSubscriptionDestinationOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestination) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func (o WebHookEventSubscriptionDestinationOutput) MaxEventsPerBatch() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestination) *int { return v.MaxEventsPerBatch }).(pulumi.IntPtrOutput)
}

func (o WebHookEventSubscriptionDestinationOutput) PreferredBatchSizeInKilobytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestination) *int { return v.PreferredBatchSizeInKilobytes }).(pulumi.IntPtrOutput)
}

type WebHookEventSubscriptionDestinationResponse struct {
	AzureActiveDirectoryApplicationIdOrUri *string       `pulumi:"azureActiveDirectoryApplicationIdOrUri"`
	AzureActiveDirectoryTenantId           *string       `pulumi:"azureActiveDirectoryTenantId"`
	DeliveryAttributeMappings              []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointBaseUrl                        string        `pulumi:"endpointBaseUrl"`
	EndpointType                           string        `pulumi:"endpointType"`
	EndpointUrl                            *string       `pulumi:"endpointUrl"`
	MaxEventsPerBatch                      *int          `pulumi:"maxEventsPerBatch"`
	PreferredBatchSizeInKilobytes          *int          `pulumi:"preferredBatchSizeInKilobytes"`
}





type WebHookEventSubscriptionDestinationResponseInput interface {
	pulumi.Input

	ToWebHookEventSubscriptionDestinationResponseOutput() WebHookEventSubscriptionDestinationResponseOutput
	ToWebHookEventSubscriptionDestinationResponseOutputWithContext(context.Context) WebHookEventSubscriptionDestinationResponseOutput
}

type WebHookEventSubscriptionDestinationResponseArgs struct {
	AzureActiveDirectoryApplicationIdOrUri pulumi.StringPtrInput `pulumi:"azureActiveDirectoryApplicationIdOrUri"`
	AzureActiveDirectoryTenantId           pulumi.StringPtrInput `pulumi:"azureActiveDirectoryTenantId"`
	DeliveryAttributeMappings              pulumi.ArrayInput     `pulumi:"deliveryAttributeMappings"`
	EndpointBaseUrl                        pulumi.StringInput    `pulumi:"endpointBaseUrl"`
	EndpointType                           pulumi.StringInput    `pulumi:"endpointType"`
	EndpointUrl                            pulumi.StringPtrInput `pulumi:"endpointUrl"`
	MaxEventsPerBatch                      pulumi.IntPtrInput    `pulumi:"maxEventsPerBatch"`
	PreferredBatchSizeInKilobytes          pulumi.IntPtrInput    `pulumi:"preferredBatchSizeInKilobytes"`
}

func (WebHookEventSubscriptionDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebHookEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (i WebHookEventSubscriptionDestinationResponseArgs) ToWebHookEventSubscriptionDestinationResponseOutput() WebHookEventSubscriptionDestinationResponseOutput {
	return i.ToWebHookEventSubscriptionDestinationResponseOutputWithContext(context.Background())
}

func (i WebHookEventSubscriptionDestinationResponseArgs) ToWebHookEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) WebHookEventSubscriptionDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebHookEventSubscriptionDestinationResponseOutput)
}

type WebHookEventSubscriptionDestinationResponseOutput struct{ *pulumi.OutputState }

func (WebHookEventSubscriptionDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebHookEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (o WebHookEventSubscriptionDestinationResponseOutput) ToWebHookEventSubscriptionDestinationResponseOutput() WebHookEventSubscriptionDestinationResponseOutput {
	return o
}

func (o WebHookEventSubscriptionDestinationResponseOutput) ToWebHookEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) WebHookEventSubscriptionDestinationResponseOutput {
	return o
}

func (o WebHookEventSubscriptionDestinationResponseOutput) AzureActiveDirectoryApplicationIdOrUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestinationResponse) *string {
		return v.AzureActiveDirectoryApplicationIdOrUri
	}).(pulumi.StringPtrOutput)
}

func (o WebHookEventSubscriptionDestinationResponseOutput) AzureActiveDirectoryTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestinationResponse) *string { return v.AzureActiveDirectoryTenantId }).(pulumi.StringPtrOutput)
}

func (o WebHookEventSubscriptionDestinationResponseOutput) DeliveryAttributeMappings() pulumi.ArrayOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestinationResponse) []interface{} { return v.DeliveryAttributeMappings }).(pulumi.ArrayOutput)
}

func (o WebHookEventSubscriptionDestinationResponseOutput) EndpointBaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestinationResponse) string { return v.EndpointBaseUrl }).(pulumi.StringOutput)
}

func (o WebHookEventSubscriptionDestinationResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestinationResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o WebHookEventSubscriptionDestinationResponseOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestinationResponse) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func (o WebHookEventSubscriptionDestinationResponseOutput) MaxEventsPerBatch() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestinationResponse) *int { return v.MaxEventsPerBatch }).(pulumi.IntPtrOutput)
}

func (o WebHookEventSubscriptionDestinationResponseOutput) PreferredBatchSizeInKilobytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestinationResponse) *int { return v.PreferredBatchSizeInKilobytes }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureFunctionEventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(AzureFunctionEventSubscriptionDestinationResponseOutput{})
	pulumi.RegisterOutputType(BoolEqualsAdvancedFilterOutput{})
	pulumi.RegisterOutputType(BoolEqualsAdvancedFilterResponseOutput{})
	pulumi.RegisterOutputType(ConnectionStateOutput{})
	pulumi.RegisterOutputType(ConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(ConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(ConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(DeadLetterWithResourceIdentityOutput{})
	pulumi.RegisterOutputType(DeadLetterWithResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(DeadLetterWithResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(DeadLetterWithResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(DeliveryWithResourceIdentityOutput{})
	pulumi.RegisterOutputType(DeliveryWithResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(DeliveryWithResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(DeliveryWithResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(DynamicDeliveryAttributeMappingOutput{})
	pulumi.RegisterOutputType(DynamicDeliveryAttributeMappingResponseOutput{})
	pulumi.RegisterOutputType(EventHubEventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(EventHubEventSubscriptionDestinationResponseOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterPtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterResponseOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionIdentityOutput{})
	pulumi.RegisterOutputType(EventSubscriptionIdentityPtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionIdentityResponseOutput{})
	pulumi.RegisterOutputType(EventSubscriptionIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(HybridConnectionEventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(HybridConnectionEventSubscriptionDestinationResponseOutput{})
	pulumi.RegisterOutputType(IdentityInfoOutput{})
	pulumi.RegisterOutputType(IdentityInfoPtrOutput{})
	pulumi.RegisterOutputType(IdentityInfoResponseOutput{})
	pulumi.RegisterOutputType(IdentityInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(InboundIpRuleOutput{})
	pulumi.RegisterOutputType(InboundIpRuleArrayOutput{})
	pulumi.RegisterOutputType(InboundIpRuleResponseOutput{})
	pulumi.RegisterOutputType(InboundIpRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(JsonFieldOutput{})
	pulumi.RegisterOutputType(JsonFieldPtrOutput{})
	pulumi.RegisterOutputType(JsonFieldResponseOutput{})
	pulumi.RegisterOutputType(JsonFieldResponsePtrOutput{})
	pulumi.RegisterOutputType(JsonFieldWithDefaultOutput{})
	pulumi.RegisterOutputType(JsonFieldWithDefaultPtrOutput{})
	pulumi.RegisterOutputType(JsonFieldWithDefaultResponseOutput{})
	pulumi.RegisterOutputType(JsonFieldWithDefaultResponsePtrOutput{})
	pulumi.RegisterOutputType(JsonInputSchemaMappingOutput{})
	pulumi.RegisterOutputType(JsonInputSchemaMappingPtrOutput{})
	pulumi.RegisterOutputType(JsonInputSchemaMappingResponseOutput{})
	pulumi.RegisterOutputType(JsonInputSchemaMappingResponsePtrOutput{})
	pulumi.RegisterOutputType(NumberGreaterThanAdvancedFilterOutput{})
	pulumi.RegisterOutputType(NumberGreaterThanAdvancedFilterResponseOutput{})
	pulumi.RegisterOutputType(NumberGreaterThanOrEqualsAdvancedFilterOutput{})
	pulumi.RegisterOutputType(NumberGreaterThanOrEqualsAdvancedFilterResponseOutput{})
	pulumi.RegisterOutputType(NumberInAdvancedFilterOutput{})
	pulumi.RegisterOutputType(NumberInAdvancedFilterResponseOutput{})
	pulumi.RegisterOutputType(NumberLessThanAdvancedFilterOutput{})
	pulumi.RegisterOutputType(NumberLessThanAdvancedFilterResponseOutput{})
	pulumi.RegisterOutputType(NumberLessThanOrEqualsAdvancedFilterOutput{})
	pulumi.RegisterOutputType(NumberLessThanOrEqualsAdvancedFilterResponseOutput{})
	pulumi.RegisterOutputType(NumberNotInAdvancedFilterOutput{})
	pulumi.RegisterOutputType(NumberNotInAdvancedFilterResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(RetryPolicyOutput{})
	pulumi.RegisterOutputType(RetryPolicyPtrOutput{})
	pulumi.RegisterOutputType(RetryPolicyResponseOutput{})
	pulumi.RegisterOutputType(RetryPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceBusQueueEventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(ServiceBusQueueEventSubscriptionDestinationResponseOutput{})
	pulumi.RegisterOutputType(ServiceBusTopicEventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(ServiceBusTopicEventSubscriptionDestinationResponseOutput{})
	pulumi.RegisterOutputType(StaticDeliveryAttributeMappingOutput{})
	pulumi.RegisterOutputType(StaticDeliveryAttributeMappingResponseOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationPtrOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationResponseOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageQueueEventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(StorageQueueEventSubscriptionDestinationResponseOutput{})
	pulumi.RegisterOutputType(StringBeginsWithAdvancedFilterOutput{})
	pulumi.RegisterOutputType(StringBeginsWithAdvancedFilterResponseOutput{})
	pulumi.RegisterOutputType(StringContainsAdvancedFilterOutput{})
	pulumi.RegisterOutputType(StringContainsAdvancedFilterResponseOutput{})
	pulumi.RegisterOutputType(StringEndsWithAdvancedFilterOutput{})
	pulumi.RegisterOutputType(StringEndsWithAdvancedFilterResponseOutput{})
	pulumi.RegisterOutputType(StringInAdvancedFilterOutput{})
	pulumi.RegisterOutputType(StringInAdvancedFilterResponseOutput{})
	pulumi.RegisterOutputType(StringNotInAdvancedFilterOutput{})
	pulumi.RegisterOutputType(StringNotInAdvancedFilterResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesMapOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseMapOutput{})
	pulumi.RegisterOutputType(WebHookEventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(WebHookEventSubscriptionDestinationResponseOutput{})
}

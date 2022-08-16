


package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureADPartnerClientAuthentication struct {
	AzureActiveDirectoryApplicationIdOrUri *string `pulumi:"azureActiveDirectoryApplicationIdOrUri"`
	AzureActiveDirectoryTenantId           *string `pulumi:"azureActiveDirectoryTenantId"`
	ClientAuthenticationType               string  `pulumi:"clientAuthenticationType"`
}


func (val *AzureADPartnerClientAuthentication) Defaults() *AzureADPartnerClientAuthentication {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClientAuthenticationType) {
		tmp.ClientAuthenticationType = "AzureAD"
	}
	return &tmp
}





type AzureADPartnerClientAuthenticationInput interface {
	pulumi.Input

	ToAzureADPartnerClientAuthenticationOutput() AzureADPartnerClientAuthenticationOutput
	ToAzureADPartnerClientAuthenticationOutputWithContext(context.Context) AzureADPartnerClientAuthenticationOutput
}

type AzureADPartnerClientAuthenticationArgs struct {
	AzureActiveDirectoryApplicationIdOrUri pulumi.StringPtrInput `pulumi:"azureActiveDirectoryApplicationIdOrUri"`
	AzureActiveDirectoryTenantId           pulumi.StringPtrInput `pulumi:"azureActiveDirectoryTenantId"`
	ClientAuthenticationType               pulumi.StringInput    `pulumi:"clientAuthenticationType"`
}


func (val *AzureADPartnerClientAuthenticationArgs) Defaults() *AzureADPartnerClientAuthenticationArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClientAuthenticationType) {
		tmp.ClientAuthenticationType = pulumi.String("AzureAD")
	}
	return &tmp
}
func (AzureADPartnerClientAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureADPartnerClientAuthentication)(nil)).Elem()
}

func (i AzureADPartnerClientAuthenticationArgs) ToAzureADPartnerClientAuthenticationOutput() AzureADPartnerClientAuthenticationOutput {
	return i.ToAzureADPartnerClientAuthenticationOutputWithContext(context.Background())
}

func (i AzureADPartnerClientAuthenticationArgs) ToAzureADPartnerClientAuthenticationOutputWithContext(ctx context.Context) AzureADPartnerClientAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureADPartnerClientAuthenticationOutput)
}

func (i AzureADPartnerClientAuthenticationArgs) ToAzureADPartnerClientAuthenticationPtrOutput() AzureADPartnerClientAuthenticationPtrOutput {
	return i.ToAzureADPartnerClientAuthenticationPtrOutputWithContext(context.Background())
}

func (i AzureADPartnerClientAuthenticationArgs) ToAzureADPartnerClientAuthenticationPtrOutputWithContext(ctx context.Context) AzureADPartnerClientAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureADPartnerClientAuthenticationOutput).ToAzureADPartnerClientAuthenticationPtrOutputWithContext(ctx)
}









type AzureADPartnerClientAuthenticationPtrInput interface {
	pulumi.Input

	ToAzureADPartnerClientAuthenticationPtrOutput() AzureADPartnerClientAuthenticationPtrOutput
	ToAzureADPartnerClientAuthenticationPtrOutputWithContext(context.Context) AzureADPartnerClientAuthenticationPtrOutput
}

type azureADPartnerClientAuthenticationPtrType AzureADPartnerClientAuthenticationArgs

func AzureADPartnerClientAuthenticationPtr(v *AzureADPartnerClientAuthenticationArgs) AzureADPartnerClientAuthenticationPtrInput {
	return (*azureADPartnerClientAuthenticationPtrType)(v)
}

func (*azureADPartnerClientAuthenticationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureADPartnerClientAuthentication)(nil)).Elem()
}

func (i *azureADPartnerClientAuthenticationPtrType) ToAzureADPartnerClientAuthenticationPtrOutput() AzureADPartnerClientAuthenticationPtrOutput {
	return i.ToAzureADPartnerClientAuthenticationPtrOutputWithContext(context.Background())
}

func (i *azureADPartnerClientAuthenticationPtrType) ToAzureADPartnerClientAuthenticationPtrOutputWithContext(ctx context.Context) AzureADPartnerClientAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureADPartnerClientAuthenticationPtrOutput)
}

type AzureADPartnerClientAuthenticationOutput struct{ *pulumi.OutputState }

func (AzureADPartnerClientAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureADPartnerClientAuthentication)(nil)).Elem()
}

func (o AzureADPartnerClientAuthenticationOutput) ToAzureADPartnerClientAuthenticationOutput() AzureADPartnerClientAuthenticationOutput {
	return o
}

func (o AzureADPartnerClientAuthenticationOutput) ToAzureADPartnerClientAuthenticationOutputWithContext(ctx context.Context) AzureADPartnerClientAuthenticationOutput {
	return o
}

func (o AzureADPartnerClientAuthenticationOutput) ToAzureADPartnerClientAuthenticationPtrOutput() AzureADPartnerClientAuthenticationPtrOutput {
	return o.ToAzureADPartnerClientAuthenticationPtrOutputWithContext(context.Background())
}

func (o AzureADPartnerClientAuthenticationOutput) ToAzureADPartnerClientAuthenticationPtrOutputWithContext(ctx context.Context) AzureADPartnerClientAuthenticationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureADPartnerClientAuthentication) *AzureADPartnerClientAuthentication {
		return &v
	}).(AzureADPartnerClientAuthenticationPtrOutput)
}

func (o AzureADPartnerClientAuthenticationOutput) AzureActiveDirectoryApplicationIdOrUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureADPartnerClientAuthentication) *string { return v.AzureActiveDirectoryApplicationIdOrUri }).(pulumi.StringPtrOutput)
}

func (o AzureADPartnerClientAuthenticationOutput) AzureActiveDirectoryTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureADPartnerClientAuthentication) *string { return v.AzureActiveDirectoryTenantId }).(pulumi.StringPtrOutput)
}

func (o AzureADPartnerClientAuthenticationOutput) ClientAuthenticationType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureADPartnerClientAuthentication) string { return v.ClientAuthenticationType }).(pulumi.StringOutput)
}

type AzureADPartnerClientAuthenticationPtrOutput struct{ *pulumi.OutputState }

func (AzureADPartnerClientAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureADPartnerClientAuthentication)(nil)).Elem()
}

func (o AzureADPartnerClientAuthenticationPtrOutput) ToAzureADPartnerClientAuthenticationPtrOutput() AzureADPartnerClientAuthenticationPtrOutput {
	return o
}

func (o AzureADPartnerClientAuthenticationPtrOutput) ToAzureADPartnerClientAuthenticationPtrOutputWithContext(ctx context.Context) AzureADPartnerClientAuthenticationPtrOutput {
	return o
}

func (o AzureADPartnerClientAuthenticationPtrOutput) Elem() AzureADPartnerClientAuthenticationOutput {
	return o.ApplyT(func(v *AzureADPartnerClientAuthentication) AzureADPartnerClientAuthentication {
		if v != nil {
			return *v
		}
		var ret AzureADPartnerClientAuthentication
		return ret
	}).(AzureADPartnerClientAuthenticationOutput)
}

func (o AzureADPartnerClientAuthenticationPtrOutput) AzureActiveDirectoryApplicationIdOrUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADPartnerClientAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.AzureActiveDirectoryApplicationIdOrUri
	}).(pulumi.StringPtrOutput)
}

func (o AzureADPartnerClientAuthenticationPtrOutput) AzureActiveDirectoryTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADPartnerClientAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.AzureActiveDirectoryTenantId
	}).(pulumi.StringPtrOutput)
}

func (o AzureADPartnerClientAuthenticationPtrOutput) ClientAuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADPartnerClientAuthentication) *string {
		if v == nil {
			return nil
		}
		return &v.ClientAuthenticationType
	}).(pulumi.StringPtrOutput)
}

type AzureADPartnerClientAuthenticationResponse struct {
	AzureActiveDirectoryApplicationIdOrUri *string `pulumi:"azureActiveDirectoryApplicationIdOrUri"`
	AzureActiveDirectoryTenantId           *string `pulumi:"azureActiveDirectoryTenantId"`
	ClientAuthenticationType               string  `pulumi:"clientAuthenticationType"`
}


func (val *AzureADPartnerClientAuthenticationResponse) Defaults() *AzureADPartnerClientAuthenticationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClientAuthenticationType) {
		tmp.ClientAuthenticationType = "AzureAD"
	}
	return &tmp
}

type AzureADPartnerClientAuthenticationResponseOutput struct{ *pulumi.OutputState }

func (AzureADPartnerClientAuthenticationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureADPartnerClientAuthenticationResponse)(nil)).Elem()
}

func (o AzureADPartnerClientAuthenticationResponseOutput) ToAzureADPartnerClientAuthenticationResponseOutput() AzureADPartnerClientAuthenticationResponseOutput {
	return o
}

func (o AzureADPartnerClientAuthenticationResponseOutput) ToAzureADPartnerClientAuthenticationResponseOutputWithContext(ctx context.Context) AzureADPartnerClientAuthenticationResponseOutput {
	return o
}

func (o AzureADPartnerClientAuthenticationResponseOutput) AzureActiveDirectoryApplicationIdOrUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureADPartnerClientAuthenticationResponse) *string {
		return v.AzureActiveDirectoryApplicationIdOrUri
	}).(pulumi.StringPtrOutput)
}

func (o AzureADPartnerClientAuthenticationResponseOutput) AzureActiveDirectoryTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureADPartnerClientAuthenticationResponse) *string { return v.AzureActiveDirectoryTenantId }).(pulumi.StringPtrOutput)
}

func (o AzureADPartnerClientAuthenticationResponseOutput) ClientAuthenticationType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureADPartnerClientAuthenticationResponse) string { return v.ClientAuthenticationType }).(pulumi.StringOutput)
}

type AzureADPartnerClientAuthenticationResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureADPartnerClientAuthenticationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureADPartnerClientAuthenticationResponse)(nil)).Elem()
}

func (o AzureADPartnerClientAuthenticationResponsePtrOutput) ToAzureADPartnerClientAuthenticationResponsePtrOutput() AzureADPartnerClientAuthenticationResponsePtrOutput {
	return o
}

func (o AzureADPartnerClientAuthenticationResponsePtrOutput) ToAzureADPartnerClientAuthenticationResponsePtrOutputWithContext(ctx context.Context) AzureADPartnerClientAuthenticationResponsePtrOutput {
	return o
}

func (o AzureADPartnerClientAuthenticationResponsePtrOutput) Elem() AzureADPartnerClientAuthenticationResponseOutput {
	return o.ApplyT(func(v *AzureADPartnerClientAuthenticationResponse) AzureADPartnerClientAuthenticationResponse {
		if v != nil {
			return *v
		}
		var ret AzureADPartnerClientAuthenticationResponse
		return ret
	}).(AzureADPartnerClientAuthenticationResponseOutput)
}

func (o AzureADPartnerClientAuthenticationResponsePtrOutput) AzureActiveDirectoryApplicationIdOrUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADPartnerClientAuthenticationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AzureActiveDirectoryApplicationIdOrUri
	}).(pulumi.StringPtrOutput)
}

func (o AzureADPartnerClientAuthenticationResponsePtrOutput) AzureActiveDirectoryTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADPartnerClientAuthenticationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AzureActiveDirectoryTenantId
	}).(pulumi.StringPtrOutput)
}

func (o AzureADPartnerClientAuthenticationResponsePtrOutput) ClientAuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADPartnerClientAuthenticationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClientAuthenticationType
	}).(pulumi.StringPtrOutput)
}

type AzureFunctionEventSubscriptionDestination struct {
	DeliveryAttributeMappings     []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType                  string        `pulumi:"endpointType"`
	MaxEventsPerBatch             *int          `pulumi:"maxEventsPerBatch"`
	PreferredBatchSizeInKilobytes *int          `pulumi:"preferredBatchSizeInKilobytes"`
	ResourceId                    *string       `pulumi:"resourceId"`
}


func (val *AzureFunctionEventSubscriptionDestination) Defaults() *AzureFunctionEventSubscriptionDestination {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxEventsPerBatch) {
		maxEventsPerBatch_ := 1
		tmp.MaxEventsPerBatch = &maxEventsPerBatch_
	}
	if isZero(tmp.PreferredBatchSizeInKilobytes) {
		preferredBatchSizeInKilobytes_ := 64
		tmp.PreferredBatchSizeInKilobytes = &preferredBatchSizeInKilobytes_
	}
	return &tmp
}

type AzureFunctionEventSubscriptionDestinationResponse struct {
	DeliveryAttributeMappings     []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType                  string        `pulumi:"endpointType"`
	MaxEventsPerBatch             *int          `pulumi:"maxEventsPerBatch"`
	PreferredBatchSizeInKilobytes *int          `pulumi:"preferredBatchSizeInKilobytes"`
	ResourceId                    *string       `pulumi:"resourceId"`
}


func (val *AzureFunctionEventSubscriptionDestinationResponse) Defaults() *AzureFunctionEventSubscriptionDestinationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxEventsPerBatch) {
		maxEventsPerBatch_ := 1
		tmp.MaxEventsPerBatch = &maxEventsPerBatch_
	}
	if isZero(tmp.PreferredBatchSizeInKilobytes) {
		preferredBatchSizeInKilobytes_ := 64
		tmp.PreferredBatchSizeInKilobytes = &preferredBatchSizeInKilobytes_
	}
	return &tmp
}

type BoolEqualsAdvancedFilter struct {
	Key          *string `pulumi:"key"`
	OperatorType string  `pulumi:"operatorType"`
	Value        *bool   `pulumi:"value"`
}

type BoolEqualsAdvancedFilterResponse struct {
	Key          *string `pulumi:"key"`
	OperatorType string  `pulumi:"operatorType"`
	Value        *bool   `pulumi:"value"`
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

type DynamicDeliveryAttributeMappingResponse struct {
	Name        *string `pulumi:"name"`
	SourceField *string `pulumi:"sourceField"`
	Type        string  `pulumi:"type"`
}

type EventChannelDestination struct {
	AzureSubscriptionId *string `pulumi:"azureSubscriptionId"`
	PartnerTopicName    *string `pulumi:"partnerTopicName"`
	ResourceGroup       *string `pulumi:"resourceGroup"`
}





type EventChannelDestinationInput interface {
	pulumi.Input

	ToEventChannelDestinationOutput() EventChannelDestinationOutput
	ToEventChannelDestinationOutputWithContext(context.Context) EventChannelDestinationOutput
}

type EventChannelDestinationArgs struct {
	AzureSubscriptionId pulumi.StringPtrInput `pulumi:"azureSubscriptionId"`
	PartnerTopicName    pulumi.StringPtrInput `pulumi:"partnerTopicName"`
	ResourceGroup       pulumi.StringPtrInput `pulumi:"resourceGroup"`
}

func (EventChannelDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventChannelDestination)(nil)).Elem()
}

func (i EventChannelDestinationArgs) ToEventChannelDestinationOutput() EventChannelDestinationOutput {
	return i.ToEventChannelDestinationOutputWithContext(context.Background())
}

func (i EventChannelDestinationArgs) ToEventChannelDestinationOutputWithContext(ctx context.Context) EventChannelDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventChannelDestinationOutput)
}

func (i EventChannelDestinationArgs) ToEventChannelDestinationPtrOutput() EventChannelDestinationPtrOutput {
	return i.ToEventChannelDestinationPtrOutputWithContext(context.Background())
}

func (i EventChannelDestinationArgs) ToEventChannelDestinationPtrOutputWithContext(ctx context.Context) EventChannelDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventChannelDestinationOutput).ToEventChannelDestinationPtrOutputWithContext(ctx)
}









type EventChannelDestinationPtrInput interface {
	pulumi.Input

	ToEventChannelDestinationPtrOutput() EventChannelDestinationPtrOutput
	ToEventChannelDestinationPtrOutputWithContext(context.Context) EventChannelDestinationPtrOutput
}

type eventChannelDestinationPtrType EventChannelDestinationArgs

func EventChannelDestinationPtr(v *EventChannelDestinationArgs) EventChannelDestinationPtrInput {
	return (*eventChannelDestinationPtrType)(v)
}

func (*eventChannelDestinationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventChannelDestination)(nil)).Elem()
}

func (i *eventChannelDestinationPtrType) ToEventChannelDestinationPtrOutput() EventChannelDestinationPtrOutput {
	return i.ToEventChannelDestinationPtrOutputWithContext(context.Background())
}

func (i *eventChannelDestinationPtrType) ToEventChannelDestinationPtrOutputWithContext(ctx context.Context) EventChannelDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventChannelDestinationPtrOutput)
}

type EventChannelDestinationOutput struct{ *pulumi.OutputState }

func (EventChannelDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventChannelDestination)(nil)).Elem()
}

func (o EventChannelDestinationOutput) ToEventChannelDestinationOutput() EventChannelDestinationOutput {
	return o
}

func (o EventChannelDestinationOutput) ToEventChannelDestinationOutputWithContext(ctx context.Context) EventChannelDestinationOutput {
	return o
}

func (o EventChannelDestinationOutput) ToEventChannelDestinationPtrOutput() EventChannelDestinationPtrOutput {
	return o.ToEventChannelDestinationPtrOutputWithContext(context.Background())
}

func (o EventChannelDestinationOutput) ToEventChannelDestinationPtrOutputWithContext(ctx context.Context) EventChannelDestinationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventChannelDestination) *EventChannelDestination {
		return &v
	}).(EventChannelDestinationPtrOutput)
}

func (o EventChannelDestinationOutput) AzureSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventChannelDestination) *string { return v.AzureSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o EventChannelDestinationOutput) PartnerTopicName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventChannelDestination) *string { return v.PartnerTopicName }).(pulumi.StringPtrOutput)
}

func (o EventChannelDestinationOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventChannelDestination) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

type EventChannelDestinationPtrOutput struct{ *pulumi.OutputState }

func (EventChannelDestinationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventChannelDestination)(nil)).Elem()
}

func (o EventChannelDestinationPtrOutput) ToEventChannelDestinationPtrOutput() EventChannelDestinationPtrOutput {
	return o
}

func (o EventChannelDestinationPtrOutput) ToEventChannelDestinationPtrOutputWithContext(ctx context.Context) EventChannelDestinationPtrOutput {
	return o
}

func (o EventChannelDestinationPtrOutput) Elem() EventChannelDestinationOutput {
	return o.ApplyT(func(v *EventChannelDestination) EventChannelDestination {
		if v != nil {
			return *v
		}
		var ret EventChannelDestination
		return ret
	}).(EventChannelDestinationOutput)
}

func (o EventChannelDestinationPtrOutput) AzureSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventChannelDestination) *string {
		if v == nil {
			return nil
		}
		return v.AzureSubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o EventChannelDestinationPtrOutput) PartnerTopicName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventChannelDestination) *string {
		if v == nil {
			return nil
		}
		return v.PartnerTopicName
	}).(pulumi.StringPtrOutput)
}

func (o EventChannelDestinationPtrOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventChannelDestination) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroup
	}).(pulumi.StringPtrOutput)
}

type EventChannelDestinationResponse struct {
	AzureSubscriptionId *string `pulumi:"azureSubscriptionId"`
	PartnerTopicName    *string `pulumi:"partnerTopicName"`
	ResourceGroup       *string `pulumi:"resourceGroup"`
}

type EventChannelDestinationResponseOutput struct{ *pulumi.OutputState }

func (EventChannelDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventChannelDestinationResponse)(nil)).Elem()
}

func (o EventChannelDestinationResponseOutput) ToEventChannelDestinationResponseOutput() EventChannelDestinationResponseOutput {
	return o
}

func (o EventChannelDestinationResponseOutput) ToEventChannelDestinationResponseOutputWithContext(ctx context.Context) EventChannelDestinationResponseOutput {
	return o
}

func (o EventChannelDestinationResponseOutput) AzureSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventChannelDestinationResponse) *string { return v.AzureSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o EventChannelDestinationResponseOutput) PartnerTopicName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventChannelDestinationResponse) *string { return v.PartnerTopicName }).(pulumi.StringPtrOutput)
}

func (o EventChannelDestinationResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventChannelDestinationResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

type EventChannelDestinationResponsePtrOutput struct{ *pulumi.OutputState }

func (EventChannelDestinationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventChannelDestinationResponse)(nil)).Elem()
}

func (o EventChannelDestinationResponsePtrOutput) ToEventChannelDestinationResponsePtrOutput() EventChannelDestinationResponsePtrOutput {
	return o
}

func (o EventChannelDestinationResponsePtrOutput) ToEventChannelDestinationResponsePtrOutputWithContext(ctx context.Context) EventChannelDestinationResponsePtrOutput {
	return o
}

func (o EventChannelDestinationResponsePtrOutput) Elem() EventChannelDestinationResponseOutput {
	return o.ApplyT(func(v *EventChannelDestinationResponse) EventChannelDestinationResponse {
		if v != nil {
			return *v
		}
		var ret EventChannelDestinationResponse
		return ret
	}).(EventChannelDestinationResponseOutput)
}

func (o EventChannelDestinationResponsePtrOutput) AzureSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventChannelDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AzureSubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o EventChannelDestinationResponsePtrOutput) PartnerTopicName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventChannelDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PartnerTopicName
	}).(pulumi.StringPtrOutput)
}

func (o EventChannelDestinationResponsePtrOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventChannelDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroup
	}).(pulumi.StringPtrOutput)
}

type EventChannelFilter struct {
	AdvancedFilters                 []interface{} `pulumi:"advancedFilters"`
	EnableAdvancedFilteringOnArrays *bool         `pulumi:"enableAdvancedFilteringOnArrays"`
}


func (val *EventChannelFilter) Defaults() *EventChannelFilter {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableAdvancedFilteringOnArrays) {
		enableAdvancedFilteringOnArrays_ := false
		tmp.EnableAdvancedFilteringOnArrays = &enableAdvancedFilteringOnArrays_
	}
	return &tmp
}





type EventChannelFilterInput interface {
	pulumi.Input

	ToEventChannelFilterOutput() EventChannelFilterOutput
	ToEventChannelFilterOutputWithContext(context.Context) EventChannelFilterOutput
}

type EventChannelFilterArgs struct {
	AdvancedFilters                 pulumi.ArrayInput   `pulumi:"advancedFilters"`
	EnableAdvancedFilteringOnArrays pulumi.BoolPtrInput `pulumi:"enableAdvancedFilteringOnArrays"`
}


func (val *EventChannelFilterArgs) Defaults() *EventChannelFilterArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableAdvancedFilteringOnArrays) {
		tmp.EnableAdvancedFilteringOnArrays = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (EventChannelFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventChannelFilter)(nil)).Elem()
}

func (i EventChannelFilterArgs) ToEventChannelFilterOutput() EventChannelFilterOutput {
	return i.ToEventChannelFilterOutputWithContext(context.Background())
}

func (i EventChannelFilterArgs) ToEventChannelFilterOutputWithContext(ctx context.Context) EventChannelFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventChannelFilterOutput)
}

func (i EventChannelFilterArgs) ToEventChannelFilterPtrOutput() EventChannelFilterPtrOutput {
	return i.ToEventChannelFilterPtrOutputWithContext(context.Background())
}

func (i EventChannelFilterArgs) ToEventChannelFilterPtrOutputWithContext(ctx context.Context) EventChannelFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventChannelFilterOutput).ToEventChannelFilterPtrOutputWithContext(ctx)
}









type EventChannelFilterPtrInput interface {
	pulumi.Input

	ToEventChannelFilterPtrOutput() EventChannelFilterPtrOutput
	ToEventChannelFilterPtrOutputWithContext(context.Context) EventChannelFilterPtrOutput
}

type eventChannelFilterPtrType EventChannelFilterArgs

func EventChannelFilterPtr(v *EventChannelFilterArgs) EventChannelFilterPtrInput {
	return (*eventChannelFilterPtrType)(v)
}

func (*eventChannelFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventChannelFilter)(nil)).Elem()
}

func (i *eventChannelFilterPtrType) ToEventChannelFilterPtrOutput() EventChannelFilterPtrOutput {
	return i.ToEventChannelFilterPtrOutputWithContext(context.Background())
}

func (i *eventChannelFilterPtrType) ToEventChannelFilterPtrOutputWithContext(ctx context.Context) EventChannelFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventChannelFilterPtrOutput)
}

type EventChannelFilterOutput struct{ *pulumi.OutputState }

func (EventChannelFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventChannelFilter)(nil)).Elem()
}

func (o EventChannelFilterOutput) ToEventChannelFilterOutput() EventChannelFilterOutput {
	return o
}

func (o EventChannelFilterOutput) ToEventChannelFilterOutputWithContext(ctx context.Context) EventChannelFilterOutput {
	return o
}

func (o EventChannelFilterOutput) ToEventChannelFilterPtrOutput() EventChannelFilterPtrOutput {
	return o.ToEventChannelFilterPtrOutputWithContext(context.Background())
}

func (o EventChannelFilterOutput) ToEventChannelFilterPtrOutputWithContext(ctx context.Context) EventChannelFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventChannelFilter) *EventChannelFilter {
		return &v
	}).(EventChannelFilterPtrOutput)
}

func (o EventChannelFilterOutput) AdvancedFilters() pulumi.ArrayOutput {
	return o.ApplyT(func(v EventChannelFilter) []interface{} { return v.AdvancedFilters }).(pulumi.ArrayOutput)
}

func (o EventChannelFilterOutput) EnableAdvancedFilteringOnArrays() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EventChannelFilter) *bool { return v.EnableAdvancedFilteringOnArrays }).(pulumi.BoolPtrOutput)
}

type EventChannelFilterPtrOutput struct{ *pulumi.OutputState }

func (EventChannelFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventChannelFilter)(nil)).Elem()
}

func (o EventChannelFilterPtrOutput) ToEventChannelFilterPtrOutput() EventChannelFilterPtrOutput {
	return o
}

func (o EventChannelFilterPtrOutput) ToEventChannelFilterPtrOutputWithContext(ctx context.Context) EventChannelFilterPtrOutput {
	return o
}

func (o EventChannelFilterPtrOutput) Elem() EventChannelFilterOutput {
	return o.ApplyT(func(v *EventChannelFilter) EventChannelFilter {
		if v != nil {
			return *v
		}
		var ret EventChannelFilter
		return ret
	}).(EventChannelFilterOutput)
}

func (o EventChannelFilterPtrOutput) AdvancedFilters() pulumi.ArrayOutput {
	return o.ApplyT(func(v *EventChannelFilter) []interface{} {
		if v == nil {
			return nil
		}
		return v.AdvancedFilters
	}).(pulumi.ArrayOutput)
}

func (o EventChannelFilterPtrOutput) EnableAdvancedFilteringOnArrays() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventChannelFilter) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAdvancedFilteringOnArrays
	}).(pulumi.BoolPtrOutput)
}

type EventChannelFilterResponse struct {
	AdvancedFilters                 []interface{} `pulumi:"advancedFilters"`
	EnableAdvancedFilteringOnArrays *bool         `pulumi:"enableAdvancedFilteringOnArrays"`
}


func (val *EventChannelFilterResponse) Defaults() *EventChannelFilterResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableAdvancedFilteringOnArrays) {
		enableAdvancedFilteringOnArrays_ := false
		tmp.EnableAdvancedFilteringOnArrays = &enableAdvancedFilteringOnArrays_
	}
	return &tmp
}

type EventChannelFilterResponseOutput struct{ *pulumi.OutputState }

func (EventChannelFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventChannelFilterResponse)(nil)).Elem()
}

func (o EventChannelFilterResponseOutput) ToEventChannelFilterResponseOutput() EventChannelFilterResponseOutput {
	return o
}

func (o EventChannelFilterResponseOutput) ToEventChannelFilterResponseOutputWithContext(ctx context.Context) EventChannelFilterResponseOutput {
	return o
}

func (o EventChannelFilterResponseOutput) AdvancedFilters() pulumi.ArrayOutput {
	return o.ApplyT(func(v EventChannelFilterResponse) []interface{} { return v.AdvancedFilters }).(pulumi.ArrayOutput)
}

func (o EventChannelFilterResponseOutput) EnableAdvancedFilteringOnArrays() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EventChannelFilterResponse) *bool { return v.EnableAdvancedFilteringOnArrays }).(pulumi.BoolPtrOutput)
}

type EventChannelFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (EventChannelFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventChannelFilterResponse)(nil)).Elem()
}

func (o EventChannelFilterResponsePtrOutput) ToEventChannelFilterResponsePtrOutput() EventChannelFilterResponsePtrOutput {
	return o
}

func (o EventChannelFilterResponsePtrOutput) ToEventChannelFilterResponsePtrOutputWithContext(ctx context.Context) EventChannelFilterResponsePtrOutput {
	return o
}

func (o EventChannelFilterResponsePtrOutput) Elem() EventChannelFilterResponseOutput {
	return o.ApplyT(func(v *EventChannelFilterResponse) EventChannelFilterResponse {
		if v != nil {
			return *v
		}
		var ret EventChannelFilterResponse
		return ret
	}).(EventChannelFilterResponseOutput)
}

func (o EventChannelFilterResponsePtrOutput) AdvancedFilters() pulumi.ArrayOutput {
	return o.ApplyT(func(v *EventChannelFilterResponse) []interface{} {
		if v == nil {
			return nil
		}
		return v.AdvancedFilters
	}).(pulumi.ArrayOutput)
}

func (o EventChannelFilterResponsePtrOutput) EnableAdvancedFilteringOnArrays() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventChannelFilterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAdvancedFilteringOnArrays
	}).(pulumi.BoolPtrOutput)
}

type EventChannelSource struct {
	Source *string `pulumi:"source"`
}





type EventChannelSourceInput interface {
	pulumi.Input

	ToEventChannelSourceOutput() EventChannelSourceOutput
	ToEventChannelSourceOutputWithContext(context.Context) EventChannelSourceOutput
}

type EventChannelSourceArgs struct {
	Source pulumi.StringPtrInput `pulumi:"source"`
}

func (EventChannelSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventChannelSource)(nil)).Elem()
}

func (i EventChannelSourceArgs) ToEventChannelSourceOutput() EventChannelSourceOutput {
	return i.ToEventChannelSourceOutputWithContext(context.Background())
}

func (i EventChannelSourceArgs) ToEventChannelSourceOutputWithContext(ctx context.Context) EventChannelSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventChannelSourceOutput)
}

func (i EventChannelSourceArgs) ToEventChannelSourcePtrOutput() EventChannelSourcePtrOutput {
	return i.ToEventChannelSourcePtrOutputWithContext(context.Background())
}

func (i EventChannelSourceArgs) ToEventChannelSourcePtrOutputWithContext(ctx context.Context) EventChannelSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventChannelSourceOutput).ToEventChannelSourcePtrOutputWithContext(ctx)
}









type EventChannelSourcePtrInput interface {
	pulumi.Input

	ToEventChannelSourcePtrOutput() EventChannelSourcePtrOutput
	ToEventChannelSourcePtrOutputWithContext(context.Context) EventChannelSourcePtrOutput
}

type eventChannelSourcePtrType EventChannelSourceArgs

func EventChannelSourcePtr(v *EventChannelSourceArgs) EventChannelSourcePtrInput {
	return (*eventChannelSourcePtrType)(v)
}

func (*eventChannelSourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventChannelSource)(nil)).Elem()
}

func (i *eventChannelSourcePtrType) ToEventChannelSourcePtrOutput() EventChannelSourcePtrOutput {
	return i.ToEventChannelSourcePtrOutputWithContext(context.Background())
}

func (i *eventChannelSourcePtrType) ToEventChannelSourcePtrOutputWithContext(ctx context.Context) EventChannelSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventChannelSourcePtrOutput)
}

type EventChannelSourceOutput struct{ *pulumi.OutputState }

func (EventChannelSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventChannelSource)(nil)).Elem()
}

func (o EventChannelSourceOutput) ToEventChannelSourceOutput() EventChannelSourceOutput {
	return o
}

func (o EventChannelSourceOutput) ToEventChannelSourceOutputWithContext(ctx context.Context) EventChannelSourceOutput {
	return o
}

func (o EventChannelSourceOutput) ToEventChannelSourcePtrOutput() EventChannelSourcePtrOutput {
	return o.ToEventChannelSourcePtrOutputWithContext(context.Background())
}

func (o EventChannelSourceOutput) ToEventChannelSourcePtrOutputWithContext(ctx context.Context) EventChannelSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventChannelSource) *EventChannelSource {
		return &v
	}).(EventChannelSourcePtrOutput)
}

func (o EventChannelSourceOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventChannelSource) *string { return v.Source }).(pulumi.StringPtrOutput)
}

type EventChannelSourcePtrOutput struct{ *pulumi.OutputState }

func (EventChannelSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventChannelSource)(nil)).Elem()
}

func (o EventChannelSourcePtrOutput) ToEventChannelSourcePtrOutput() EventChannelSourcePtrOutput {
	return o
}

func (o EventChannelSourcePtrOutput) ToEventChannelSourcePtrOutputWithContext(ctx context.Context) EventChannelSourcePtrOutput {
	return o
}

func (o EventChannelSourcePtrOutput) Elem() EventChannelSourceOutput {
	return o.ApplyT(func(v *EventChannelSource) EventChannelSource {
		if v != nil {
			return *v
		}
		var ret EventChannelSource
		return ret
	}).(EventChannelSourceOutput)
}

func (o EventChannelSourcePtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventChannelSource) *string {
		if v == nil {
			return nil
		}
		return v.Source
	}).(pulumi.StringPtrOutput)
}

type EventChannelSourceResponse struct {
	Source *string `pulumi:"source"`
}

type EventChannelSourceResponseOutput struct{ *pulumi.OutputState }

func (EventChannelSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventChannelSourceResponse)(nil)).Elem()
}

func (o EventChannelSourceResponseOutput) ToEventChannelSourceResponseOutput() EventChannelSourceResponseOutput {
	return o
}

func (o EventChannelSourceResponseOutput) ToEventChannelSourceResponseOutputWithContext(ctx context.Context) EventChannelSourceResponseOutput {
	return o
}

func (o EventChannelSourceResponseOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventChannelSourceResponse) *string { return v.Source }).(pulumi.StringPtrOutput)
}

type EventChannelSourceResponsePtrOutput struct{ *pulumi.OutputState }

func (EventChannelSourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventChannelSourceResponse)(nil)).Elem()
}

func (o EventChannelSourceResponsePtrOutput) ToEventChannelSourceResponsePtrOutput() EventChannelSourceResponsePtrOutput {
	return o
}

func (o EventChannelSourceResponsePtrOutput) ToEventChannelSourceResponsePtrOutputWithContext(ctx context.Context) EventChannelSourceResponsePtrOutput {
	return o
}

func (o EventChannelSourceResponsePtrOutput) Elem() EventChannelSourceResponseOutput {
	return o.ApplyT(func(v *EventChannelSourceResponse) EventChannelSourceResponse {
		if v != nil {
			return *v
		}
		var ret EventChannelSourceResponse
		return ret
	}).(EventChannelSourceResponseOutput)
}

func (o EventChannelSourceResponsePtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventChannelSourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Source
	}).(pulumi.StringPtrOutput)
}

type EventHubEventSubscriptionDestination struct {
	DeliveryAttributeMappings []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType              string        `pulumi:"endpointType"`
	ResourceId                *string       `pulumi:"resourceId"`
}

type EventHubEventSubscriptionDestinationResponse struct {
	DeliveryAttributeMappings []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType              string        `pulumi:"endpointType"`
	ResourceId                *string       `pulumi:"resourceId"`
}

type EventSubscriptionFilter struct {
	AdvancedFilters                 []interface{} `pulumi:"advancedFilters"`
	EnableAdvancedFilteringOnArrays *bool         `pulumi:"enableAdvancedFilteringOnArrays"`
	IncludedEventTypes              []string      `pulumi:"includedEventTypes"`
	IsSubjectCaseSensitive          *bool         `pulumi:"isSubjectCaseSensitive"`
	SubjectBeginsWith               *string       `pulumi:"subjectBeginsWith"`
	SubjectEndsWith                 *string       `pulumi:"subjectEndsWith"`
}


func (val *EventSubscriptionFilter) Defaults() *EventSubscriptionFilter {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSubjectCaseSensitive) {
		isSubjectCaseSensitive_ := false
		tmp.IsSubjectCaseSensitive = &isSubjectCaseSensitive_
	}
	return &tmp
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


func (val *EventSubscriptionFilterArgs) Defaults() *EventSubscriptionFilterArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSubjectCaseSensitive) {
		tmp.IsSubjectCaseSensitive = pulumi.BoolPtr(false)
	}
	return &tmp
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


func (val *EventSubscriptionFilterResponse) Defaults() *EventSubscriptionFilterResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSubjectCaseSensitive) {
		isSubjectCaseSensitive_ := false
		tmp.IsSubjectCaseSensitive = &isSubjectCaseSensitive_
	}
	return &tmp
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

type EventTypeInfo struct {
	InlineEventTypes map[string]InlineEventProperties `pulumi:"inlineEventTypes"`
	Kind             *string                          `pulumi:"kind"`
}





type EventTypeInfoInput interface {
	pulumi.Input

	ToEventTypeInfoOutput() EventTypeInfoOutput
	ToEventTypeInfoOutputWithContext(context.Context) EventTypeInfoOutput
}

type EventTypeInfoArgs struct {
	InlineEventTypes InlineEventPropertiesMapInput `pulumi:"inlineEventTypes"`
	Kind             pulumi.StringPtrInput         `pulumi:"kind"`
}

func (EventTypeInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventTypeInfo)(nil)).Elem()
}

func (i EventTypeInfoArgs) ToEventTypeInfoOutput() EventTypeInfoOutput {
	return i.ToEventTypeInfoOutputWithContext(context.Background())
}

func (i EventTypeInfoArgs) ToEventTypeInfoOutputWithContext(ctx context.Context) EventTypeInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTypeInfoOutput)
}

func (i EventTypeInfoArgs) ToEventTypeInfoPtrOutput() EventTypeInfoPtrOutput {
	return i.ToEventTypeInfoPtrOutputWithContext(context.Background())
}

func (i EventTypeInfoArgs) ToEventTypeInfoPtrOutputWithContext(ctx context.Context) EventTypeInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTypeInfoOutput).ToEventTypeInfoPtrOutputWithContext(ctx)
}









type EventTypeInfoPtrInput interface {
	pulumi.Input

	ToEventTypeInfoPtrOutput() EventTypeInfoPtrOutput
	ToEventTypeInfoPtrOutputWithContext(context.Context) EventTypeInfoPtrOutput
}

type eventTypeInfoPtrType EventTypeInfoArgs

func EventTypeInfoPtr(v *EventTypeInfoArgs) EventTypeInfoPtrInput {
	return (*eventTypeInfoPtrType)(v)
}

func (*eventTypeInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventTypeInfo)(nil)).Elem()
}

func (i *eventTypeInfoPtrType) ToEventTypeInfoPtrOutput() EventTypeInfoPtrOutput {
	return i.ToEventTypeInfoPtrOutputWithContext(context.Background())
}

func (i *eventTypeInfoPtrType) ToEventTypeInfoPtrOutputWithContext(ctx context.Context) EventTypeInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTypeInfoPtrOutput)
}

type EventTypeInfoOutput struct{ *pulumi.OutputState }

func (EventTypeInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventTypeInfo)(nil)).Elem()
}

func (o EventTypeInfoOutput) ToEventTypeInfoOutput() EventTypeInfoOutput {
	return o
}

func (o EventTypeInfoOutput) ToEventTypeInfoOutputWithContext(ctx context.Context) EventTypeInfoOutput {
	return o
}

func (o EventTypeInfoOutput) ToEventTypeInfoPtrOutput() EventTypeInfoPtrOutput {
	return o.ToEventTypeInfoPtrOutputWithContext(context.Background())
}

func (o EventTypeInfoOutput) ToEventTypeInfoPtrOutputWithContext(ctx context.Context) EventTypeInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventTypeInfo) *EventTypeInfo {
		return &v
	}).(EventTypeInfoPtrOutput)
}

func (o EventTypeInfoOutput) InlineEventTypes() InlineEventPropertiesMapOutput {
	return o.ApplyT(func(v EventTypeInfo) map[string]InlineEventProperties { return v.InlineEventTypes }).(InlineEventPropertiesMapOutput)
}

func (o EventTypeInfoOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventTypeInfo) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

type EventTypeInfoPtrOutput struct{ *pulumi.OutputState }

func (EventTypeInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventTypeInfo)(nil)).Elem()
}

func (o EventTypeInfoPtrOutput) ToEventTypeInfoPtrOutput() EventTypeInfoPtrOutput {
	return o
}

func (o EventTypeInfoPtrOutput) ToEventTypeInfoPtrOutputWithContext(ctx context.Context) EventTypeInfoPtrOutput {
	return o
}

func (o EventTypeInfoPtrOutput) Elem() EventTypeInfoOutput {
	return o.ApplyT(func(v *EventTypeInfo) EventTypeInfo {
		if v != nil {
			return *v
		}
		var ret EventTypeInfo
		return ret
	}).(EventTypeInfoOutput)
}

func (o EventTypeInfoPtrOutput) InlineEventTypes() InlineEventPropertiesMapOutput {
	return o.ApplyT(func(v *EventTypeInfo) map[string]InlineEventProperties {
		if v == nil {
			return nil
		}
		return v.InlineEventTypes
	}).(InlineEventPropertiesMapOutput)
}

func (o EventTypeInfoPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventTypeInfo) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

type EventTypeInfoResponse struct {
	InlineEventTypes map[string]InlineEventPropertiesResponse `pulumi:"inlineEventTypes"`
	Kind             *string                                  `pulumi:"kind"`
}

type EventTypeInfoResponseOutput struct{ *pulumi.OutputState }

func (EventTypeInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventTypeInfoResponse)(nil)).Elem()
}

func (o EventTypeInfoResponseOutput) ToEventTypeInfoResponseOutput() EventTypeInfoResponseOutput {
	return o
}

func (o EventTypeInfoResponseOutput) ToEventTypeInfoResponseOutputWithContext(ctx context.Context) EventTypeInfoResponseOutput {
	return o
}

func (o EventTypeInfoResponseOutput) InlineEventTypes() InlineEventPropertiesResponseMapOutput {
	return o.ApplyT(func(v EventTypeInfoResponse) map[string]InlineEventPropertiesResponse { return v.InlineEventTypes }).(InlineEventPropertiesResponseMapOutput)
}

func (o EventTypeInfoResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventTypeInfoResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

type EventTypeInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (EventTypeInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventTypeInfoResponse)(nil)).Elem()
}

func (o EventTypeInfoResponsePtrOutput) ToEventTypeInfoResponsePtrOutput() EventTypeInfoResponsePtrOutput {
	return o
}

func (o EventTypeInfoResponsePtrOutput) ToEventTypeInfoResponsePtrOutputWithContext(ctx context.Context) EventTypeInfoResponsePtrOutput {
	return o
}

func (o EventTypeInfoResponsePtrOutput) Elem() EventTypeInfoResponseOutput {
	return o.ApplyT(func(v *EventTypeInfoResponse) EventTypeInfoResponse {
		if v != nil {
			return *v
		}
		var ret EventTypeInfoResponse
		return ret
	}).(EventTypeInfoResponseOutput)
}

func (o EventTypeInfoResponsePtrOutput) InlineEventTypes() InlineEventPropertiesResponseMapOutput {
	return o.ApplyT(func(v *EventTypeInfoResponse) map[string]InlineEventPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.InlineEventTypes
	}).(InlineEventPropertiesResponseMapOutput)
}

func (o EventTypeInfoResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventTypeInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ExtendedLocationInput interface {
	pulumi.Input

	ToExtendedLocationOutput() ExtendedLocationOutput
	ToExtendedLocationOutputWithContext(context.Context) ExtendedLocationOutput
}

type ExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (i ExtendedLocationArgs) ToExtendedLocationOutput() ExtendedLocationOutput {
	return i.ToExtendedLocationOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput)
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput).ToExtendedLocationPtrOutputWithContext(ctx)
}









type ExtendedLocationPtrInput interface {
	pulumi.Input

	ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput
	ToExtendedLocationPtrOutputWithContext(context.Context) ExtendedLocationPtrOutput
}

type extendedLocationPtrType ExtendedLocationArgs

func ExtendedLocationPtr(v *ExtendedLocationArgs) ExtendedLocationPtrInput {
	return (*extendedLocationPtrType)(v)
}

func (*extendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationPtrOutput)
}

type ExtendedLocationOutput struct{ *pulumi.OutputState }

func (ExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationOutput) ToExtendedLocationOutput() ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocation) *ExtendedLocation {
		return &v
	}).(ExtendedLocationPtrOutput)
}

func (o ExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) Elem() ExtendedLocationOutput {
	return o.ApplyT(func(v *ExtendedLocation) ExtendedLocation {
		if v != nil {
			return *v
		}
		var ret ExtendedLocation
		return ret
	}).(ExtendedLocationOutput)
}

func (o ExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type ExtendedLocationResponseOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutputWithContext(ctx context.Context) ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponsePtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) Elem() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) ExtendedLocationResponse {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationResponse
		return ret
	}).(ExtendedLocationResponseOutput)
}

func (o ExtendedLocationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type HybridConnectionEventSubscriptionDestination struct {
	DeliveryAttributeMappings []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType              string        `pulumi:"endpointType"`
	ResourceId                *string       `pulumi:"resourceId"`
}

type HybridConnectionEventSubscriptionDestinationResponse struct {
	DeliveryAttributeMappings []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType              string        `pulumi:"endpointType"`
	ResourceId                *string       `pulumi:"resourceId"`
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

type InlineEventProperties struct {
	DataSchemaUrl    *string `pulumi:"dataSchemaUrl"`
	Description      *string `pulumi:"description"`
	DocumentationUrl *string `pulumi:"documentationUrl"`
}





type InlineEventPropertiesInput interface {
	pulumi.Input

	ToInlineEventPropertiesOutput() InlineEventPropertiesOutput
	ToInlineEventPropertiesOutputWithContext(context.Context) InlineEventPropertiesOutput
}

type InlineEventPropertiesArgs struct {
	DataSchemaUrl    pulumi.StringPtrInput `pulumi:"dataSchemaUrl"`
	Description      pulumi.StringPtrInput `pulumi:"description"`
	DocumentationUrl pulumi.StringPtrInput `pulumi:"documentationUrl"`
}

func (InlineEventPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InlineEventProperties)(nil)).Elem()
}

func (i InlineEventPropertiesArgs) ToInlineEventPropertiesOutput() InlineEventPropertiesOutput {
	return i.ToInlineEventPropertiesOutputWithContext(context.Background())
}

func (i InlineEventPropertiesArgs) ToInlineEventPropertiesOutputWithContext(ctx context.Context) InlineEventPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InlineEventPropertiesOutput)
}





type InlineEventPropertiesMapInput interface {
	pulumi.Input

	ToInlineEventPropertiesMapOutput() InlineEventPropertiesMapOutput
	ToInlineEventPropertiesMapOutputWithContext(context.Context) InlineEventPropertiesMapOutput
}

type InlineEventPropertiesMap map[string]InlineEventPropertiesInput

func (InlineEventPropertiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InlineEventProperties)(nil)).Elem()
}

func (i InlineEventPropertiesMap) ToInlineEventPropertiesMapOutput() InlineEventPropertiesMapOutput {
	return i.ToInlineEventPropertiesMapOutputWithContext(context.Background())
}

func (i InlineEventPropertiesMap) ToInlineEventPropertiesMapOutputWithContext(ctx context.Context) InlineEventPropertiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InlineEventPropertiesMapOutput)
}

type InlineEventPropertiesOutput struct{ *pulumi.OutputState }

func (InlineEventPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InlineEventProperties)(nil)).Elem()
}

func (o InlineEventPropertiesOutput) ToInlineEventPropertiesOutput() InlineEventPropertiesOutput {
	return o
}

func (o InlineEventPropertiesOutput) ToInlineEventPropertiesOutputWithContext(ctx context.Context) InlineEventPropertiesOutput {
	return o
}

func (o InlineEventPropertiesOutput) DataSchemaUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InlineEventProperties) *string { return v.DataSchemaUrl }).(pulumi.StringPtrOutput)
}

func (o InlineEventPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InlineEventProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o InlineEventPropertiesOutput) DocumentationUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InlineEventProperties) *string { return v.DocumentationUrl }).(pulumi.StringPtrOutput)
}

type InlineEventPropertiesMapOutput struct{ *pulumi.OutputState }

func (InlineEventPropertiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InlineEventProperties)(nil)).Elem()
}

func (o InlineEventPropertiesMapOutput) ToInlineEventPropertiesMapOutput() InlineEventPropertiesMapOutput {
	return o
}

func (o InlineEventPropertiesMapOutput) ToInlineEventPropertiesMapOutputWithContext(ctx context.Context) InlineEventPropertiesMapOutput {
	return o
}

func (o InlineEventPropertiesMapOutput) MapIndex(k pulumi.StringInput) InlineEventPropertiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InlineEventProperties {
		return vs[0].(map[string]InlineEventProperties)[vs[1].(string)]
	}).(InlineEventPropertiesOutput)
}

type InlineEventPropertiesResponse struct {
	DataSchemaUrl    *string `pulumi:"dataSchemaUrl"`
	Description      *string `pulumi:"description"`
	DocumentationUrl *string `pulumi:"documentationUrl"`
}

type InlineEventPropertiesResponseOutput struct{ *pulumi.OutputState }

func (InlineEventPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InlineEventPropertiesResponse)(nil)).Elem()
}

func (o InlineEventPropertiesResponseOutput) ToInlineEventPropertiesResponseOutput() InlineEventPropertiesResponseOutput {
	return o
}

func (o InlineEventPropertiesResponseOutput) ToInlineEventPropertiesResponseOutputWithContext(ctx context.Context) InlineEventPropertiesResponseOutput {
	return o
}

func (o InlineEventPropertiesResponseOutput) DataSchemaUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InlineEventPropertiesResponse) *string { return v.DataSchemaUrl }).(pulumi.StringPtrOutput)
}

func (o InlineEventPropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InlineEventPropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o InlineEventPropertiesResponseOutput) DocumentationUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InlineEventPropertiesResponse) *string { return v.DocumentationUrl }).(pulumi.StringPtrOutput)
}

type InlineEventPropertiesResponseMapOutput struct{ *pulumi.OutputState }

func (InlineEventPropertiesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InlineEventPropertiesResponse)(nil)).Elem()
}

func (o InlineEventPropertiesResponseMapOutput) ToInlineEventPropertiesResponseMapOutput() InlineEventPropertiesResponseMapOutput {
	return o
}

func (o InlineEventPropertiesResponseMapOutput) ToInlineEventPropertiesResponseMapOutputWithContext(ctx context.Context) InlineEventPropertiesResponseMapOutput {
	return o
}

func (o InlineEventPropertiesResponseMapOutput) MapIndex(k pulumi.StringInput) InlineEventPropertiesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InlineEventPropertiesResponse {
		return vs[0].(map[string]InlineEventPropertiesResponse)[vs[1].(string)]
	}).(InlineEventPropertiesResponseOutput)
}

type IsNotNullAdvancedFilter struct {
	Key          *string `pulumi:"key"`
	OperatorType string  `pulumi:"operatorType"`
}

type IsNotNullAdvancedFilterResponse struct {
	Key          *string `pulumi:"key"`
	OperatorType string  `pulumi:"operatorType"`
}

type IsNullOrUndefinedAdvancedFilter struct {
	Key          *string `pulumi:"key"`
	OperatorType string  `pulumi:"operatorType"`
}

type IsNullOrUndefinedAdvancedFilterResponse struct {
	Key          *string `pulumi:"key"`
	OperatorType string  `pulumi:"operatorType"`
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

type NumberGreaterThanAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}

type NumberGreaterThanOrEqualsAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}

type NumberGreaterThanOrEqualsAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}

type NumberInAdvancedFilter struct {
	Key          *string   `pulumi:"key"`
	OperatorType string    `pulumi:"operatorType"`
	Values       []float64 `pulumi:"values"`
}

type NumberInAdvancedFilterResponse struct {
	Key          *string   `pulumi:"key"`
	OperatorType string    `pulumi:"operatorType"`
	Values       []float64 `pulumi:"values"`
}

type NumberInRangeAdvancedFilter struct {
	Key          *string     `pulumi:"key"`
	OperatorType string      `pulumi:"operatorType"`
	Values       [][]float64 `pulumi:"values"`
}

type NumberInRangeAdvancedFilterResponse struct {
	Key          *string     `pulumi:"key"`
	OperatorType string      `pulumi:"operatorType"`
	Values       [][]float64 `pulumi:"values"`
}

type NumberLessThanAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}

type NumberLessThanAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}

type NumberLessThanOrEqualsAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}

type NumberLessThanOrEqualsAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}

type NumberNotInAdvancedFilter struct {
	Key          *string   `pulumi:"key"`
	OperatorType string    `pulumi:"operatorType"`
	Values       []float64 `pulumi:"values"`
}

type NumberNotInAdvancedFilterResponse struct {
	Key          *string   `pulumi:"key"`
	OperatorType string    `pulumi:"operatorType"`
	Values       []float64 `pulumi:"values"`
}

type NumberNotInRangeAdvancedFilter struct {
	Key          *string     `pulumi:"key"`
	OperatorType string      `pulumi:"operatorType"`
	Values       [][]float64 `pulumi:"values"`
}

type NumberNotInRangeAdvancedFilterResponse struct {
	Key          *string     `pulumi:"key"`
	OperatorType string      `pulumi:"operatorType"`
	Values       [][]float64 `pulumi:"values"`
}

type Partner struct {
	AuthorizationExpirationTimeInUtc *string `pulumi:"authorizationExpirationTimeInUtc"`
	PartnerName                      *string `pulumi:"partnerName"`
	PartnerRegistrationImmutableId   *string `pulumi:"partnerRegistrationImmutableId"`
}





type PartnerInput interface {
	pulumi.Input

	ToPartnerOutput() PartnerOutput
	ToPartnerOutputWithContext(context.Context) PartnerOutput
}

type PartnerArgs struct {
	AuthorizationExpirationTimeInUtc pulumi.StringPtrInput `pulumi:"authorizationExpirationTimeInUtc"`
	PartnerName                      pulumi.StringPtrInput `pulumi:"partnerName"`
	PartnerRegistrationImmutableId   pulumi.StringPtrInput `pulumi:"partnerRegistrationImmutableId"`
}

func (PartnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Partner)(nil)).Elem()
}

func (i PartnerArgs) ToPartnerOutput() PartnerOutput {
	return i.ToPartnerOutputWithContext(context.Background())
}

func (i PartnerArgs) ToPartnerOutputWithContext(ctx context.Context) PartnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerOutput)
}





type PartnerArrayInput interface {
	pulumi.Input

	ToPartnerArrayOutput() PartnerArrayOutput
	ToPartnerArrayOutputWithContext(context.Context) PartnerArrayOutput
}

type PartnerArray []PartnerInput

func (PartnerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Partner)(nil)).Elem()
}

func (i PartnerArray) ToPartnerArrayOutput() PartnerArrayOutput {
	return i.ToPartnerArrayOutputWithContext(context.Background())
}

func (i PartnerArray) ToPartnerArrayOutputWithContext(ctx context.Context) PartnerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerArrayOutput)
}

type PartnerOutput struct{ *pulumi.OutputState }

func (PartnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Partner)(nil)).Elem()
}

func (o PartnerOutput) ToPartnerOutput() PartnerOutput {
	return o
}

func (o PartnerOutput) ToPartnerOutputWithContext(ctx context.Context) PartnerOutput {
	return o
}

func (o PartnerOutput) AuthorizationExpirationTimeInUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Partner) *string { return v.AuthorizationExpirationTimeInUtc }).(pulumi.StringPtrOutput)
}

func (o PartnerOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Partner) *string { return v.PartnerName }).(pulumi.StringPtrOutput)
}

func (o PartnerOutput) PartnerRegistrationImmutableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Partner) *string { return v.PartnerRegistrationImmutableId }).(pulumi.StringPtrOutput)
}

type PartnerArrayOutput struct{ *pulumi.OutputState }

func (PartnerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Partner)(nil)).Elem()
}

func (o PartnerArrayOutput) ToPartnerArrayOutput() PartnerArrayOutput {
	return o
}

func (o PartnerArrayOutput) ToPartnerArrayOutputWithContext(ctx context.Context) PartnerArrayOutput {
	return o
}

func (o PartnerArrayOutput) Index(i pulumi.IntInput) PartnerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Partner {
		return vs[0].([]Partner)[vs[1].(int)]
	}).(PartnerOutput)
}

type PartnerAuthorization struct {
	AuthorizedPartnersList             []Partner `pulumi:"authorizedPartnersList"`
	DefaultMaximumExpirationTimeInDays *int      `pulumi:"defaultMaximumExpirationTimeInDays"`
}





type PartnerAuthorizationInput interface {
	pulumi.Input

	ToPartnerAuthorizationOutput() PartnerAuthorizationOutput
	ToPartnerAuthorizationOutputWithContext(context.Context) PartnerAuthorizationOutput
}

type PartnerAuthorizationArgs struct {
	AuthorizedPartnersList             PartnerArrayInput  `pulumi:"authorizedPartnersList"`
	DefaultMaximumExpirationTimeInDays pulumi.IntPtrInput `pulumi:"defaultMaximumExpirationTimeInDays"`
}

func (PartnerAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerAuthorization)(nil)).Elem()
}

func (i PartnerAuthorizationArgs) ToPartnerAuthorizationOutput() PartnerAuthorizationOutput {
	return i.ToPartnerAuthorizationOutputWithContext(context.Background())
}

func (i PartnerAuthorizationArgs) ToPartnerAuthorizationOutputWithContext(ctx context.Context) PartnerAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerAuthorizationOutput)
}

func (i PartnerAuthorizationArgs) ToPartnerAuthorizationPtrOutput() PartnerAuthorizationPtrOutput {
	return i.ToPartnerAuthorizationPtrOutputWithContext(context.Background())
}

func (i PartnerAuthorizationArgs) ToPartnerAuthorizationPtrOutputWithContext(ctx context.Context) PartnerAuthorizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerAuthorizationOutput).ToPartnerAuthorizationPtrOutputWithContext(ctx)
}









type PartnerAuthorizationPtrInput interface {
	pulumi.Input

	ToPartnerAuthorizationPtrOutput() PartnerAuthorizationPtrOutput
	ToPartnerAuthorizationPtrOutputWithContext(context.Context) PartnerAuthorizationPtrOutput
}

type partnerAuthorizationPtrType PartnerAuthorizationArgs

func PartnerAuthorizationPtr(v *PartnerAuthorizationArgs) PartnerAuthorizationPtrInput {
	return (*partnerAuthorizationPtrType)(v)
}

func (*partnerAuthorizationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerAuthorization)(nil)).Elem()
}

func (i *partnerAuthorizationPtrType) ToPartnerAuthorizationPtrOutput() PartnerAuthorizationPtrOutput {
	return i.ToPartnerAuthorizationPtrOutputWithContext(context.Background())
}

func (i *partnerAuthorizationPtrType) ToPartnerAuthorizationPtrOutputWithContext(ctx context.Context) PartnerAuthorizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerAuthorizationPtrOutput)
}

type PartnerAuthorizationOutput struct{ *pulumi.OutputState }

func (PartnerAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerAuthorization)(nil)).Elem()
}

func (o PartnerAuthorizationOutput) ToPartnerAuthorizationOutput() PartnerAuthorizationOutput {
	return o
}

func (o PartnerAuthorizationOutput) ToPartnerAuthorizationOutputWithContext(ctx context.Context) PartnerAuthorizationOutput {
	return o
}

func (o PartnerAuthorizationOutput) ToPartnerAuthorizationPtrOutput() PartnerAuthorizationPtrOutput {
	return o.ToPartnerAuthorizationPtrOutputWithContext(context.Background())
}

func (o PartnerAuthorizationOutput) ToPartnerAuthorizationPtrOutputWithContext(ctx context.Context) PartnerAuthorizationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PartnerAuthorization) *PartnerAuthorization {
		return &v
	}).(PartnerAuthorizationPtrOutput)
}

func (o PartnerAuthorizationOutput) AuthorizedPartnersList() PartnerArrayOutput {
	return o.ApplyT(func(v PartnerAuthorization) []Partner { return v.AuthorizedPartnersList }).(PartnerArrayOutput)
}

func (o PartnerAuthorizationOutput) DefaultMaximumExpirationTimeInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PartnerAuthorization) *int { return v.DefaultMaximumExpirationTimeInDays }).(pulumi.IntPtrOutput)
}

type PartnerAuthorizationPtrOutput struct{ *pulumi.OutputState }

func (PartnerAuthorizationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerAuthorization)(nil)).Elem()
}

func (o PartnerAuthorizationPtrOutput) ToPartnerAuthorizationPtrOutput() PartnerAuthorizationPtrOutput {
	return o
}

func (o PartnerAuthorizationPtrOutput) ToPartnerAuthorizationPtrOutputWithContext(ctx context.Context) PartnerAuthorizationPtrOutput {
	return o
}

func (o PartnerAuthorizationPtrOutput) Elem() PartnerAuthorizationOutput {
	return o.ApplyT(func(v *PartnerAuthorization) PartnerAuthorization {
		if v != nil {
			return *v
		}
		var ret PartnerAuthorization
		return ret
	}).(PartnerAuthorizationOutput)
}

func (o PartnerAuthorizationPtrOutput) AuthorizedPartnersList() PartnerArrayOutput {
	return o.ApplyT(func(v *PartnerAuthorization) []Partner {
		if v == nil {
			return nil
		}
		return v.AuthorizedPartnersList
	}).(PartnerArrayOutput)
}

func (o PartnerAuthorizationPtrOutput) DefaultMaximumExpirationTimeInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PartnerAuthorization) *int {
		if v == nil {
			return nil
		}
		return v.DefaultMaximumExpirationTimeInDays
	}).(pulumi.IntPtrOutput)
}

type PartnerAuthorizationResponse struct {
	AuthorizedPartnersList             []PartnerResponse `pulumi:"authorizedPartnersList"`
	DefaultMaximumExpirationTimeInDays *int              `pulumi:"defaultMaximumExpirationTimeInDays"`
}

type PartnerAuthorizationResponseOutput struct{ *pulumi.OutputState }

func (PartnerAuthorizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerAuthorizationResponse)(nil)).Elem()
}

func (o PartnerAuthorizationResponseOutput) ToPartnerAuthorizationResponseOutput() PartnerAuthorizationResponseOutput {
	return o
}

func (o PartnerAuthorizationResponseOutput) ToPartnerAuthorizationResponseOutputWithContext(ctx context.Context) PartnerAuthorizationResponseOutput {
	return o
}

func (o PartnerAuthorizationResponseOutput) AuthorizedPartnersList() PartnerResponseArrayOutput {
	return o.ApplyT(func(v PartnerAuthorizationResponse) []PartnerResponse { return v.AuthorizedPartnersList }).(PartnerResponseArrayOutput)
}

func (o PartnerAuthorizationResponseOutput) DefaultMaximumExpirationTimeInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PartnerAuthorizationResponse) *int { return v.DefaultMaximumExpirationTimeInDays }).(pulumi.IntPtrOutput)
}

type PartnerAuthorizationResponsePtrOutput struct{ *pulumi.OutputState }

func (PartnerAuthorizationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerAuthorizationResponse)(nil)).Elem()
}

func (o PartnerAuthorizationResponsePtrOutput) ToPartnerAuthorizationResponsePtrOutput() PartnerAuthorizationResponsePtrOutput {
	return o
}

func (o PartnerAuthorizationResponsePtrOutput) ToPartnerAuthorizationResponsePtrOutputWithContext(ctx context.Context) PartnerAuthorizationResponsePtrOutput {
	return o
}

func (o PartnerAuthorizationResponsePtrOutput) Elem() PartnerAuthorizationResponseOutput {
	return o.ApplyT(func(v *PartnerAuthorizationResponse) PartnerAuthorizationResponse {
		if v != nil {
			return *v
		}
		var ret PartnerAuthorizationResponse
		return ret
	}).(PartnerAuthorizationResponseOutput)
}

func (o PartnerAuthorizationResponsePtrOutput) AuthorizedPartnersList() PartnerResponseArrayOutput {
	return o.ApplyT(func(v *PartnerAuthorizationResponse) []PartnerResponse {
		if v == nil {
			return nil
		}
		return v.AuthorizedPartnersList
	}).(PartnerResponseArrayOutput)
}

func (o PartnerAuthorizationResponsePtrOutput) DefaultMaximumExpirationTimeInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PartnerAuthorizationResponse) *int {
		if v == nil {
			return nil
		}
		return v.DefaultMaximumExpirationTimeInDays
	}).(pulumi.IntPtrOutput)
}

type PartnerEventSubscriptionDestination struct {
	EndpointType string  `pulumi:"endpointType"`
	ResourceId   *string `pulumi:"resourceId"`
}

type PartnerEventSubscriptionDestinationResponse struct {
	EndpointType string  `pulumi:"endpointType"`
	ResourceId   *string `pulumi:"resourceId"`
}

type PartnerResponse struct {
	AuthorizationExpirationTimeInUtc *string `pulumi:"authorizationExpirationTimeInUtc"`
	PartnerName                      *string `pulumi:"partnerName"`
	PartnerRegistrationImmutableId   *string `pulumi:"partnerRegistrationImmutableId"`
}

type PartnerResponseOutput struct{ *pulumi.OutputState }

func (PartnerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerResponse)(nil)).Elem()
}

func (o PartnerResponseOutput) ToPartnerResponseOutput() PartnerResponseOutput {
	return o
}

func (o PartnerResponseOutput) ToPartnerResponseOutputWithContext(ctx context.Context) PartnerResponseOutput {
	return o
}

func (o PartnerResponseOutput) AuthorizationExpirationTimeInUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartnerResponse) *string { return v.AuthorizationExpirationTimeInUtc }).(pulumi.StringPtrOutput)
}

func (o PartnerResponseOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartnerResponse) *string { return v.PartnerName }).(pulumi.StringPtrOutput)
}

func (o PartnerResponseOutput) PartnerRegistrationImmutableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartnerResponse) *string { return v.PartnerRegistrationImmutableId }).(pulumi.StringPtrOutput)
}

type PartnerResponseArrayOutput struct{ *pulumi.OutputState }

func (PartnerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartnerResponse)(nil)).Elem()
}

func (o PartnerResponseArrayOutput) ToPartnerResponseArrayOutput() PartnerResponseArrayOutput {
	return o
}

func (o PartnerResponseArrayOutput) ToPartnerResponseArrayOutputWithContext(ctx context.Context) PartnerResponseArrayOutput {
	return o
}

func (o PartnerResponseArrayOutput) Index(i pulumi.IntInput) PartnerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PartnerResponse {
		return vs[0].([]PartnerResponse)[vs[1].(int)]
	}).(PartnerResponseOutput)
}

type PartnerTopicInfo struct {
	AzureSubscriptionId *string        `pulumi:"azureSubscriptionId"`
	EventTypeInfo       *EventTypeInfo `pulumi:"eventTypeInfo"`
	Name                *string        `pulumi:"name"`
	ResourceGroupName   *string        `pulumi:"resourceGroupName"`
	Source              *string        `pulumi:"source"`
}





type PartnerTopicInfoInput interface {
	pulumi.Input

	ToPartnerTopicInfoOutput() PartnerTopicInfoOutput
	ToPartnerTopicInfoOutputWithContext(context.Context) PartnerTopicInfoOutput
}

type PartnerTopicInfoArgs struct {
	AzureSubscriptionId pulumi.StringPtrInput `pulumi:"azureSubscriptionId"`
	EventTypeInfo       EventTypeInfoPtrInput `pulumi:"eventTypeInfo"`
	Name                pulumi.StringPtrInput `pulumi:"name"`
	ResourceGroupName   pulumi.StringPtrInput `pulumi:"resourceGroupName"`
	Source              pulumi.StringPtrInput `pulumi:"source"`
}

func (PartnerTopicInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerTopicInfo)(nil)).Elem()
}

func (i PartnerTopicInfoArgs) ToPartnerTopicInfoOutput() PartnerTopicInfoOutput {
	return i.ToPartnerTopicInfoOutputWithContext(context.Background())
}

func (i PartnerTopicInfoArgs) ToPartnerTopicInfoOutputWithContext(ctx context.Context) PartnerTopicInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerTopicInfoOutput)
}

func (i PartnerTopicInfoArgs) ToPartnerTopicInfoPtrOutput() PartnerTopicInfoPtrOutput {
	return i.ToPartnerTopicInfoPtrOutputWithContext(context.Background())
}

func (i PartnerTopicInfoArgs) ToPartnerTopicInfoPtrOutputWithContext(ctx context.Context) PartnerTopicInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerTopicInfoOutput).ToPartnerTopicInfoPtrOutputWithContext(ctx)
}









type PartnerTopicInfoPtrInput interface {
	pulumi.Input

	ToPartnerTopicInfoPtrOutput() PartnerTopicInfoPtrOutput
	ToPartnerTopicInfoPtrOutputWithContext(context.Context) PartnerTopicInfoPtrOutput
}

type partnerTopicInfoPtrType PartnerTopicInfoArgs

func PartnerTopicInfoPtr(v *PartnerTopicInfoArgs) PartnerTopicInfoPtrInput {
	return (*partnerTopicInfoPtrType)(v)
}

func (*partnerTopicInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerTopicInfo)(nil)).Elem()
}

func (i *partnerTopicInfoPtrType) ToPartnerTopicInfoPtrOutput() PartnerTopicInfoPtrOutput {
	return i.ToPartnerTopicInfoPtrOutputWithContext(context.Background())
}

func (i *partnerTopicInfoPtrType) ToPartnerTopicInfoPtrOutputWithContext(ctx context.Context) PartnerTopicInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerTopicInfoPtrOutput)
}

type PartnerTopicInfoOutput struct{ *pulumi.OutputState }

func (PartnerTopicInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerTopicInfo)(nil)).Elem()
}

func (o PartnerTopicInfoOutput) ToPartnerTopicInfoOutput() PartnerTopicInfoOutput {
	return o
}

func (o PartnerTopicInfoOutput) ToPartnerTopicInfoOutputWithContext(ctx context.Context) PartnerTopicInfoOutput {
	return o
}

func (o PartnerTopicInfoOutput) ToPartnerTopicInfoPtrOutput() PartnerTopicInfoPtrOutput {
	return o.ToPartnerTopicInfoPtrOutputWithContext(context.Background())
}

func (o PartnerTopicInfoOutput) ToPartnerTopicInfoPtrOutputWithContext(ctx context.Context) PartnerTopicInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PartnerTopicInfo) *PartnerTopicInfo {
		return &v
	}).(PartnerTopicInfoPtrOutput)
}

func (o PartnerTopicInfoOutput) AzureSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartnerTopicInfo) *string { return v.AzureSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o PartnerTopicInfoOutput) EventTypeInfo() EventTypeInfoPtrOutput {
	return o.ApplyT(func(v PartnerTopicInfo) *EventTypeInfo { return v.EventTypeInfo }).(EventTypeInfoPtrOutput)
}

func (o PartnerTopicInfoOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartnerTopicInfo) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PartnerTopicInfoOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartnerTopicInfo) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o PartnerTopicInfoOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartnerTopicInfo) *string { return v.Source }).(pulumi.StringPtrOutput)
}

type PartnerTopicInfoPtrOutput struct{ *pulumi.OutputState }

func (PartnerTopicInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerTopicInfo)(nil)).Elem()
}

func (o PartnerTopicInfoPtrOutput) ToPartnerTopicInfoPtrOutput() PartnerTopicInfoPtrOutput {
	return o
}

func (o PartnerTopicInfoPtrOutput) ToPartnerTopicInfoPtrOutputWithContext(ctx context.Context) PartnerTopicInfoPtrOutput {
	return o
}

func (o PartnerTopicInfoPtrOutput) Elem() PartnerTopicInfoOutput {
	return o.ApplyT(func(v *PartnerTopicInfo) PartnerTopicInfo {
		if v != nil {
			return *v
		}
		var ret PartnerTopicInfo
		return ret
	}).(PartnerTopicInfoOutput)
}

func (o PartnerTopicInfoPtrOutput) AzureSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopicInfo) *string {
		if v == nil {
			return nil
		}
		return v.AzureSubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o PartnerTopicInfoPtrOutput) EventTypeInfo() EventTypeInfoPtrOutput {
	return o.ApplyT(func(v *PartnerTopicInfo) *EventTypeInfo {
		if v == nil {
			return nil
		}
		return v.EventTypeInfo
	}).(EventTypeInfoPtrOutput)
}

func (o PartnerTopicInfoPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopicInfo) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PartnerTopicInfoPtrOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopicInfo) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroupName
	}).(pulumi.StringPtrOutput)
}

func (o PartnerTopicInfoPtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopicInfo) *string {
		if v == nil {
			return nil
		}
		return v.Source
	}).(pulumi.StringPtrOutput)
}

type PartnerTopicInfoResponse struct {
	AzureSubscriptionId *string                `pulumi:"azureSubscriptionId"`
	EventTypeInfo       *EventTypeInfoResponse `pulumi:"eventTypeInfo"`
	Name                *string                `pulumi:"name"`
	ResourceGroupName   *string                `pulumi:"resourceGroupName"`
	Source              *string                `pulumi:"source"`
}

type PartnerTopicInfoResponseOutput struct{ *pulumi.OutputState }

func (PartnerTopicInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerTopicInfoResponse)(nil)).Elem()
}

func (o PartnerTopicInfoResponseOutput) ToPartnerTopicInfoResponseOutput() PartnerTopicInfoResponseOutput {
	return o
}

func (o PartnerTopicInfoResponseOutput) ToPartnerTopicInfoResponseOutputWithContext(ctx context.Context) PartnerTopicInfoResponseOutput {
	return o
}

func (o PartnerTopicInfoResponseOutput) AzureSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartnerTopicInfoResponse) *string { return v.AzureSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o PartnerTopicInfoResponseOutput) EventTypeInfo() EventTypeInfoResponsePtrOutput {
	return o.ApplyT(func(v PartnerTopicInfoResponse) *EventTypeInfoResponse { return v.EventTypeInfo }).(EventTypeInfoResponsePtrOutput)
}

func (o PartnerTopicInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartnerTopicInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PartnerTopicInfoResponseOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartnerTopicInfoResponse) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o PartnerTopicInfoResponseOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartnerTopicInfoResponse) *string { return v.Source }).(pulumi.StringPtrOutput)
}

type PartnerTopicInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (PartnerTopicInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerTopicInfoResponse)(nil)).Elem()
}

func (o PartnerTopicInfoResponsePtrOutput) ToPartnerTopicInfoResponsePtrOutput() PartnerTopicInfoResponsePtrOutput {
	return o
}

func (o PartnerTopicInfoResponsePtrOutput) ToPartnerTopicInfoResponsePtrOutputWithContext(ctx context.Context) PartnerTopicInfoResponsePtrOutput {
	return o
}

func (o PartnerTopicInfoResponsePtrOutput) Elem() PartnerTopicInfoResponseOutput {
	return o.ApplyT(func(v *PartnerTopicInfoResponse) PartnerTopicInfoResponse {
		if v != nil {
			return *v
		}
		var ret PartnerTopicInfoResponse
		return ret
	}).(PartnerTopicInfoResponseOutput)
}

func (o PartnerTopicInfoResponsePtrOutput) AzureSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopicInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.AzureSubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o PartnerTopicInfoResponsePtrOutput) EventTypeInfo() EventTypeInfoResponsePtrOutput {
	return o.ApplyT(func(v *PartnerTopicInfoResponse) *EventTypeInfoResponse {
		if v == nil {
			return nil
		}
		return v.EventTypeInfo
	}).(EventTypeInfoResponsePtrOutput)
}

func (o PartnerTopicInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopicInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PartnerTopicInfoResponsePtrOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopicInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroupName
	}).(pulumi.StringPtrOutput)
}

func (o PartnerTopicInfoResponsePtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopicInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Source
	}).(pulumi.StringPtrOutput)
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

type ResourceMoveChangeHistory struct {
	AzureSubscriptionId *string `pulumi:"azureSubscriptionId"`
	ChangedTimeUtc      *string `pulumi:"changedTimeUtc"`
	ResourceGroupName   *string `pulumi:"resourceGroupName"`
}





type ResourceMoveChangeHistoryInput interface {
	pulumi.Input

	ToResourceMoveChangeHistoryOutput() ResourceMoveChangeHistoryOutput
	ToResourceMoveChangeHistoryOutputWithContext(context.Context) ResourceMoveChangeHistoryOutput
}

type ResourceMoveChangeHistoryArgs struct {
	AzureSubscriptionId pulumi.StringPtrInput `pulumi:"azureSubscriptionId"`
	ChangedTimeUtc      pulumi.StringPtrInput `pulumi:"changedTimeUtc"`
	ResourceGroupName   pulumi.StringPtrInput `pulumi:"resourceGroupName"`
}

func (ResourceMoveChangeHistoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceMoveChangeHistory)(nil)).Elem()
}

func (i ResourceMoveChangeHistoryArgs) ToResourceMoveChangeHistoryOutput() ResourceMoveChangeHistoryOutput {
	return i.ToResourceMoveChangeHistoryOutputWithContext(context.Background())
}

func (i ResourceMoveChangeHistoryArgs) ToResourceMoveChangeHistoryOutputWithContext(ctx context.Context) ResourceMoveChangeHistoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMoveChangeHistoryOutput)
}





type ResourceMoveChangeHistoryArrayInput interface {
	pulumi.Input

	ToResourceMoveChangeHistoryArrayOutput() ResourceMoveChangeHistoryArrayOutput
	ToResourceMoveChangeHistoryArrayOutputWithContext(context.Context) ResourceMoveChangeHistoryArrayOutput
}

type ResourceMoveChangeHistoryArray []ResourceMoveChangeHistoryInput

func (ResourceMoveChangeHistoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceMoveChangeHistory)(nil)).Elem()
}

func (i ResourceMoveChangeHistoryArray) ToResourceMoveChangeHistoryArrayOutput() ResourceMoveChangeHistoryArrayOutput {
	return i.ToResourceMoveChangeHistoryArrayOutputWithContext(context.Background())
}

func (i ResourceMoveChangeHistoryArray) ToResourceMoveChangeHistoryArrayOutputWithContext(ctx context.Context) ResourceMoveChangeHistoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMoveChangeHistoryArrayOutput)
}

type ResourceMoveChangeHistoryOutput struct{ *pulumi.OutputState }

func (ResourceMoveChangeHistoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceMoveChangeHistory)(nil)).Elem()
}

func (o ResourceMoveChangeHistoryOutput) ToResourceMoveChangeHistoryOutput() ResourceMoveChangeHistoryOutput {
	return o
}

func (o ResourceMoveChangeHistoryOutput) ToResourceMoveChangeHistoryOutputWithContext(ctx context.Context) ResourceMoveChangeHistoryOutput {
	return o
}

func (o ResourceMoveChangeHistoryOutput) AzureSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMoveChangeHistory) *string { return v.AzureSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o ResourceMoveChangeHistoryOutput) ChangedTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMoveChangeHistory) *string { return v.ChangedTimeUtc }).(pulumi.StringPtrOutput)
}

func (o ResourceMoveChangeHistoryOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMoveChangeHistory) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

type ResourceMoveChangeHistoryArrayOutput struct{ *pulumi.OutputState }

func (ResourceMoveChangeHistoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceMoveChangeHistory)(nil)).Elem()
}

func (o ResourceMoveChangeHistoryArrayOutput) ToResourceMoveChangeHistoryArrayOutput() ResourceMoveChangeHistoryArrayOutput {
	return o
}

func (o ResourceMoveChangeHistoryArrayOutput) ToResourceMoveChangeHistoryArrayOutputWithContext(ctx context.Context) ResourceMoveChangeHistoryArrayOutput {
	return o
}

func (o ResourceMoveChangeHistoryArrayOutput) Index(i pulumi.IntInput) ResourceMoveChangeHistoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceMoveChangeHistory {
		return vs[0].([]ResourceMoveChangeHistory)[vs[1].(int)]
	}).(ResourceMoveChangeHistoryOutput)
}

type ResourceMoveChangeHistoryResponse struct {
	AzureSubscriptionId *string `pulumi:"azureSubscriptionId"`
	ChangedTimeUtc      *string `pulumi:"changedTimeUtc"`
	ResourceGroupName   *string `pulumi:"resourceGroupName"`
}

type ResourceMoveChangeHistoryResponseOutput struct{ *pulumi.OutputState }

func (ResourceMoveChangeHistoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceMoveChangeHistoryResponse)(nil)).Elem()
}

func (o ResourceMoveChangeHistoryResponseOutput) ToResourceMoveChangeHistoryResponseOutput() ResourceMoveChangeHistoryResponseOutput {
	return o
}

func (o ResourceMoveChangeHistoryResponseOutput) ToResourceMoveChangeHistoryResponseOutputWithContext(ctx context.Context) ResourceMoveChangeHistoryResponseOutput {
	return o
}

func (o ResourceMoveChangeHistoryResponseOutput) AzureSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMoveChangeHistoryResponse) *string { return v.AzureSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o ResourceMoveChangeHistoryResponseOutput) ChangedTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMoveChangeHistoryResponse) *string { return v.ChangedTimeUtc }).(pulumi.StringPtrOutput)
}

func (o ResourceMoveChangeHistoryResponseOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMoveChangeHistoryResponse) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

type ResourceMoveChangeHistoryResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceMoveChangeHistoryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceMoveChangeHistoryResponse)(nil)).Elem()
}

func (o ResourceMoveChangeHistoryResponseArrayOutput) ToResourceMoveChangeHistoryResponseArrayOutput() ResourceMoveChangeHistoryResponseArrayOutput {
	return o
}

func (o ResourceMoveChangeHistoryResponseArrayOutput) ToResourceMoveChangeHistoryResponseArrayOutputWithContext(ctx context.Context) ResourceMoveChangeHistoryResponseArrayOutput {
	return o
}

func (o ResourceMoveChangeHistoryResponseArrayOutput) Index(i pulumi.IntInput) ResourceMoveChangeHistoryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceMoveChangeHistoryResponse {
		return vs[0].([]ResourceMoveChangeHistoryResponse)[vs[1].(int)]
	}).(ResourceMoveChangeHistoryResponseOutput)
}

type ResourceSku struct {
	Name *string `pulumi:"name"`
}


func (val *ResourceSku) Defaults() *ResourceSku {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Name) {
		name_ := "Basic"
		tmp.Name = &name_
	}
	return &tmp
}





type ResourceSkuInput interface {
	pulumi.Input

	ToResourceSkuOutput() ResourceSkuOutput
	ToResourceSkuOutputWithContext(context.Context) ResourceSkuOutput
}

type ResourceSkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}


func (val *ResourceSkuArgs) Defaults() *ResourceSkuArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Name) {
		tmp.Name = pulumi.StringPtr("Basic")
	}
	return &tmp
}
func (ResourceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (i ResourceSkuArgs) ToResourceSkuOutput() ResourceSkuOutput {
	return i.ToResourceSkuOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput)
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput).ToResourceSkuPtrOutputWithContext(ctx)
}









type ResourceSkuPtrInput interface {
	pulumi.Input

	ToResourceSkuPtrOutput() ResourceSkuPtrOutput
	ToResourceSkuPtrOutputWithContext(context.Context) ResourceSkuPtrOutput
}

type resourceSkuPtrType ResourceSkuArgs

func ResourceSkuPtr(v *ResourceSkuArgs) ResourceSkuPtrInput {
	return (*resourceSkuPtrType)(v)
}

func (*resourceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuPtrOutput)
}

type ResourceSkuOutput struct{ *pulumi.OutputState }

func (ResourceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (o ResourceSkuOutput) ToResourceSkuOutput() ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSku) *ResourceSku {
		return &v
	}).(ResourceSkuPtrOutput)
}

func (o ResourceSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ResourceSkuPtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) Elem() ResourceSkuOutput {
	return o.ApplyT(func(v *ResourceSku) ResourceSku {
		if v != nil {
			return *v
		}
		var ret ResourceSku
		return ret
	}).(ResourceSkuOutput)
}

func (o ResourceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type ResourceSkuResponse struct {
	Name *string `pulumi:"name"`
}


func (val *ResourceSkuResponse) Defaults() *ResourceSkuResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Name) {
		name_ := "Basic"
		tmp.Name = &name_
	}
	return &tmp
}

type ResourceSkuResponseOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ResourceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) Elem() ResourceSkuResponseOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) ResourceSkuResponse {
		if v != nil {
			return *v
		}
		var ret ResourceSkuResponse
		return ret
	}).(ResourceSkuResponseOutput)
}

func (o ResourceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type RetryPolicy struct {
	EventTimeToLiveInMinutes *int `pulumi:"eventTimeToLiveInMinutes"`
	MaxDeliveryAttempts      *int `pulumi:"maxDeliveryAttempts"`
}


func (val *RetryPolicy) Defaults() *RetryPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EventTimeToLiveInMinutes) {
		eventTimeToLiveInMinutes_ := 1440
		tmp.EventTimeToLiveInMinutes = &eventTimeToLiveInMinutes_
	}
	if isZero(tmp.MaxDeliveryAttempts) {
		maxDeliveryAttempts_ := 30
		tmp.MaxDeliveryAttempts = &maxDeliveryAttempts_
	}
	return &tmp
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


func (val *RetryPolicyArgs) Defaults() *RetryPolicyArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EventTimeToLiveInMinutes) {
		tmp.EventTimeToLiveInMinutes = pulumi.IntPtr(1440)
	}
	if isZero(tmp.MaxDeliveryAttempts) {
		tmp.MaxDeliveryAttempts = pulumi.IntPtr(30)
	}
	return &tmp
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


func (val *RetryPolicyResponse) Defaults() *RetryPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EventTimeToLiveInMinutes) {
		eventTimeToLiveInMinutes_ := 1440
		tmp.EventTimeToLiveInMinutes = &eventTimeToLiveInMinutes_
	}
	if isZero(tmp.MaxDeliveryAttempts) {
		maxDeliveryAttempts_ := 30
		tmp.MaxDeliveryAttempts = &maxDeliveryAttempts_
	}
	return &tmp
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

type ServiceBusQueueEventSubscriptionDestinationResponse struct {
	DeliveryAttributeMappings []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType              string        `pulumi:"endpointType"`
	ResourceId                *string       `pulumi:"resourceId"`
}

type ServiceBusTopicEventSubscriptionDestination struct {
	DeliveryAttributeMappings []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType              string        `pulumi:"endpointType"`
	ResourceId                *string       `pulumi:"resourceId"`
}

type ServiceBusTopicEventSubscriptionDestinationResponse struct {
	DeliveryAttributeMappings []interface{} `pulumi:"deliveryAttributeMappings"`
	EndpointType              string        `pulumi:"endpointType"`
	ResourceId                *string       `pulumi:"resourceId"`
}

type StaticDeliveryAttributeMapping struct {
	IsSecret *bool   `pulumi:"isSecret"`
	Name     *string `pulumi:"name"`
	Type     string  `pulumi:"type"`
	Value    *string `pulumi:"value"`
}


func (val *StaticDeliveryAttributeMapping) Defaults() *StaticDeliveryAttributeMapping {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSecret) {
		isSecret_ := false
		tmp.IsSecret = &isSecret_
	}
	return &tmp
}

type StaticDeliveryAttributeMappingResponse struct {
	IsSecret *bool   `pulumi:"isSecret"`
	Name     *string `pulumi:"name"`
	Type     string  `pulumi:"type"`
	Value    *string `pulumi:"value"`
}


func (val *StaticDeliveryAttributeMappingResponse) Defaults() *StaticDeliveryAttributeMappingResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSecret) {
		isSecret_ := false
		tmp.IsSecret = &isSecret_
	}
	return &tmp
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

type StorageQueueEventSubscriptionDestinationResponse struct {
	EndpointType                    string   `pulumi:"endpointType"`
	QueueMessageTimeToLiveInSeconds *float64 `pulumi:"queueMessageTimeToLiveInSeconds"`
	QueueName                       *string  `pulumi:"queueName"`
	ResourceId                      *string  `pulumi:"resourceId"`
}

type StringBeginsWithAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringBeginsWithAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringContainsAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringContainsAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringEndsWithAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringEndsWithAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringInAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringInAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringNotBeginsWithAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringNotBeginsWithAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringNotContainsAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringNotContainsAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringNotEndsWithAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringNotEndsWithAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringNotInAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringNotInAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
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


func (val *WebHookEventSubscriptionDestination) Defaults() *WebHookEventSubscriptionDestination {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxEventsPerBatch) {
		maxEventsPerBatch_ := 1
		tmp.MaxEventsPerBatch = &maxEventsPerBatch_
	}
	if isZero(tmp.PreferredBatchSizeInKilobytes) {
		preferredBatchSizeInKilobytes_ := 64
		tmp.PreferredBatchSizeInKilobytes = &preferredBatchSizeInKilobytes_
	}
	return &tmp
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


func (val *WebHookEventSubscriptionDestinationResponse) Defaults() *WebHookEventSubscriptionDestinationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxEventsPerBatch) {
		maxEventsPerBatch_ := 1
		tmp.MaxEventsPerBatch = &maxEventsPerBatch_
	}
	if isZero(tmp.PreferredBatchSizeInKilobytes) {
		preferredBatchSizeInKilobytes_ := 64
		tmp.PreferredBatchSizeInKilobytes = &preferredBatchSizeInKilobytes_
	}
	return &tmp
}

type WebhookPartnerDestinationInfo struct {
	AzureSubscriptionId       *string                             `pulumi:"azureSubscriptionId"`
	ClientAuthentication      *AzureADPartnerClientAuthentication `pulumi:"clientAuthentication"`
	EndpointBaseUrl           *string                             `pulumi:"endpointBaseUrl"`
	EndpointServiceContext    *string                             `pulumi:"endpointServiceContext"`
	EndpointType              string                              `pulumi:"endpointType"`
	EndpointUrl               *string                             `pulumi:"endpointUrl"`
	Name                      *string                             `pulumi:"name"`
	ResourceGroupName         *string                             `pulumi:"resourceGroupName"`
	ResourceMoveChangeHistory []ResourceMoveChangeHistory         `pulumi:"resourceMoveChangeHistory"`
}


func (val *WebhookPartnerDestinationInfo) Defaults() *WebhookPartnerDestinationInfo {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ClientAuthentication = tmp.ClientAuthentication.Defaults()

	if isZero(tmp.EndpointType) {
		tmp.EndpointType = "WebHook"
	}
	return &tmp
}





type WebhookPartnerDestinationInfoInput interface {
	pulumi.Input

	ToWebhookPartnerDestinationInfoOutput() WebhookPartnerDestinationInfoOutput
	ToWebhookPartnerDestinationInfoOutputWithContext(context.Context) WebhookPartnerDestinationInfoOutput
}

type WebhookPartnerDestinationInfoArgs struct {
	AzureSubscriptionId       pulumi.StringPtrInput                      `pulumi:"azureSubscriptionId"`
	ClientAuthentication      AzureADPartnerClientAuthenticationPtrInput `pulumi:"clientAuthentication"`
	EndpointBaseUrl           pulumi.StringPtrInput                      `pulumi:"endpointBaseUrl"`
	EndpointServiceContext    pulumi.StringPtrInput                      `pulumi:"endpointServiceContext"`
	EndpointType              pulumi.StringInput                         `pulumi:"endpointType"`
	EndpointUrl               pulumi.StringPtrInput                      `pulumi:"endpointUrl"`
	Name                      pulumi.StringPtrInput                      `pulumi:"name"`
	ResourceGroupName         pulumi.StringPtrInput                      `pulumi:"resourceGroupName"`
	ResourceMoveChangeHistory ResourceMoveChangeHistoryArrayInput        `pulumi:"resourceMoveChangeHistory"`
}


func (val *WebhookPartnerDestinationInfoArgs) Defaults() *WebhookPartnerDestinationInfoArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	if isZero(tmp.EndpointType) {
		tmp.EndpointType = pulumi.String("WebHook")
	}
	return &tmp
}
func (WebhookPartnerDestinationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookPartnerDestinationInfo)(nil)).Elem()
}

func (i WebhookPartnerDestinationInfoArgs) ToWebhookPartnerDestinationInfoOutput() WebhookPartnerDestinationInfoOutput {
	return i.ToWebhookPartnerDestinationInfoOutputWithContext(context.Background())
}

func (i WebhookPartnerDestinationInfoArgs) ToWebhookPartnerDestinationInfoOutputWithContext(ctx context.Context) WebhookPartnerDestinationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookPartnerDestinationInfoOutput)
}

func (i WebhookPartnerDestinationInfoArgs) ToWebhookPartnerDestinationInfoPtrOutput() WebhookPartnerDestinationInfoPtrOutput {
	return i.ToWebhookPartnerDestinationInfoPtrOutputWithContext(context.Background())
}

func (i WebhookPartnerDestinationInfoArgs) ToWebhookPartnerDestinationInfoPtrOutputWithContext(ctx context.Context) WebhookPartnerDestinationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookPartnerDestinationInfoOutput).ToWebhookPartnerDestinationInfoPtrOutputWithContext(ctx)
}









type WebhookPartnerDestinationInfoPtrInput interface {
	pulumi.Input

	ToWebhookPartnerDestinationInfoPtrOutput() WebhookPartnerDestinationInfoPtrOutput
	ToWebhookPartnerDestinationInfoPtrOutputWithContext(context.Context) WebhookPartnerDestinationInfoPtrOutput
}

type webhookPartnerDestinationInfoPtrType WebhookPartnerDestinationInfoArgs

func WebhookPartnerDestinationInfoPtr(v *WebhookPartnerDestinationInfoArgs) WebhookPartnerDestinationInfoPtrInput {
	return (*webhookPartnerDestinationInfoPtrType)(v)
}

func (*webhookPartnerDestinationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookPartnerDestinationInfo)(nil)).Elem()
}

func (i *webhookPartnerDestinationInfoPtrType) ToWebhookPartnerDestinationInfoPtrOutput() WebhookPartnerDestinationInfoPtrOutput {
	return i.ToWebhookPartnerDestinationInfoPtrOutputWithContext(context.Background())
}

func (i *webhookPartnerDestinationInfoPtrType) ToWebhookPartnerDestinationInfoPtrOutputWithContext(ctx context.Context) WebhookPartnerDestinationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookPartnerDestinationInfoPtrOutput)
}

type WebhookPartnerDestinationInfoOutput struct{ *pulumi.OutputState }

func (WebhookPartnerDestinationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookPartnerDestinationInfo)(nil)).Elem()
}

func (o WebhookPartnerDestinationInfoOutput) ToWebhookPartnerDestinationInfoOutput() WebhookPartnerDestinationInfoOutput {
	return o
}

func (o WebhookPartnerDestinationInfoOutput) ToWebhookPartnerDestinationInfoOutputWithContext(ctx context.Context) WebhookPartnerDestinationInfoOutput {
	return o
}

func (o WebhookPartnerDestinationInfoOutput) ToWebhookPartnerDestinationInfoPtrOutput() WebhookPartnerDestinationInfoPtrOutput {
	return o.ToWebhookPartnerDestinationInfoPtrOutputWithContext(context.Background())
}

func (o WebhookPartnerDestinationInfoOutput) ToWebhookPartnerDestinationInfoPtrOutputWithContext(ctx context.Context) WebhookPartnerDestinationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebhookPartnerDestinationInfo) *WebhookPartnerDestinationInfo {
		return &v
	}).(WebhookPartnerDestinationInfoPtrOutput)
}

func (o WebhookPartnerDestinationInfoOutput) AzureSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfo) *string { return v.AzureSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoOutput) ClientAuthentication() AzureADPartnerClientAuthenticationPtrOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfo) *AzureADPartnerClientAuthentication {
		return v.ClientAuthentication
	}).(AzureADPartnerClientAuthenticationPtrOutput)
}

func (o WebhookPartnerDestinationInfoOutput) EndpointBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfo) *string { return v.EndpointBaseUrl }).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoOutput) EndpointServiceContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfo) *string { return v.EndpointServiceContext }).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfo) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o WebhookPartnerDestinationInfoOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfo) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfo) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfo) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoOutput) ResourceMoveChangeHistory() ResourceMoveChangeHistoryArrayOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfo) []ResourceMoveChangeHistory { return v.ResourceMoveChangeHistory }).(ResourceMoveChangeHistoryArrayOutput)
}

type WebhookPartnerDestinationInfoPtrOutput struct{ *pulumi.OutputState }

func (WebhookPartnerDestinationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookPartnerDestinationInfo)(nil)).Elem()
}

func (o WebhookPartnerDestinationInfoPtrOutput) ToWebhookPartnerDestinationInfoPtrOutput() WebhookPartnerDestinationInfoPtrOutput {
	return o
}

func (o WebhookPartnerDestinationInfoPtrOutput) ToWebhookPartnerDestinationInfoPtrOutputWithContext(ctx context.Context) WebhookPartnerDestinationInfoPtrOutput {
	return o
}

func (o WebhookPartnerDestinationInfoPtrOutput) Elem() WebhookPartnerDestinationInfoOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfo) WebhookPartnerDestinationInfo {
		if v != nil {
			return *v
		}
		var ret WebhookPartnerDestinationInfo
		return ret
	}).(WebhookPartnerDestinationInfoOutput)
}

func (o WebhookPartnerDestinationInfoPtrOutput) AzureSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfo) *string {
		if v == nil {
			return nil
		}
		return v.AzureSubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoPtrOutput) ClientAuthentication() AzureADPartnerClientAuthenticationPtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfo) *AzureADPartnerClientAuthentication {
		if v == nil {
			return nil
		}
		return v.ClientAuthentication
	}).(AzureADPartnerClientAuthenticationPtrOutput)
}

func (o WebhookPartnerDestinationInfoPtrOutput) EndpointBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfo) *string {
		if v == nil {
			return nil
		}
		return v.EndpointBaseUrl
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoPtrOutput) EndpointServiceContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfo) *string {
		if v == nil {
			return nil
		}
		return v.EndpointServiceContext
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoPtrOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfo) *string {
		if v == nil {
			return nil
		}
		return &v.EndpointType
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoPtrOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfo) *string {
		if v == nil {
			return nil
		}
		return v.EndpointUrl
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfo) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoPtrOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfo) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroupName
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoPtrOutput) ResourceMoveChangeHistory() ResourceMoveChangeHistoryArrayOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfo) []ResourceMoveChangeHistory {
		if v == nil {
			return nil
		}
		return v.ResourceMoveChangeHistory
	}).(ResourceMoveChangeHistoryArrayOutput)
}

type WebhookPartnerDestinationInfoResponse struct {
	AzureSubscriptionId       *string                                     `pulumi:"azureSubscriptionId"`
	ClientAuthentication      *AzureADPartnerClientAuthenticationResponse `pulumi:"clientAuthentication"`
	EndpointBaseUrl           *string                                     `pulumi:"endpointBaseUrl"`
	EndpointServiceContext    *string                                     `pulumi:"endpointServiceContext"`
	EndpointType              string                                      `pulumi:"endpointType"`
	EndpointUrl               *string                                     `pulumi:"endpointUrl"`
	Name                      *string                                     `pulumi:"name"`
	ResourceGroupName         *string                                     `pulumi:"resourceGroupName"`
	ResourceMoveChangeHistory []ResourceMoveChangeHistoryResponse         `pulumi:"resourceMoveChangeHistory"`
}


func (val *WebhookPartnerDestinationInfoResponse) Defaults() *WebhookPartnerDestinationInfoResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ClientAuthentication = tmp.ClientAuthentication.Defaults()

	if isZero(tmp.EndpointType) {
		tmp.EndpointType = "WebHook"
	}
	return &tmp
}

type WebhookPartnerDestinationInfoResponseOutput struct{ *pulumi.OutputState }

func (WebhookPartnerDestinationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookPartnerDestinationInfoResponse)(nil)).Elem()
}

func (o WebhookPartnerDestinationInfoResponseOutput) ToWebhookPartnerDestinationInfoResponseOutput() WebhookPartnerDestinationInfoResponseOutput {
	return o
}

func (o WebhookPartnerDestinationInfoResponseOutput) ToWebhookPartnerDestinationInfoResponseOutputWithContext(ctx context.Context) WebhookPartnerDestinationInfoResponseOutput {
	return o
}

func (o WebhookPartnerDestinationInfoResponseOutput) AzureSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfoResponse) *string { return v.AzureSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoResponseOutput) ClientAuthentication() AzureADPartnerClientAuthenticationResponsePtrOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfoResponse) *AzureADPartnerClientAuthenticationResponse {
		return v.ClientAuthentication
	}).(AzureADPartnerClientAuthenticationResponsePtrOutput)
}

func (o WebhookPartnerDestinationInfoResponseOutput) EndpointBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfoResponse) *string { return v.EndpointBaseUrl }).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoResponseOutput) EndpointServiceContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfoResponse) *string { return v.EndpointServiceContext }).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfoResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o WebhookPartnerDestinationInfoResponseOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfoResponse) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoResponseOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfoResponse) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoResponseOutput) ResourceMoveChangeHistory() ResourceMoveChangeHistoryResponseArrayOutput {
	return o.ApplyT(func(v WebhookPartnerDestinationInfoResponse) []ResourceMoveChangeHistoryResponse {
		return v.ResourceMoveChangeHistory
	}).(ResourceMoveChangeHistoryResponseArrayOutput)
}

type WebhookPartnerDestinationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (WebhookPartnerDestinationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookPartnerDestinationInfoResponse)(nil)).Elem()
}

func (o WebhookPartnerDestinationInfoResponsePtrOutput) ToWebhookPartnerDestinationInfoResponsePtrOutput() WebhookPartnerDestinationInfoResponsePtrOutput {
	return o
}

func (o WebhookPartnerDestinationInfoResponsePtrOutput) ToWebhookPartnerDestinationInfoResponsePtrOutputWithContext(ctx context.Context) WebhookPartnerDestinationInfoResponsePtrOutput {
	return o
}

func (o WebhookPartnerDestinationInfoResponsePtrOutput) Elem() WebhookPartnerDestinationInfoResponseOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfoResponse) WebhookPartnerDestinationInfoResponse {
		if v != nil {
			return *v
		}
		var ret WebhookPartnerDestinationInfoResponse
		return ret
	}).(WebhookPartnerDestinationInfoResponseOutput)
}

func (o WebhookPartnerDestinationInfoResponsePtrOutput) AzureSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.AzureSubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoResponsePtrOutput) ClientAuthentication() AzureADPartnerClientAuthenticationResponsePtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfoResponse) *AzureADPartnerClientAuthenticationResponse {
		if v == nil {
			return nil
		}
		return v.ClientAuthentication
	}).(AzureADPartnerClientAuthenticationResponsePtrOutput)
}

func (o WebhookPartnerDestinationInfoResponsePtrOutput) EndpointBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndpointBaseUrl
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoResponsePtrOutput) EndpointServiceContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndpointServiceContext
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoResponsePtrOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EndpointType
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoResponsePtrOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndpointUrl
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoResponsePtrOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroupName
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPartnerDestinationInfoResponsePtrOutput) ResourceMoveChangeHistory() ResourceMoveChangeHistoryResponseArrayOutput {
	return o.ApplyT(func(v *WebhookPartnerDestinationInfoResponse) []ResourceMoveChangeHistoryResponse {
		if v == nil {
			return nil
		}
		return v.ResourceMoveChangeHistory
	}).(ResourceMoveChangeHistoryResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureADPartnerClientAuthenticationOutput{})
	pulumi.RegisterOutputType(AzureADPartnerClientAuthenticationPtrOutput{})
	pulumi.RegisterOutputType(AzureADPartnerClientAuthenticationResponseOutput{})
	pulumi.RegisterOutputType(AzureADPartnerClientAuthenticationResponsePtrOutput{})
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
	pulumi.RegisterOutputType(EventChannelDestinationOutput{})
	pulumi.RegisterOutputType(EventChannelDestinationPtrOutput{})
	pulumi.RegisterOutputType(EventChannelDestinationResponseOutput{})
	pulumi.RegisterOutputType(EventChannelDestinationResponsePtrOutput{})
	pulumi.RegisterOutputType(EventChannelFilterOutput{})
	pulumi.RegisterOutputType(EventChannelFilterPtrOutput{})
	pulumi.RegisterOutputType(EventChannelFilterResponseOutput{})
	pulumi.RegisterOutputType(EventChannelFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(EventChannelSourceOutput{})
	pulumi.RegisterOutputType(EventChannelSourcePtrOutput{})
	pulumi.RegisterOutputType(EventChannelSourceResponseOutput{})
	pulumi.RegisterOutputType(EventChannelSourceResponsePtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterPtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterResponseOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionIdentityOutput{})
	pulumi.RegisterOutputType(EventSubscriptionIdentityPtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionIdentityResponseOutput{})
	pulumi.RegisterOutputType(EventSubscriptionIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(EventTypeInfoOutput{})
	pulumi.RegisterOutputType(EventTypeInfoPtrOutput{})
	pulumi.RegisterOutputType(EventTypeInfoResponseOutput{})
	pulumi.RegisterOutputType(EventTypeInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityInfoOutput{})
	pulumi.RegisterOutputType(IdentityInfoPtrOutput{})
	pulumi.RegisterOutputType(IdentityInfoResponseOutput{})
	pulumi.RegisterOutputType(IdentityInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(InboundIpRuleOutput{})
	pulumi.RegisterOutputType(InboundIpRuleArrayOutput{})
	pulumi.RegisterOutputType(InboundIpRuleResponseOutput{})
	pulumi.RegisterOutputType(InboundIpRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(InlineEventPropertiesOutput{})
	pulumi.RegisterOutputType(InlineEventPropertiesMapOutput{})
	pulumi.RegisterOutputType(InlineEventPropertiesResponseOutput{})
	pulumi.RegisterOutputType(InlineEventPropertiesResponseMapOutput{})
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
	pulumi.RegisterOutputType(PartnerOutput{})
	pulumi.RegisterOutputType(PartnerArrayOutput{})
	pulumi.RegisterOutputType(PartnerAuthorizationOutput{})
	pulumi.RegisterOutputType(PartnerAuthorizationPtrOutput{})
	pulumi.RegisterOutputType(PartnerAuthorizationResponseOutput{})
	pulumi.RegisterOutputType(PartnerAuthorizationResponsePtrOutput{})
	pulumi.RegisterOutputType(PartnerResponseOutput{})
	pulumi.RegisterOutputType(PartnerResponseArrayOutput{})
	pulumi.RegisterOutputType(PartnerTopicInfoOutput{})
	pulumi.RegisterOutputType(PartnerTopicInfoPtrOutput{})
	pulumi.RegisterOutputType(PartnerTopicInfoResponseOutput{})
	pulumi.RegisterOutputType(PartnerTopicInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceMoveChangeHistoryOutput{})
	pulumi.RegisterOutputType(ResourceMoveChangeHistoryArrayOutput{})
	pulumi.RegisterOutputType(ResourceMoveChangeHistoryResponseOutput{})
	pulumi.RegisterOutputType(ResourceMoveChangeHistoryResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuPtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(RetryPolicyOutput{})
	pulumi.RegisterOutputType(RetryPolicyPtrOutput{})
	pulumi.RegisterOutputType(RetryPolicyResponseOutput{})
	pulumi.RegisterOutputType(RetryPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationPtrOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationResponseOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesMapOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseMapOutput{})
	pulumi.RegisterOutputType(WebhookPartnerDestinationInfoOutput{})
	pulumi.RegisterOutputType(WebhookPartnerDestinationInfoPtrOutput{})
	pulumi.RegisterOutputType(WebhookPartnerDestinationInfoResponseOutput{})
	pulumi.RegisterOutputType(WebhookPartnerDestinationInfoResponsePtrOutput{})
}

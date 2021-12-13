


package digitaltwins

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectionPropertiesPrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     string  `pulumi:"description"`
	Status          string  `pulumi:"status"`
}





type ConnectionPropertiesPrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToConnectionPropertiesPrivateLinkServiceConnectionStateOutput() ConnectionPropertiesPrivateLinkServiceConnectionStateOutput
	ToConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStateOutput
}

type ConnectionPropertiesPrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringInput    `pulumi:"description"`
	Status          pulumi.StringInput    `pulumi:"status"`
}

func (ConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i ConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToConnectionPropertiesPrivateLinkServiceConnectionStateOutput() ConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return i.ToConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i ConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesPrivateLinkServiceConnectionStateOutput)
}

func (i ConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return i.ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i ConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesPrivateLinkServiceConnectionStateOutput).ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type ConnectionPropertiesPrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput
	ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput
}

type connectionPropertiesPrivateLinkServiceConnectionStatePtrType ConnectionPropertiesPrivateLinkServiceConnectionStateArgs

func ConnectionPropertiesPrivateLinkServiceConnectionStatePtr(v *ConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ConnectionPropertiesPrivateLinkServiceConnectionStatePtrInput {
	return (*connectionPropertiesPrivateLinkServiceConnectionStatePtrType)(v)
}

func (*connectionPropertiesPrivateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *connectionPropertiesPrivateLinkServiceConnectionStatePtrType) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return i.ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *connectionPropertiesPrivateLinkServiceConnectionStatePtrType) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
}

type ConnectionPropertiesPrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToConnectionPropertiesPrivateLinkServiceConnectionStateOutput() ConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return o
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return o
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionPropertiesPrivateLinkServiceConnectionState) *ConnectionPropertiesPrivateLinkServiceConnectionState {
		return &v
	}).(ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionPropertiesPrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionPropertiesPrivateLinkServiceConnectionState) string { return v.Description }).(pulumi.StringOutput)
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionPropertiesPrivateLinkServiceConnectionState) string { return v.Status }).(pulumi.StringOutput)
}

type ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) Elem() ConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *ConnectionPropertiesPrivateLinkServiceConnectionState) ConnectionPropertiesPrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret ConnectionPropertiesPrivateLinkServiceConnectionState
		return ret
	}).(ConnectionPropertiesPrivateLinkServiceConnectionStateOutput)
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionPropertiesPrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionPropertiesPrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionPropertiesPrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type ConnectionPropertiesResponsePrivateEndpoint struct {
	Id string `pulumi:"id"`
}

type ConnectionPropertiesResponsePrivateEndpointOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesResponsePrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionPropertiesResponsePrivateEndpoint)(nil)).Elem()
}

func (o ConnectionPropertiesResponsePrivateEndpointOutput) ToConnectionPropertiesResponsePrivateEndpointOutput() ConnectionPropertiesResponsePrivateEndpointOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateEndpointOutput) ToConnectionPropertiesResponsePrivateEndpointOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateEndpointOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateEndpointOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionPropertiesResponsePrivateEndpoint) string { return v.Id }).(pulumi.StringOutput)
}

type ConnectionPropertiesResponsePrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesResponsePrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionPropertiesResponsePrivateEndpoint)(nil)).Elem()
}

func (o ConnectionPropertiesResponsePrivateEndpointPtrOutput) ToConnectionPropertiesResponsePrivateEndpointPtrOutput() ConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateEndpointPtrOutput) ToConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateEndpointPtrOutput) Elem() ConnectionPropertiesResponsePrivateEndpointOutput {
	return o.ApplyT(func(v *ConnectionPropertiesResponsePrivateEndpoint) ConnectionPropertiesResponsePrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret ConnectionPropertiesResponsePrivateEndpoint
		return ret
	}).(ConnectionPropertiesResponsePrivateEndpointOutput)
}

func (o ConnectionPropertiesResponsePrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionPropertiesResponsePrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ConnectionPropertiesResponsePrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     string  `pulumi:"description"`
	Status          string  `pulumi:"status"`
}

type ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionPropertiesResponsePrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput() ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionPropertiesResponsePrivateLinkServiceConnectionState) string { return v.Description }).(pulumi.StringOutput)
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionPropertiesResponsePrivateLinkServiceConnectionState) string { return v.Status }).(pulumi.StringOutput)
}

type ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionPropertiesResponsePrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) Elem() ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *ConnectionPropertiesResponsePrivateLinkServiceConnectionState) ConnectionPropertiesResponsePrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret ConnectionPropertiesResponsePrivateLinkServiceConnectionState
		return ret
	}).(ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput)
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type DigitalTwinsIdentity struct {
	Type *string `pulumi:"type"`
}





type DigitalTwinsIdentityInput interface {
	pulumi.Input

	ToDigitalTwinsIdentityOutput() DigitalTwinsIdentityOutput
	ToDigitalTwinsIdentityOutputWithContext(context.Context) DigitalTwinsIdentityOutput
}

type DigitalTwinsIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (DigitalTwinsIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwinsIdentity)(nil)).Elem()
}

func (i DigitalTwinsIdentityArgs) ToDigitalTwinsIdentityOutput() DigitalTwinsIdentityOutput {
	return i.ToDigitalTwinsIdentityOutputWithContext(context.Background())
}

func (i DigitalTwinsIdentityArgs) ToDigitalTwinsIdentityOutputWithContext(ctx context.Context) DigitalTwinsIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinsIdentityOutput)
}

func (i DigitalTwinsIdentityArgs) ToDigitalTwinsIdentityPtrOutput() DigitalTwinsIdentityPtrOutput {
	return i.ToDigitalTwinsIdentityPtrOutputWithContext(context.Background())
}

func (i DigitalTwinsIdentityArgs) ToDigitalTwinsIdentityPtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinsIdentityOutput).ToDigitalTwinsIdentityPtrOutputWithContext(ctx)
}









type DigitalTwinsIdentityPtrInput interface {
	pulumi.Input

	ToDigitalTwinsIdentityPtrOutput() DigitalTwinsIdentityPtrOutput
	ToDigitalTwinsIdentityPtrOutputWithContext(context.Context) DigitalTwinsIdentityPtrOutput
}

type digitalTwinsIdentityPtrType DigitalTwinsIdentityArgs

func DigitalTwinsIdentityPtr(v *DigitalTwinsIdentityArgs) DigitalTwinsIdentityPtrInput {
	return (*digitalTwinsIdentityPtrType)(v)
}

func (*digitalTwinsIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwinsIdentity)(nil)).Elem()
}

func (i *digitalTwinsIdentityPtrType) ToDigitalTwinsIdentityPtrOutput() DigitalTwinsIdentityPtrOutput {
	return i.ToDigitalTwinsIdentityPtrOutputWithContext(context.Background())
}

func (i *digitalTwinsIdentityPtrType) ToDigitalTwinsIdentityPtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinsIdentityPtrOutput)
}

type DigitalTwinsIdentityOutput struct{ *pulumi.OutputState }

func (DigitalTwinsIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwinsIdentity)(nil)).Elem()
}

func (o DigitalTwinsIdentityOutput) ToDigitalTwinsIdentityOutput() DigitalTwinsIdentityOutput {
	return o
}

func (o DigitalTwinsIdentityOutput) ToDigitalTwinsIdentityOutputWithContext(ctx context.Context) DigitalTwinsIdentityOutput {
	return o
}

func (o DigitalTwinsIdentityOutput) ToDigitalTwinsIdentityPtrOutput() DigitalTwinsIdentityPtrOutput {
	return o.ToDigitalTwinsIdentityPtrOutputWithContext(context.Background())
}

func (o DigitalTwinsIdentityOutput) ToDigitalTwinsIdentityPtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DigitalTwinsIdentity) *DigitalTwinsIdentity {
		return &v
	}).(DigitalTwinsIdentityPtrOutput)
}

func (o DigitalTwinsIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DigitalTwinsIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type DigitalTwinsIdentityPtrOutput struct{ *pulumi.OutputState }

func (DigitalTwinsIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwinsIdentity)(nil)).Elem()
}

func (o DigitalTwinsIdentityPtrOutput) ToDigitalTwinsIdentityPtrOutput() DigitalTwinsIdentityPtrOutput {
	return o
}

func (o DigitalTwinsIdentityPtrOutput) ToDigitalTwinsIdentityPtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityPtrOutput {
	return o
}

func (o DigitalTwinsIdentityPtrOutput) Elem() DigitalTwinsIdentityOutput {
	return o.ApplyT(func(v *DigitalTwinsIdentity) DigitalTwinsIdentity {
		if v != nil {
			return *v
		}
		var ret DigitalTwinsIdentity
		return ret
	}).(DigitalTwinsIdentityOutput)
}

func (o DigitalTwinsIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DigitalTwinsIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type DigitalTwinsIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type DigitalTwinsIdentityResponseOutput struct{ *pulumi.OutputState }

func (DigitalTwinsIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwinsIdentityResponse)(nil)).Elem()
}

func (o DigitalTwinsIdentityResponseOutput) ToDigitalTwinsIdentityResponseOutput() DigitalTwinsIdentityResponseOutput {
	return o
}

func (o DigitalTwinsIdentityResponseOutput) ToDigitalTwinsIdentityResponseOutputWithContext(ctx context.Context) DigitalTwinsIdentityResponseOutput {
	return o
}

func (o DigitalTwinsIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v DigitalTwinsIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o DigitalTwinsIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v DigitalTwinsIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o DigitalTwinsIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DigitalTwinsIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type DigitalTwinsIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (DigitalTwinsIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwinsIdentityResponse)(nil)).Elem()
}

func (o DigitalTwinsIdentityResponsePtrOutput) ToDigitalTwinsIdentityResponsePtrOutput() DigitalTwinsIdentityResponsePtrOutput {
	return o
}

func (o DigitalTwinsIdentityResponsePtrOutput) ToDigitalTwinsIdentityResponsePtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityResponsePtrOutput {
	return o
}

func (o DigitalTwinsIdentityResponsePtrOutput) Elem() DigitalTwinsIdentityResponseOutput {
	return o.ApplyT(func(v *DigitalTwinsIdentityResponse) DigitalTwinsIdentityResponse {
		if v != nil {
			return *v
		}
		var ret DigitalTwinsIdentityResponse
		return ret
	}).(DigitalTwinsIdentityResponseOutput)
}

func (o DigitalTwinsIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DigitalTwinsIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o DigitalTwinsIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DigitalTwinsIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o DigitalTwinsIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DigitalTwinsIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type EventGrid struct {
	AccessKey1         string  `pulumi:"accessKey1"`
	AccessKey2         *string `pulumi:"accessKey2"`
	AuthenticationType *string `pulumi:"authenticationType"`
	DeadLetterSecret   *string `pulumi:"deadLetterSecret"`
	DeadLetterUri      *string `pulumi:"deadLetterUri"`
	EndpointType       string  `pulumi:"endpointType"`
	TopicEndpoint      string  `pulumi:"topicEndpoint"`
}

type EventGridResponse struct {
	AccessKey1         string  `pulumi:"accessKey1"`
	AccessKey2         *string `pulumi:"accessKey2"`
	AuthenticationType *string `pulumi:"authenticationType"`
	CreatedTime        string  `pulumi:"createdTime"`
	DeadLetterSecret   *string `pulumi:"deadLetterSecret"`
	DeadLetterUri      *string `pulumi:"deadLetterUri"`
	EndpointType       string  `pulumi:"endpointType"`
	ProvisioningState  string  `pulumi:"provisioningState"`
	TopicEndpoint      string  `pulumi:"topicEndpoint"`
}

type EventHub struct {
	AuthenticationType           *string `pulumi:"authenticationType"`
	ConnectionStringPrimaryKey   *string `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey *string `pulumi:"connectionStringSecondaryKey"`
	DeadLetterSecret             *string `pulumi:"deadLetterSecret"`
	DeadLetterUri                *string `pulumi:"deadLetterUri"`
	EndpointType                 string  `pulumi:"endpointType"`
	EndpointUri                  *string `pulumi:"endpointUri"`
	EntityPath                   *string `pulumi:"entityPath"`
}

type EventHubResponse struct {
	AuthenticationType           *string `pulumi:"authenticationType"`
	ConnectionStringPrimaryKey   *string `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey *string `pulumi:"connectionStringSecondaryKey"`
	CreatedTime                  string  `pulumi:"createdTime"`
	DeadLetterSecret             *string `pulumi:"deadLetterSecret"`
	DeadLetterUri                *string `pulumi:"deadLetterUri"`
	EndpointType                 string  `pulumi:"endpointType"`
	EndpointUri                  *string `pulumi:"endpointUri"`
	EntityPath                   *string `pulumi:"entityPath"`
	ProvisioningState            string  `pulumi:"provisioningState"`
}

type PrivateEndpointConnectionType struct {
	Properties PrivateEndpointConnectionProperties `pulumi:"properties"`
}





type PrivateEndpointConnectionTypeInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput
	ToPrivateEndpointConnectionTypeOutputWithContext(context.Context) PrivateEndpointConnectionTypeOutput
}

type PrivateEndpointConnectionTypeArgs struct {
	Properties PrivateEndpointConnectionPropertiesInput `pulumi:"properties"`
}

func (PrivateEndpointConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return i.ToPrivateEndpointConnectionTypeOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeOutput)
}





type PrivateEndpointConnectionTypeArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput
	ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Context) PrivateEndpointConnectionTypeArrayOutput
}

type PrivateEndpointConnectionTypeArray []PrivateEndpointConnectionTypeInput

func (PrivateEndpointConnectionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return i.ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeArrayOutput)
}

type PrivateEndpointConnectionTypeOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) Properties() PrivateEndpointConnectionPropertiesOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) PrivateEndpointConnectionProperties { return v.Properties }).(PrivateEndpointConnectionPropertiesOutput)
}

type PrivateEndpointConnectionTypeArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionType {
		return vs[0].([]PrivateEndpointConnectionType)[vs[1].(int)]
	}).(PrivateEndpointConnectionTypeOutput)
}

type PrivateEndpointConnectionProperties struct {
	GroupIds                          []string                                               `pulumi:"groupIds"`
	PrivateLinkServiceConnectionState *ConnectionPropertiesPrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionPropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput
	ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesOutput
}

type PrivateEndpointConnectionPropertiesArgs struct {
	GroupIds                          pulumi.StringArrayInput                                       `pulumi:"groupIds"`
	PrivateLinkServiceConnectionState ConnectionPropertiesPrivateLinkServiceConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return i.ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput)
}

type PrivateEndpointConnectionPropertiesOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) PrivateLinkServiceConnectionState() ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) *ConnectionPropertiesPrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionResponseProperties `pulumi:"properties"`
	Type       string                                      `pulumi:"type"`
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

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Properties() PrivateEndpointConnectionResponsePropertiesOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateEndpointConnectionResponseProperties {
		return v.Properties
	}).(PrivateEndpointConnectionResponsePropertiesOutput)
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

type PrivateEndpointConnectionResponseProperties struct {
	GroupIds                          []string                                                       `pulumi:"groupIds"`
	PrivateEndpoint                   *ConnectionPropertiesResponsePrivateEndpoint                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *ConnectionPropertiesResponsePrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                                         `pulumi:"provisioningState"`
}

type PrivateEndpointConnectionResponsePropertiesOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponseProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponsePropertiesOutput) ToPrivateEndpointConnectionResponsePropertiesOutput() PrivateEndpointConnectionResponsePropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionResponsePropertiesOutput) ToPrivateEndpointConnectionResponsePropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponsePropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionResponsePropertiesOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponseProperties) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionResponsePropertiesOutput) PrivateEndpoint() ConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponseProperties) *ConnectionPropertiesResponsePrivateEndpoint {
		return v.PrivateEndpoint
	}).(ConnectionPropertiesResponsePrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionResponsePropertiesOutput) PrivateLinkServiceConnectionState() ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponseProperties) *ConnectionPropertiesResponsePrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateEndpointConnectionResponsePropertiesOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponseProperties) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type ServiceBus struct {
	AuthenticationType        *string `pulumi:"authenticationType"`
	DeadLetterSecret          *string `pulumi:"deadLetterSecret"`
	DeadLetterUri             *string `pulumi:"deadLetterUri"`
	EndpointType              string  `pulumi:"endpointType"`
	EndpointUri               *string `pulumi:"endpointUri"`
	EntityPath                *string `pulumi:"entityPath"`
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
}

type ServiceBusResponse struct {
	AuthenticationType        *string `pulumi:"authenticationType"`
	CreatedTime               string  `pulumi:"createdTime"`
	DeadLetterSecret          *string `pulumi:"deadLetterSecret"`
	DeadLetterUri             *string `pulumi:"deadLetterUri"`
	EndpointType              string  `pulumi:"endpointType"`
	EndpointUri               *string `pulumi:"endpointUri"`
	EntityPath                *string `pulumi:"entityPath"`
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	ProvisioningState         string  `pulumi:"provisioningState"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
}

func init() {
	pulumi.RegisterOutputType(ConnectionPropertiesPrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesResponsePrivateEndpointOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesResponsePrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(DigitalTwinsIdentityOutput{})
	pulumi.RegisterOutputType(DigitalTwinsIdentityPtrOutput{})
	pulumi.RegisterOutputType(DigitalTwinsIdentityResponseOutput{})
	pulumi.RegisterOutputType(DigitalTwinsIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponsePropertiesOutput{})
}

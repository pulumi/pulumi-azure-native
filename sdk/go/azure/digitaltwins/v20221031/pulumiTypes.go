


package v20221031

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureDataExplorerConnectionProperties struct {
	AdxDatabaseName             string                    `pulumi:"adxDatabaseName"`
	AdxEndpointUri              string                    `pulumi:"adxEndpointUri"`
	AdxResourceId               string                    `pulumi:"adxResourceId"`
	AdxTableName                *string                   `pulumi:"adxTableName"`
	ConnectionType              string                    `pulumi:"connectionType"`
	EventHubConsumerGroup       *string                   `pulumi:"eventHubConsumerGroup"`
	EventHubEndpointUri         string                    `pulumi:"eventHubEndpointUri"`
	EventHubEntityPath          string                    `pulumi:"eventHubEntityPath"`
	EventHubNamespaceResourceId string                    `pulumi:"eventHubNamespaceResourceId"`
	Identity                    *ManagedIdentityReference `pulumi:"identity"`
}


func (val *AzureDataExplorerConnectionProperties) Defaults() *AzureDataExplorerConnectionProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AdxTableName) {
		adxTableName_ := "AdtPropertyEvents"
		tmp.AdxTableName = &adxTableName_
	}
	if isZero(tmp.EventHubConsumerGroup) {
		eventHubConsumerGroup_ := "$Default"
		tmp.EventHubConsumerGroup = &eventHubConsumerGroup_
	}
	return &tmp
}





type AzureDataExplorerConnectionPropertiesInput interface {
	pulumi.Input

	ToAzureDataExplorerConnectionPropertiesOutput() AzureDataExplorerConnectionPropertiesOutput
	ToAzureDataExplorerConnectionPropertiesOutputWithContext(context.Context) AzureDataExplorerConnectionPropertiesOutput
}

type AzureDataExplorerConnectionPropertiesArgs struct {
	AdxDatabaseName             pulumi.StringInput               `pulumi:"adxDatabaseName"`
	AdxEndpointUri              pulumi.StringInput               `pulumi:"adxEndpointUri"`
	AdxResourceId               pulumi.StringInput               `pulumi:"adxResourceId"`
	AdxTableName                pulumi.StringPtrInput            `pulumi:"adxTableName"`
	ConnectionType              pulumi.StringInput               `pulumi:"connectionType"`
	EventHubConsumerGroup       pulumi.StringPtrInput            `pulumi:"eventHubConsumerGroup"`
	EventHubEndpointUri         pulumi.StringInput               `pulumi:"eventHubEndpointUri"`
	EventHubEntityPath          pulumi.StringInput               `pulumi:"eventHubEntityPath"`
	EventHubNamespaceResourceId pulumi.StringInput               `pulumi:"eventHubNamespaceResourceId"`
	Identity                    ManagedIdentityReferencePtrInput `pulumi:"identity"`
}


func (val *AzureDataExplorerConnectionPropertiesArgs) Defaults() *AzureDataExplorerConnectionPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AdxTableName) {
		tmp.AdxTableName = pulumi.StringPtr("AdtPropertyEvents")
	}
	if isZero(tmp.EventHubConsumerGroup) {
		tmp.EventHubConsumerGroup = pulumi.StringPtr("$Default")
	}
	return &tmp
}
func (AzureDataExplorerConnectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDataExplorerConnectionProperties)(nil)).Elem()
}

func (i AzureDataExplorerConnectionPropertiesArgs) ToAzureDataExplorerConnectionPropertiesOutput() AzureDataExplorerConnectionPropertiesOutput {
	return i.ToAzureDataExplorerConnectionPropertiesOutputWithContext(context.Background())
}

func (i AzureDataExplorerConnectionPropertiesArgs) ToAzureDataExplorerConnectionPropertiesOutputWithContext(ctx context.Context) AzureDataExplorerConnectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDataExplorerConnectionPropertiesOutput)
}

func (i AzureDataExplorerConnectionPropertiesArgs) ToAzureDataExplorerConnectionPropertiesPtrOutput() AzureDataExplorerConnectionPropertiesPtrOutput {
	return i.ToAzureDataExplorerConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i AzureDataExplorerConnectionPropertiesArgs) ToAzureDataExplorerConnectionPropertiesPtrOutputWithContext(ctx context.Context) AzureDataExplorerConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDataExplorerConnectionPropertiesOutput).ToAzureDataExplorerConnectionPropertiesPtrOutputWithContext(ctx)
}









type AzureDataExplorerConnectionPropertiesPtrInput interface {
	pulumi.Input

	ToAzureDataExplorerConnectionPropertiesPtrOutput() AzureDataExplorerConnectionPropertiesPtrOutput
	ToAzureDataExplorerConnectionPropertiesPtrOutputWithContext(context.Context) AzureDataExplorerConnectionPropertiesPtrOutput
}

type azureDataExplorerConnectionPropertiesPtrType AzureDataExplorerConnectionPropertiesArgs

func AzureDataExplorerConnectionPropertiesPtr(v *AzureDataExplorerConnectionPropertiesArgs) AzureDataExplorerConnectionPropertiesPtrInput {
	return (*azureDataExplorerConnectionPropertiesPtrType)(v)
}

func (*azureDataExplorerConnectionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureDataExplorerConnectionProperties)(nil)).Elem()
}

func (i *azureDataExplorerConnectionPropertiesPtrType) ToAzureDataExplorerConnectionPropertiesPtrOutput() AzureDataExplorerConnectionPropertiesPtrOutput {
	return i.ToAzureDataExplorerConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i *azureDataExplorerConnectionPropertiesPtrType) ToAzureDataExplorerConnectionPropertiesPtrOutputWithContext(ctx context.Context) AzureDataExplorerConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDataExplorerConnectionPropertiesPtrOutput)
}

type AzureDataExplorerConnectionPropertiesOutput struct{ *pulumi.OutputState }

func (AzureDataExplorerConnectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDataExplorerConnectionProperties)(nil)).Elem()
}

func (o AzureDataExplorerConnectionPropertiesOutput) ToAzureDataExplorerConnectionPropertiesOutput() AzureDataExplorerConnectionPropertiesOutput {
	return o
}

func (o AzureDataExplorerConnectionPropertiesOutput) ToAzureDataExplorerConnectionPropertiesOutputWithContext(ctx context.Context) AzureDataExplorerConnectionPropertiesOutput {
	return o
}

func (o AzureDataExplorerConnectionPropertiesOutput) ToAzureDataExplorerConnectionPropertiesPtrOutput() AzureDataExplorerConnectionPropertiesPtrOutput {
	return o.ToAzureDataExplorerConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (o AzureDataExplorerConnectionPropertiesOutput) ToAzureDataExplorerConnectionPropertiesPtrOutputWithContext(ctx context.Context) AzureDataExplorerConnectionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureDataExplorerConnectionProperties) *AzureDataExplorerConnectionProperties {
		return &v
	}).(AzureDataExplorerConnectionPropertiesPtrOutput)
}

func (o AzureDataExplorerConnectionPropertiesOutput) AdxDatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionProperties) string { return v.AdxDatabaseName }).(pulumi.StringOutput)
}

func (o AzureDataExplorerConnectionPropertiesOutput) AdxEndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionProperties) string { return v.AdxEndpointUri }).(pulumi.StringOutput)
}

func (o AzureDataExplorerConnectionPropertiesOutput) AdxResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionProperties) string { return v.AdxResourceId }).(pulumi.StringOutput)
}

func (o AzureDataExplorerConnectionPropertiesOutput) AdxTableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionProperties) *string { return v.AdxTableName }).(pulumi.StringPtrOutput)
}

func (o AzureDataExplorerConnectionPropertiesOutput) ConnectionType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionProperties) string { return v.ConnectionType }).(pulumi.StringOutput)
}

func (o AzureDataExplorerConnectionPropertiesOutput) EventHubConsumerGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionProperties) *string { return v.EventHubConsumerGroup }).(pulumi.StringPtrOutput)
}

func (o AzureDataExplorerConnectionPropertiesOutput) EventHubEndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionProperties) string { return v.EventHubEndpointUri }).(pulumi.StringOutput)
}

func (o AzureDataExplorerConnectionPropertiesOutput) EventHubEntityPath() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionProperties) string { return v.EventHubEntityPath }).(pulumi.StringOutput)
}

func (o AzureDataExplorerConnectionPropertiesOutput) EventHubNamespaceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionProperties) string { return v.EventHubNamespaceResourceId }).(pulumi.StringOutput)
}

func (o AzureDataExplorerConnectionPropertiesOutput) Identity() ManagedIdentityReferencePtrOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionProperties) *ManagedIdentityReference { return v.Identity }).(ManagedIdentityReferencePtrOutput)
}

type AzureDataExplorerConnectionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AzureDataExplorerConnectionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureDataExplorerConnectionProperties)(nil)).Elem()
}

func (o AzureDataExplorerConnectionPropertiesPtrOutput) ToAzureDataExplorerConnectionPropertiesPtrOutput() AzureDataExplorerConnectionPropertiesPtrOutput {
	return o
}

func (o AzureDataExplorerConnectionPropertiesPtrOutput) ToAzureDataExplorerConnectionPropertiesPtrOutputWithContext(ctx context.Context) AzureDataExplorerConnectionPropertiesPtrOutput {
	return o
}

func (o AzureDataExplorerConnectionPropertiesPtrOutput) Elem() AzureDataExplorerConnectionPropertiesOutput {
	return o.ApplyT(func(v *AzureDataExplorerConnectionProperties) AzureDataExplorerConnectionProperties {
		if v != nil {
			return *v
		}
		var ret AzureDataExplorerConnectionProperties
		return ret
	}).(AzureDataExplorerConnectionPropertiesOutput)
}

func (o AzureDataExplorerConnectionPropertiesPtrOutput) AdxDatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataExplorerConnectionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AdxDatabaseName
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataExplorerConnectionPropertiesPtrOutput) AdxEndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataExplorerConnectionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AdxEndpointUri
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataExplorerConnectionPropertiesPtrOutput) AdxResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataExplorerConnectionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AdxResourceId
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataExplorerConnectionPropertiesPtrOutput) AdxTableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataExplorerConnectionProperties) *string {
		if v == nil {
			return nil
		}
		return v.AdxTableName
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataExplorerConnectionPropertiesPtrOutput) ConnectionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataExplorerConnectionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ConnectionType
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataExplorerConnectionPropertiesPtrOutput) EventHubConsumerGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataExplorerConnectionProperties) *string {
		if v == nil {
			return nil
		}
		return v.EventHubConsumerGroup
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataExplorerConnectionPropertiesPtrOutput) EventHubEndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataExplorerConnectionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.EventHubEndpointUri
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataExplorerConnectionPropertiesPtrOutput) EventHubEntityPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataExplorerConnectionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.EventHubEntityPath
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataExplorerConnectionPropertiesPtrOutput) EventHubNamespaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataExplorerConnectionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.EventHubNamespaceResourceId
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataExplorerConnectionPropertiesPtrOutput) Identity() ManagedIdentityReferencePtrOutput {
	return o.ApplyT(func(v *AzureDataExplorerConnectionProperties) *ManagedIdentityReference {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(ManagedIdentityReferencePtrOutput)
}

type AzureDataExplorerConnectionPropertiesResponse struct {
	AdxDatabaseName             string                            `pulumi:"adxDatabaseName"`
	AdxEndpointUri              string                            `pulumi:"adxEndpointUri"`
	AdxResourceId               string                            `pulumi:"adxResourceId"`
	AdxTableName                *string                           `pulumi:"adxTableName"`
	ConnectionType              string                            `pulumi:"connectionType"`
	EventHubConsumerGroup       *string                           `pulumi:"eventHubConsumerGroup"`
	EventHubEndpointUri         string                            `pulumi:"eventHubEndpointUri"`
	EventHubEntityPath          string                            `pulumi:"eventHubEntityPath"`
	EventHubNamespaceResourceId string                            `pulumi:"eventHubNamespaceResourceId"`
	Identity                    *ManagedIdentityReferenceResponse `pulumi:"identity"`
	ProvisioningState           string                            `pulumi:"provisioningState"`
}


func (val *AzureDataExplorerConnectionPropertiesResponse) Defaults() *AzureDataExplorerConnectionPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AdxTableName) {
		adxTableName_ := "AdtPropertyEvents"
		tmp.AdxTableName = &adxTableName_
	}
	if isZero(tmp.EventHubConsumerGroup) {
		eventHubConsumerGroup_ := "$Default"
		tmp.EventHubConsumerGroup = &eventHubConsumerGroup_
	}
	return &tmp
}

type AzureDataExplorerConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AzureDataExplorerConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDataExplorerConnectionPropertiesResponse)(nil)).Elem()
}

func (o AzureDataExplorerConnectionPropertiesResponseOutput) ToAzureDataExplorerConnectionPropertiesResponseOutput() AzureDataExplorerConnectionPropertiesResponseOutput {
	return o
}

func (o AzureDataExplorerConnectionPropertiesResponseOutput) ToAzureDataExplorerConnectionPropertiesResponseOutputWithContext(ctx context.Context) AzureDataExplorerConnectionPropertiesResponseOutput {
	return o
}

func (o AzureDataExplorerConnectionPropertiesResponseOutput) AdxDatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionPropertiesResponse) string { return v.AdxDatabaseName }).(pulumi.StringOutput)
}

func (o AzureDataExplorerConnectionPropertiesResponseOutput) AdxEndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionPropertiesResponse) string { return v.AdxEndpointUri }).(pulumi.StringOutput)
}

func (o AzureDataExplorerConnectionPropertiesResponseOutput) AdxResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionPropertiesResponse) string { return v.AdxResourceId }).(pulumi.StringOutput)
}

func (o AzureDataExplorerConnectionPropertiesResponseOutput) AdxTableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionPropertiesResponse) *string { return v.AdxTableName }).(pulumi.StringPtrOutput)
}

func (o AzureDataExplorerConnectionPropertiesResponseOutput) ConnectionType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionPropertiesResponse) string { return v.ConnectionType }).(pulumi.StringOutput)
}

func (o AzureDataExplorerConnectionPropertiesResponseOutput) EventHubConsumerGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionPropertiesResponse) *string { return v.EventHubConsumerGroup }).(pulumi.StringPtrOutput)
}

func (o AzureDataExplorerConnectionPropertiesResponseOutput) EventHubEndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionPropertiesResponse) string { return v.EventHubEndpointUri }).(pulumi.StringOutput)
}

func (o AzureDataExplorerConnectionPropertiesResponseOutput) EventHubEntityPath() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionPropertiesResponse) string { return v.EventHubEntityPath }).(pulumi.StringOutput)
}

func (o AzureDataExplorerConnectionPropertiesResponseOutput) EventHubNamespaceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionPropertiesResponse) string { return v.EventHubNamespaceResourceId }).(pulumi.StringOutput)
}

func (o AzureDataExplorerConnectionPropertiesResponseOutput) Identity() ManagedIdentityReferenceResponsePtrOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionPropertiesResponse) *ManagedIdentityReferenceResponse {
		return v.Identity
	}).(ManagedIdentityReferenceResponsePtrOutput)
}

func (o AzureDataExplorerConnectionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataExplorerConnectionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type ConnectionProperties struct {
	GroupIds                          []string                                               `pulumi:"groupIds"`
	PrivateLinkServiceConnectionState *ConnectionPropertiesPrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type ConnectionPropertiesInput interface {
	pulumi.Input

	ToConnectionPropertiesOutput() ConnectionPropertiesOutput
	ToConnectionPropertiesOutputWithContext(context.Context) ConnectionPropertiesOutput
}

type ConnectionPropertiesArgs struct {
	GroupIds                          pulumi.StringArrayInput                                       `pulumi:"groupIds"`
	PrivateLinkServiceConnectionState ConnectionPropertiesPrivateLinkServiceConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
}

func (ConnectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionProperties)(nil)).Elem()
}

func (i ConnectionPropertiesArgs) ToConnectionPropertiesOutput() ConnectionPropertiesOutput {
	return i.ToConnectionPropertiesOutputWithContext(context.Background())
}

func (i ConnectionPropertiesArgs) ToConnectionPropertiesOutputWithContext(ctx context.Context) ConnectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesOutput)
}

type ConnectionPropertiesOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionProperties)(nil)).Elem()
}

func (o ConnectionPropertiesOutput) ToConnectionPropertiesOutput() ConnectionPropertiesOutput {
	return o
}

func (o ConnectionPropertiesOutput) ToConnectionPropertiesOutputWithContext(ctx context.Context) ConnectionPropertiesOutput {
	return o
}

func (o ConnectionPropertiesOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConnectionProperties) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o ConnectionPropertiesOutput) PrivateLinkServiceConnectionState() ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v ConnectionProperties) *ConnectionPropertiesPrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
}

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

type ConnectionPropertiesResponse struct {
	GroupIds                          []string                                                       `pulumi:"groupIds"`
	PrivateEndpoint                   *PrivateEndpointResponse                                       `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *ConnectionPropertiesResponsePrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                                         `pulumi:"provisioningState"`
}

type ConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionPropertiesResponse)(nil)).Elem()
}

func (o ConnectionPropertiesResponseOutput) ToConnectionPropertiesResponseOutput() ConnectionPropertiesResponseOutput {
	return o
}

func (o ConnectionPropertiesResponseOutput) ToConnectionPropertiesResponseOutputWithContext(ctx context.Context) ConnectionPropertiesResponseOutput {
	return o
}

func (o ConnectionPropertiesResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConnectionPropertiesResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o ConnectionPropertiesResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v ConnectionPropertiesResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o ConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v ConnectionPropertiesResponse) *ConnectionPropertiesResponsePrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput)
}

func (o ConnectionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
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
	Type                   *string                `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type DigitalTwinsIdentityInput interface {
	pulumi.Input

	ToDigitalTwinsIdentityOutput() DigitalTwinsIdentityOutput
	ToDigitalTwinsIdentityOutputWithContext(context.Context) DigitalTwinsIdentityOutput
}

type DigitalTwinsIdentityArgs struct {
	Type                   pulumi.StringPtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput       `pulumi:"userAssignedIdentities"`
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

func (o DigitalTwinsIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v DigitalTwinsIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
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

func (o DigitalTwinsIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *DigitalTwinsIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type DigitalTwinsIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
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

func (o DigitalTwinsIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v DigitalTwinsIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
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

func (o DigitalTwinsIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *DigitalTwinsIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type EventGrid struct {
	AccessKey1         string                    `pulumi:"accessKey1"`
	AccessKey2         *string                   `pulumi:"accessKey2"`
	AuthenticationType *string                   `pulumi:"authenticationType"`
	DeadLetterSecret   *string                   `pulumi:"deadLetterSecret"`
	DeadLetterUri      *string                   `pulumi:"deadLetterUri"`
	EndpointType       string                    `pulumi:"endpointType"`
	Identity           *ManagedIdentityReference `pulumi:"identity"`
	TopicEndpoint      string                    `pulumi:"topicEndpoint"`
}

type EventGridResponse struct {
	AccessKey1         string                            `pulumi:"accessKey1"`
	AccessKey2         *string                           `pulumi:"accessKey2"`
	AuthenticationType *string                           `pulumi:"authenticationType"`
	CreatedTime        string                            `pulumi:"createdTime"`
	DeadLetterSecret   *string                           `pulumi:"deadLetterSecret"`
	DeadLetterUri      *string                           `pulumi:"deadLetterUri"`
	EndpointType       string                            `pulumi:"endpointType"`
	Identity           *ManagedIdentityReferenceResponse `pulumi:"identity"`
	ProvisioningState  string                            `pulumi:"provisioningState"`
	TopicEndpoint      string                            `pulumi:"topicEndpoint"`
}

type EventHub struct {
	AuthenticationType           *string                   `pulumi:"authenticationType"`
	ConnectionStringPrimaryKey   *string                   `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey *string                   `pulumi:"connectionStringSecondaryKey"`
	DeadLetterSecret             *string                   `pulumi:"deadLetterSecret"`
	DeadLetterUri                *string                   `pulumi:"deadLetterUri"`
	EndpointType                 string                    `pulumi:"endpointType"`
	EndpointUri                  *string                   `pulumi:"endpointUri"`
	EntityPath                   *string                   `pulumi:"entityPath"`
	Identity                     *ManagedIdentityReference `pulumi:"identity"`
}

type EventHubResponse struct {
	AuthenticationType           *string                           `pulumi:"authenticationType"`
	ConnectionStringPrimaryKey   *string                           `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey *string                           `pulumi:"connectionStringSecondaryKey"`
	CreatedTime                  string                            `pulumi:"createdTime"`
	DeadLetterSecret             *string                           `pulumi:"deadLetterSecret"`
	DeadLetterUri                *string                           `pulumi:"deadLetterUri"`
	EndpointType                 string                            `pulumi:"endpointType"`
	EndpointUri                  *string                           `pulumi:"endpointUri"`
	EntityPath                   *string                           `pulumi:"entityPath"`
	Identity                     *ManagedIdentityReferenceResponse `pulumi:"identity"`
	ProvisioningState            string                            `pulumi:"provisioningState"`
}

type ManagedIdentityReference struct {
	Type                 *string `pulumi:"type"`
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}





type ManagedIdentityReferenceInput interface {
	pulumi.Input

	ToManagedIdentityReferenceOutput() ManagedIdentityReferenceOutput
	ToManagedIdentityReferenceOutputWithContext(context.Context) ManagedIdentityReferenceOutput
}

type ManagedIdentityReferenceArgs struct {
	Type                 pulumi.StringPtrInput `pulumi:"type"`
	UserAssignedIdentity pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (ManagedIdentityReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityReference)(nil)).Elem()
}

func (i ManagedIdentityReferenceArgs) ToManagedIdentityReferenceOutput() ManagedIdentityReferenceOutput {
	return i.ToManagedIdentityReferenceOutputWithContext(context.Background())
}

func (i ManagedIdentityReferenceArgs) ToManagedIdentityReferenceOutputWithContext(ctx context.Context) ManagedIdentityReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityReferenceOutput)
}

func (i ManagedIdentityReferenceArgs) ToManagedIdentityReferencePtrOutput() ManagedIdentityReferencePtrOutput {
	return i.ToManagedIdentityReferencePtrOutputWithContext(context.Background())
}

func (i ManagedIdentityReferenceArgs) ToManagedIdentityReferencePtrOutputWithContext(ctx context.Context) ManagedIdentityReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityReferenceOutput).ToManagedIdentityReferencePtrOutputWithContext(ctx)
}









type ManagedIdentityReferencePtrInput interface {
	pulumi.Input

	ToManagedIdentityReferencePtrOutput() ManagedIdentityReferencePtrOutput
	ToManagedIdentityReferencePtrOutputWithContext(context.Context) ManagedIdentityReferencePtrOutput
}

type managedIdentityReferencePtrType ManagedIdentityReferenceArgs

func ManagedIdentityReferencePtr(v *ManagedIdentityReferenceArgs) ManagedIdentityReferencePtrInput {
	return (*managedIdentityReferencePtrType)(v)
}

func (*managedIdentityReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityReference)(nil)).Elem()
}

func (i *managedIdentityReferencePtrType) ToManagedIdentityReferencePtrOutput() ManagedIdentityReferencePtrOutput {
	return i.ToManagedIdentityReferencePtrOutputWithContext(context.Background())
}

func (i *managedIdentityReferencePtrType) ToManagedIdentityReferencePtrOutputWithContext(ctx context.Context) ManagedIdentityReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityReferencePtrOutput)
}

type ManagedIdentityReferenceOutput struct{ *pulumi.OutputState }

func (ManagedIdentityReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityReference)(nil)).Elem()
}

func (o ManagedIdentityReferenceOutput) ToManagedIdentityReferenceOutput() ManagedIdentityReferenceOutput {
	return o
}

func (o ManagedIdentityReferenceOutput) ToManagedIdentityReferenceOutputWithContext(ctx context.Context) ManagedIdentityReferenceOutput {
	return o
}

func (o ManagedIdentityReferenceOutput) ToManagedIdentityReferencePtrOutput() ManagedIdentityReferencePtrOutput {
	return o.ToManagedIdentityReferencePtrOutputWithContext(context.Background())
}

func (o ManagedIdentityReferenceOutput) ToManagedIdentityReferencePtrOutputWithContext(ctx context.Context) ManagedIdentityReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentityReference) *ManagedIdentityReference {
		return &v
	}).(ManagedIdentityReferencePtrOutput)
}

func (o ManagedIdentityReferenceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityReference) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityReferenceOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityReference) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type ManagedIdentityReferencePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityReference)(nil)).Elem()
}

func (o ManagedIdentityReferencePtrOutput) ToManagedIdentityReferencePtrOutput() ManagedIdentityReferencePtrOutput {
	return o
}

func (o ManagedIdentityReferencePtrOutput) ToManagedIdentityReferencePtrOutputWithContext(ctx context.Context) ManagedIdentityReferencePtrOutput {
	return o
}

func (o ManagedIdentityReferencePtrOutput) Elem() ManagedIdentityReferenceOutput {
	return o.ApplyT(func(v *ManagedIdentityReference) ManagedIdentityReference {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityReference
		return ret
	}).(ManagedIdentityReferenceOutput)
}

func (o ManagedIdentityReferencePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityReference) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityReferencePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityReference) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type ManagedIdentityReferenceResponse struct {
	Type                 *string `pulumi:"type"`
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}

type ManagedIdentityReferenceResponseOutput struct{ *pulumi.OutputState }

func (ManagedIdentityReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityReferenceResponse)(nil)).Elem()
}

func (o ManagedIdentityReferenceResponseOutput) ToManagedIdentityReferenceResponseOutput() ManagedIdentityReferenceResponseOutput {
	return o
}

func (o ManagedIdentityReferenceResponseOutput) ToManagedIdentityReferenceResponseOutputWithContext(ctx context.Context) ManagedIdentityReferenceResponseOutput {
	return o
}

func (o ManagedIdentityReferenceResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityReferenceResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityReferenceResponseOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityReferenceResponse) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type ManagedIdentityReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityReferenceResponse)(nil)).Elem()
}

func (o ManagedIdentityReferenceResponsePtrOutput) ToManagedIdentityReferenceResponsePtrOutput() ManagedIdentityReferenceResponsePtrOutput {
	return o
}

func (o ManagedIdentityReferenceResponsePtrOutput) ToManagedIdentityReferenceResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityReferenceResponsePtrOutput {
	return o
}

func (o ManagedIdentityReferenceResponsePtrOutput) Elem() ManagedIdentityReferenceResponseOutput {
	return o.ApplyT(func(v *ManagedIdentityReferenceResponse) ManagedIdentityReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityReferenceResponse
		return ret
	}).(ManagedIdentityReferenceResponseOutput)
}

func (o ManagedIdentityReferenceResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityReferenceResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionType struct {
	Properties ConnectionProperties `pulumi:"properties"`
}





type PrivateEndpointConnectionTypeInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput
	ToPrivateEndpointConnectionTypeOutputWithContext(context.Context) PrivateEndpointConnectionTypeOutput
}

type PrivateEndpointConnectionTypeArgs struct {
	Properties ConnectionPropertiesInput `pulumi:"properties"`
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

func (o PrivateEndpointConnectionTypeOutput) Properties() ConnectionPropertiesOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) ConnectionProperties { return v.Properties }).(ConnectionPropertiesOutput)
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

type PrivateEndpointConnectionResponse struct {
	Id         string                       `pulumi:"id"`
	Name       string                       `pulumi:"name"`
	Properties ConnectionPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse           `pulumi:"systemData"`
	Type       string                       `pulumi:"type"`
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

func (o PrivateEndpointConnectionResponseOutput) Properties() ConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) ConnectionPropertiesResponse { return v.Properties }).(ConnectionPropertiesResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
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
	Id string `pulumi:"id"`
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

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
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
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ServiceBus struct {
	AuthenticationType        *string                   `pulumi:"authenticationType"`
	DeadLetterSecret          *string                   `pulumi:"deadLetterSecret"`
	DeadLetterUri             *string                   `pulumi:"deadLetterUri"`
	EndpointType              string                    `pulumi:"endpointType"`
	EndpointUri               *string                   `pulumi:"endpointUri"`
	EntityPath                *string                   `pulumi:"entityPath"`
	Identity                  *ManagedIdentityReference `pulumi:"identity"`
	PrimaryConnectionString   *string                   `pulumi:"primaryConnectionString"`
	SecondaryConnectionString *string                   `pulumi:"secondaryConnectionString"`
}

type ServiceBusResponse struct {
	AuthenticationType        *string                           `pulumi:"authenticationType"`
	CreatedTime               string                            `pulumi:"createdTime"`
	DeadLetterSecret          *string                           `pulumi:"deadLetterSecret"`
	DeadLetterUri             *string                           `pulumi:"deadLetterUri"`
	EndpointType              string                            `pulumi:"endpointType"`
	EndpointUri               *string                           `pulumi:"endpointUri"`
	EntityPath                *string                           `pulumi:"entityPath"`
	Identity                  *ManagedIdentityReferenceResponse `pulumi:"identity"`
	PrimaryConnectionString   *string                           `pulumi:"primaryConnectionString"`
	ProvisioningState         string                            `pulumi:"provisioningState"`
	SecondaryConnectionString *string                           `pulumi:"secondaryConnectionString"`
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

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureDataExplorerConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(AzureDataExplorerConnectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AzureDataExplorerConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesPrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(DigitalTwinsIdentityOutput{})
	pulumi.RegisterOutputType(DigitalTwinsIdentityPtrOutput{})
	pulumi.RegisterOutputType(DigitalTwinsIdentityResponseOutput{})
	pulumi.RegisterOutputType(DigitalTwinsIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityReferenceOutput{})
	pulumi.RegisterOutputType(ManagedIdentityReferencePtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityReferenceResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentityReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}

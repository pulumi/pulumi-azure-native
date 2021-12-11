


package v20211015

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CertificatePropertiesResponse struct {
	Certificate string `pulumi:"certificate"`
	Created     string `pulumi:"created"`
	Expiry      string `pulumi:"expiry"`
	IsVerified  bool   `pulumi:"isVerified"`
	Subject     string `pulumi:"subject"`
	Thumbprint  string `pulumi:"thumbprint"`
	Updated     string `pulumi:"updated"`
}





type CertificatePropertiesResponseInput interface {
	pulumi.Input

	ToCertificatePropertiesResponseOutput() CertificatePropertiesResponseOutput
	ToCertificatePropertiesResponseOutputWithContext(context.Context) CertificatePropertiesResponseOutput
}

type CertificatePropertiesResponseArgs struct {
	Certificate pulumi.StringInput `pulumi:"certificate"`
	Created     pulumi.StringInput `pulumi:"created"`
	Expiry      pulumi.StringInput `pulumi:"expiry"`
	IsVerified  pulumi.BoolInput   `pulumi:"isVerified"`
	Subject     pulumi.StringInput `pulumi:"subject"`
	Thumbprint  pulumi.StringInput `pulumi:"thumbprint"`
	Updated     pulumi.StringInput `pulumi:"updated"`
}

func (CertificatePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificatePropertiesResponse)(nil)).Elem()
}

func (i CertificatePropertiesResponseArgs) ToCertificatePropertiesResponseOutput() CertificatePropertiesResponseOutput {
	return i.ToCertificatePropertiesResponseOutputWithContext(context.Background())
}

func (i CertificatePropertiesResponseArgs) ToCertificatePropertiesResponseOutputWithContext(ctx context.Context) CertificatePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesResponseOutput)
}

func (i CertificatePropertiesResponseArgs) ToCertificatePropertiesResponsePtrOutput() CertificatePropertiesResponsePtrOutput {
	return i.ToCertificatePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CertificatePropertiesResponseArgs) ToCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) CertificatePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesResponseOutput).ToCertificatePropertiesResponsePtrOutputWithContext(ctx)
}









type CertificatePropertiesResponsePtrInput interface {
	pulumi.Input

	ToCertificatePropertiesResponsePtrOutput() CertificatePropertiesResponsePtrOutput
	ToCertificatePropertiesResponsePtrOutputWithContext(context.Context) CertificatePropertiesResponsePtrOutput
}

type certificatePropertiesResponsePtrType CertificatePropertiesResponseArgs

func CertificatePropertiesResponsePtr(v *CertificatePropertiesResponseArgs) CertificatePropertiesResponsePtrInput {
	return (*certificatePropertiesResponsePtrType)(v)
}

func (*certificatePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificatePropertiesResponse)(nil)).Elem()
}

func (i *certificatePropertiesResponsePtrType) ToCertificatePropertiesResponsePtrOutput() CertificatePropertiesResponsePtrOutput {
	return i.ToCertificatePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *certificatePropertiesResponsePtrType) ToCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) CertificatePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesResponsePtrOutput)
}

type CertificatePropertiesResponseOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificatePropertiesResponse)(nil)).Elem()
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponseOutput() CertificatePropertiesResponseOutput {
	return o
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponseOutputWithContext(ctx context.Context) CertificatePropertiesResponseOutput {
	return o
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponsePtrOutput() CertificatePropertiesResponsePtrOutput {
	return o.ToCertificatePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) CertificatePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificatePropertiesResponse) *CertificatePropertiesResponse {
		return &v
	}).(CertificatePropertiesResponsePtrOutput)
}

func (o CertificatePropertiesResponseOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Certificate }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Created }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Expiry() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Expiry }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) IsVerified() pulumi.BoolOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) bool { return v.IsVerified }).(pulumi.BoolOutput)
}

func (o CertificatePropertiesResponseOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Subject }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Updated }).(pulumi.StringOutput)
}

type CertificatePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificatePropertiesResponse)(nil)).Elem()
}

func (o CertificatePropertiesResponsePtrOutput) ToCertificatePropertiesResponsePtrOutput() CertificatePropertiesResponsePtrOutput {
	return o
}

func (o CertificatePropertiesResponsePtrOutput) ToCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) CertificatePropertiesResponsePtrOutput {
	return o
}

func (o CertificatePropertiesResponsePtrOutput) Elem() CertificatePropertiesResponseOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) CertificatePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CertificatePropertiesResponse
		return ret
	}).(CertificatePropertiesResponseOutput)
}

func (o CertificatePropertiesResponsePtrOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Certificate
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Created
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Expiry
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) IsVerified() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsVerified
	}).(pulumi.BoolPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) Subject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Subject
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) Updated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Updated
	}).(pulumi.StringPtrOutput)
}

type IotDpsPropertiesDescription struct {
	AllocationPolicy           *string                                                         `pulumi:"allocationPolicy"`
	AuthorizationPolicies      []SharedAccessSignatureAuthorizationRuleAccessRightsDescription `pulumi:"authorizationPolicies"`
	EnableDataResidency        *bool                                                           `pulumi:"enableDataResidency"`
	IotHubs                    []IotHubDefinitionDescription                                   `pulumi:"iotHubs"`
	IpFilterRules              []TargetIpFilterRule                                            `pulumi:"ipFilterRules"`
	PrivateEndpointConnections []PrivateEndpointConnection                                     `pulumi:"privateEndpointConnections"`
	ProvisioningState          *string                                                         `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                                                         `pulumi:"publicNetworkAccess"`
	State                      *string                                                         `pulumi:"state"`
}





type IotDpsPropertiesDescriptionInput interface {
	pulumi.Input

	ToIotDpsPropertiesDescriptionOutput() IotDpsPropertiesDescriptionOutput
	ToIotDpsPropertiesDescriptionOutputWithContext(context.Context) IotDpsPropertiesDescriptionOutput
}

type IotDpsPropertiesDescriptionArgs struct {
	AllocationPolicy           pulumi.StringPtrInput                                                   `pulumi:"allocationPolicy"`
	AuthorizationPolicies      SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayInput `pulumi:"authorizationPolicies"`
	EnableDataResidency        pulumi.BoolPtrInput                                                     `pulumi:"enableDataResidency"`
	IotHubs                    IotHubDefinitionDescriptionArrayInput                                   `pulumi:"iotHubs"`
	IpFilterRules              TargetIpFilterRuleArrayInput                                            `pulumi:"ipFilterRules"`
	PrivateEndpointConnections PrivateEndpointConnectionArrayInput                                     `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringPtrInput                                                   `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrInput                                                   `pulumi:"publicNetworkAccess"`
	State                      pulumi.StringPtrInput                                                   `pulumi:"state"`
}

func (IotDpsPropertiesDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsPropertiesDescription)(nil)).Elem()
}

func (i IotDpsPropertiesDescriptionArgs) ToIotDpsPropertiesDescriptionOutput() IotDpsPropertiesDescriptionOutput {
	return i.ToIotDpsPropertiesDescriptionOutputWithContext(context.Background())
}

func (i IotDpsPropertiesDescriptionArgs) ToIotDpsPropertiesDescriptionOutputWithContext(ctx context.Context) IotDpsPropertiesDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsPropertiesDescriptionOutput)
}

func (i IotDpsPropertiesDescriptionArgs) ToIotDpsPropertiesDescriptionPtrOutput() IotDpsPropertiesDescriptionPtrOutput {
	return i.ToIotDpsPropertiesDescriptionPtrOutputWithContext(context.Background())
}

func (i IotDpsPropertiesDescriptionArgs) ToIotDpsPropertiesDescriptionPtrOutputWithContext(ctx context.Context) IotDpsPropertiesDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsPropertiesDescriptionOutput).ToIotDpsPropertiesDescriptionPtrOutputWithContext(ctx)
}









type IotDpsPropertiesDescriptionPtrInput interface {
	pulumi.Input

	ToIotDpsPropertiesDescriptionPtrOutput() IotDpsPropertiesDescriptionPtrOutput
	ToIotDpsPropertiesDescriptionPtrOutputWithContext(context.Context) IotDpsPropertiesDescriptionPtrOutput
}

type iotDpsPropertiesDescriptionPtrType IotDpsPropertiesDescriptionArgs

func IotDpsPropertiesDescriptionPtr(v *IotDpsPropertiesDescriptionArgs) IotDpsPropertiesDescriptionPtrInput {
	return (*iotDpsPropertiesDescriptionPtrType)(v)
}

func (*iotDpsPropertiesDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IotDpsPropertiesDescription)(nil)).Elem()
}

func (i *iotDpsPropertiesDescriptionPtrType) ToIotDpsPropertiesDescriptionPtrOutput() IotDpsPropertiesDescriptionPtrOutput {
	return i.ToIotDpsPropertiesDescriptionPtrOutputWithContext(context.Background())
}

func (i *iotDpsPropertiesDescriptionPtrType) ToIotDpsPropertiesDescriptionPtrOutputWithContext(ctx context.Context) IotDpsPropertiesDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsPropertiesDescriptionPtrOutput)
}

type IotDpsPropertiesDescriptionOutput struct{ *pulumi.OutputState }

func (IotDpsPropertiesDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsPropertiesDescription)(nil)).Elem()
}

func (o IotDpsPropertiesDescriptionOutput) ToIotDpsPropertiesDescriptionOutput() IotDpsPropertiesDescriptionOutput {
	return o
}

func (o IotDpsPropertiesDescriptionOutput) ToIotDpsPropertiesDescriptionOutputWithContext(ctx context.Context) IotDpsPropertiesDescriptionOutput {
	return o
}

func (o IotDpsPropertiesDescriptionOutput) ToIotDpsPropertiesDescriptionPtrOutput() IotDpsPropertiesDescriptionPtrOutput {
	return o.ToIotDpsPropertiesDescriptionPtrOutputWithContext(context.Background())
}

func (o IotDpsPropertiesDescriptionOutput) ToIotDpsPropertiesDescriptionPtrOutputWithContext(ctx context.Context) IotDpsPropertiesDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IotDpsPropertiesDescription) *IotDpsPropertiesDescription {
		return &v
	}).(IotDpsPropertiesDescriptionPtrOutput)
}

func (o IotDpsPropertiesDescriptionOutput) AllocationPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) *string { return v.AllocationPolicy }).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionOutput) AuthorizationPolicies() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) []SharedAccessSignatureAuthorizationRuleAccessRightsDescription {
		return v.AuthorizationPolicies
	}).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput)
}

func (o IotDpsPropertiesDescriptionOutput) EnableDataResidency() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) *bool { return v.EnableDataResidency }).(pulumi.BoolPtrOutput)
}

func (o IotDpsPropertiesDescriptionOutput) IotHubs() IotHubDefinitionDescriptionArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) []IotHubDefinitionDescription { return v.IotHubs }).(IotHubDefinitionDescriptionArrayOutput)
}

func (o IotDpsPropertiesDescriptionOutput) IpFilterRules() TargetIpFilterRuleArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) []TargetIpFilterRule { return v.IpFilterRules }).(TargetIpFilterRuleArrayOutput)
}

func (o IotDpsPropertiesDescriptionOutput) PrivateEndpointConnections() PrivateEndpointConnectionArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) []PrivateEndpointConnection { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionArrayOutput)
}

func (o IotDpsPropertiesDescriptionOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) *string { return v.State }).(pulumi.StringPtrOutput)
}

type IotDpsPropertiesDescriptionPtrOutput struct{ *pulumi.OutputState }

func (IotDpsPropertiesDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotDpsPropertiesDescription)(nil)).Elem()
}

func (o IotDpsPropertiesDescriptionPtrOutput) ToIotDpsPropertiesDescriptionPtrOutput() IotDpsPropertiesDescriptionPtrOutput {
	return o
}

func (o IotDpsPropertiesDescriptionPtrOutput) ToIotDpsPropertiesDescriptionPtrOutputWithContext(ctx context.Context) IotDpsPropertiesDescriptionPtrOutput {
	return o
}

func (o IotDpsPropertiesDescriptionPtrOutput) Elem() IotDpsPropertiesDescriptionOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescription) IotDpsPropertiesDescription {
		if v != nil {
			return *v
		}
		var ret IotDpsPropertiesDescription
		return ret
	}).(IotDpsPropertiesDescriptionOutput)
}

func (o IotDpsPropertiesDescriptionPtrOutput) AllocationPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescription) *string {
		if v == nil {
			return nil
		}
		return v.AllocationPolicy
	}).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionPtrOutput) AuthorizationPolicies() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescription) []SharedAccessSignatureAuthorizationRuleAccessRightsDescription {
		if v == nil {
			return nil
		}
		return v.AuthorizationPolicies
	}).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput)
}

func (o IotDpsPropertiesDescriptionPtrOutput) EnableDataResidency() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescription) *bool {
		if v == nil {
			return nil
		}
		return v.EnableDataResidency
	}).(pulumi.BoolPtrOutput)
}

func (o IotDpsPropertiesDescriptionPtrOutput) IotHubs() IotHubDefinitionDescriptionArrayOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescription) []IotHubDefinitionDescription {
		if v == nil {
			return nil
		}
		return v.IotHubs
	}).(IotHubDefinitionDescriptionArrayOutput)
}

func (o IotDpsPropertiesDescriptionPtrOutput) IpFilterRules() TargetIpFilterRuleArrayOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescription) []TargetIpFilterRule {
		if v == nil {
			return nil
		}
		return v.IpFilterRules
	}).(TargetIpFilterRuleArrayOutput)
}

func (o IotDpsPropertiesDescriptionPtrOutput) PrivateEndpointConnections() PrivateEndpointConnectionArrayOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescription) []PrivateEndpointConnection {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionArrayOutput)
}

func (o IotDpsPropertiesDescriptionPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescription) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescription) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescription) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type IotDpsPropertiesDescriptionResponse struct {
	AllocationPolicy           *string                                                                 `pulumi:"allocationPolicy"`
	AuthorizationPolicies      []SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse `pulumi:"authorizationPolicies"`
	DeviceProvisioningHostName string                                                                  `pulumi:"deviceProvisioningHostName"`
	EnableDataResidency        *bool                                                                   `pulumi:"enableDataResidency"`
	IdScope                    string                                                                  `pulumi:"idScope"`
	IotHubs                    []IotHubDefinitionDescriptionResponse                                   `pulumi:"iotHubs"`
	IpFilterRules              []TargetIpFilterRuleResponse                                            `pulumi:"ipFilterRules"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse                                     `pulumi:"privateEndpointConnections"`
	ProvisioningState          *string                                                                 `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                                                                 `pulumi:"publicNetworkAccess"`
	ServiceOperationsHostName  string                                                                  `pulumi:"serviceOperationsHostName"`
	State                      *string                                                                 `pulumi:"state"`
}





type IotDpsPropertiesDescriptionResponseInput interface {
	pulumi.Input

	ToIotDpsPropertiesDescriptionResponseOutput() IotDpsPropertiesDescriptionResponseOutput
	ToIotDpsPropertiesDescriptionResponseOutputWithContext(context.Context) IotDpsPropertiesDescriptionResponseOutput
}

type IotDpsPropertiesDescriptionResponseArgs struct {
	AllocationPolicy           pulumi.StringPtrInput                                                           `pulumi:"allocationPolicy"`
	AuthorizationPolicies      SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayInput `pulumi:"authorizationPolicies"`
	DeviceProvisioningHostName pulumi.StringInput                                                              `pulumi:"deviceProvisioningHostName"`
	EnableDataResidency        pulumi.BoolPtrInput                                                             `pulumi:"enableDataResidency"`
	IdScope                    pulumi.StringInput                                                              `pulumi:"idScope"`
	IotHubs                    IotHubDefinitionDescriptionResponseArrayInput                                   `pulumi:"iotHubs"`
	IpFilterRules              TargetIpFilterRuleResponseArrayInput                                            `pulumi:"ipFilterRules"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayInput                                     `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringPtrInput                                                           `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrInput                                                           `pulumi:"publicNetworkAccess"`
	ServiceOperationsHostName  pulumi.StringInput                                                              `pulumi:"serviceOperationsHostName"`
	State                      pulumi.StringPtrInput                                                           `pulumi:"state"`
}

func (IotDpsPropertiesDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsPropertiesDescriptionResponse)(nil)).Elem()
}

func (i IotDpsPropertiesDescriptionResponseArgs) ToIotDpsPropertiesDescriptionResponseOutput() IotDpsPropertiesDescriptionResponseOutput {
	return i.ToIotDpsPropertiesDescriptionResponseOutputWithContext(context.Background())
}

func (i IotDpsPropertiesDescriptionResponseArgs) ToIotDpsPropertiesDescriptionResponseOutputWithContext(ctx context.Context) IotDpsPropertiesDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsPropertiesDescriptionResponseOutput)
}

func (i IotDpsPropertiesDescriptionResponseArgs) ToIotDpsPropertiesDescriptionResponsePtrOutput() IotDpsPropertiesDescriptionResponsePtrOutput {
	return i.ToIotDpsPropertiesDescriptionResponsePtrOutputWithContext(context.Background())
}

func (i IotDpsPropertiesDescriptionResponseArgs) ToIotDpsPropertiesDescriptionResponsePtrOutputWithContext(ctx context.Context) IotDpsPropertiesDescriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsPropertiesDescriptionResponseOutput).ToIotDpsPropertiesDescriptionResponsePtrOutputWithContext(ctx)
}









type IotDpsPropertiesDescriptionResponsePtrInput interface {
	pulumi.Input

	ToIotDpsPropertiesDescriptionResponsePtrOutput() IotDpsPropertiesDescriptionResponsePtrOutput
	ToIotDpsPropertiesDescriptionResponsePtrOutputWithContext(context.Context) IotDpsPropertiesDescriptionResponsePtrOutput
}

type iotDpsPropertiesDescriptionResponsePtrType IotDpsPropertiesDescriptionResponseArgs

func IotDpsPropertiesDescriptionResponsePtr(v *IotDpsPropertiesDescriptionResponseArgs) IotDpsPropertiesDescriptionResponsePtrInput {
	return (*iotDpsPropertiesDescriptionResponsePtrType)(v)
}

func (*iotDpsPropertiesDescriptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IotDpsPropertiesDescriptionResponse)(nil)).Elem()
}

func (i *iotDpsPropertiesDescriptionResponsePtrType) ToIotDpsPropertiesDescriptionResponsePtrOutput() IotDpsPropertiesDescriptionResponsePtrOutput {
	return i.ToIotDpsPropertiesDescriptionResponsePtrOutputWithContext(context.Background())
}

func (i *iotDpsPropertiesDescriptionResponsePtrType) ToIotDpsPropertiesDescriptionResponsePtrOutputWithContext(ctx context.Context) IotDpsPropertiesDescriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsPropertiesDescriptionResponsePtrOutput)
}

type IotDpsPropertiesDescriptionResponseOutput struct{ *pulumi.OutputState }

func (IotDpsPropertiesDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsPropertiesDescriptionResponse)(nil)).Elem()
}

func (o IotDpsPropertiesDescriptionResponseOutput) ToIotDpsPropertiesDescriptionResponseOutput() IotDpsPropertiesDescriptionResponseOutput {
	return o
}

func (o IotDpsPropertiesDescriptionResponseOutput) ToIotDpsPropertiesDescriptionResponseOutputWithContext(ctx context.Context) IotDpsPropertiesDescriptionResponseOutput {
	return o
}

func (o IotDpsPropertiesDescriptionResponseOutput) ToIotDpsPropertiesDescriptionResponsePtrOutput() IotDpsPropertiesDescriptionResponsePtrOutput {
	return o.ToIotDpsPropertiesDescriptionResponsePtrOutputWithContext(context.Background())
}

func (o IotDpsPropertiesDescriptionResponseOutput) ToIotDpsPropertiesDescriptionResponsePtrOutputWithContext(ctx context.Context) IotDpsPropertiesDescriptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IotDpsPropertiesDescriptionResponse) *IotDpsPropertiesDescriptionResponse {
		return &v
	}).(IotDpsPropertiesDescriptionResponsePtrOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) AllocationPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) *string { return v.AllocationPolicy }).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) AuthorizationPolicies() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) []SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse {
		return v.AuthorizationPolicies
	}).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) DeviceProvisioningHostName() pulumi.StringOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) string { return v.DeviceProvisioningHostName }).(pulumi.StringOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) EnableDataResidency() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) *bool { return v.EnableDataResidency }).(pulumi.BoolPtrOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) IdScope() pulumi.StringOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) string { return v.IdScope }).(pulumi.StringOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) IotHubs() IotHubDefinitionDescriptionResponseArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) []IotHubDefinitionDescriptionResponse { return v.IotHubs }).(IotHubDefinitionDescriptionResponseArrayOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) IpFilterRules() TargetIpFilterRuleResponseArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) []TargetIpFilterRuleResponse { return v.IpFilterRules }).(TargetIpFilterRuleResponseArrayOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) ServiceOperationsHostName() pulumi.StringOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) string { return v.ServiceOperationsHostName }).(pulumi.StringOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type IotDpsPropertiesDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (IotDpsPropertiesDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotDpsPropertiesDescriptionResponse)(nil)).Elem()
}

func (o IotDpsPropertiesDescriptionResponsePtrOutput) ToIotDpsPropertiesDescriptionResponsePtrOutput() IotDpsPropertiesDescriptionResponsePtrOutput {
	return o
}

func (o IotDpsPropertiesDescriptionResponsePtrOutput) ToIotDpsPropertiesDescriptionResponsePtrOutputWithContext(ctx context.Context) IotDpsPropertiesDescriptionResponsePtrOutput {
	return o
}

func (o IotDpsPropertiesDescriptionResponsePtrOutput) Elem() IotDpsPropertiesDescriptionResponseOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescriptionResponse) IotDpsPropertiesDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret IotDpsPropertiesDescriptionResponse
		return ret
	}).(IotDpsPropertiesDescriptionResponseOutput)
}

func (o IotDpsPropertiesDescriptionResponsePtrOutput) AllocationPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.AllocationPolicy
	}).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionResponsePtrOutput) AuthorizationPolicies() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescriptionResponse) []SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse {
		if v == nil {
			return nil
		}
		return v.AuthorizationPolicies
	}).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput)
}

func (o IotDpsPropertiesDescriptionResponsePtrOutput) DeviceProvisioningHostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DeviceProvisioningHostName
	}).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionResponsePtrOutput) EnableDataResidency() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescriptionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableDataResidency
	}).(pulumi.BoolPtrOutput)
}

func (o IotDpsPropertiesDescriptionResponsePtrOutput) IdScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.IdScope
	}).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionResponsePtrOutput) IotHubs() IotHubDefinitionDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescriptionResponse) []IotHubDefinitionDescriptionResponse {
		if v == nil {
			return nil
		}
		return v.IotHubs
	}).(IotHubDefinitionDescriptionResponseArrayOutput)
}

func (o IotDpsPropertiesDescriptionResponsePtrOutput) IpFilterRules() TargetIpFilterRuleResponseArrayOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescriptionResponse) []TargetIpFilterRuleResponse {
		if v == nil {
			return nil
		}
		return v.IpFilterRules
	}).(TargetIpFilterRuleResponseArrayOutput)
}

func (o IotDpsPropertiesDescriptionResponsePtrOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescriptionResponse) []PrivateEndpointConnectionResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o IotDpsPropertiesDescriptionResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionResponsePtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionResponsePtrOutput) ServiceOperationsHostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ServiceOperationsHostName
	}).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type IotDpsSkuInfo struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
}





type IotDpsSkuInfoInput interface {
	pulumi.Input

	ToIotDpsSkuInfoOutput() IotDpsSkuInfoOutput
	ToIotDpsSkuInfoOutputWithContext(context.Context) IotDpsSkuInfoOutput
}

type IotDpsSkuInfoArgs struct {
	Capacity pulumi.Float64PtrInput `pulumi:"capacity"`
	Name     pulumi.StringPtrInput  `pulumi:"name"`
}

func (IotDpsSkuInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsSkuInfo)(nil)).Elem()
}

func (i IotDpsSkuInfoArgs) ToIotDpsSkuInfoOutput() IotDpsSkuInfoOutput {
	return i.ToIotDpsSkuInfoOutputWithContext(context.Background())
}

func (i IotDpsSkuInfoArgs) ToIotDpsSkuInfoOutputWithContext(ctx context.Context) IotDpsSkuInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsSkuInfoOutput)
}

func (i IotDpsSkuInfoArgs) ToIotDpsSkuInfoPtrOutput() IotDpsSkuInfoPtrOutput {
	return i.ToIotDpsSkuInfoPtrOutputWithContext(context.Background())
}

func (i IotDpsSkuInfoArgs) ToIotDpsSkuInfoPtrOutputWithContext(ctx context.Context) IotDpsSkuInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsSkuInfoOutput).ToIotDpsSkuInfoPtrOutputWithContext(ctx)
}









type IotDpsSkuInfoPtrInput interface {
	pulumi.Input

	ToIotDpsSkuInfoPtrOutput() IotDpsSkuInfoPtrOutput
	ToIotDpsSkuInfoPtrOutputWithContext(context.Context) IotDpsSkuInfoPtrOutput
}

type iotDpsSkuInfoPtrType IotDpsSkuInfoArgs

func IotDpsSkuInfoPtr(v *IotDpsSkuInfoArgs) IotDpsSkuInfoPtrInput {
	return (*iotDpsSkuInfoPtrType)(v)
}

func (*iotDpsSkuInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IotDpsSkuInfo)(nil)).Elem()
}

func (i *iotDpsSkuInfoPtrType) ToIotDpsSkuInfoPtrOutput() IotDpsSkuInfoPtrOutput {
	return i.ToIotDpsSkuInfoPtrOutputWithContext(context.Background())
}

func (i *iotDpsSkuInfoPtrType) ToIotDpsSkuInfoPtrOutputWithContext(ctx context.Context) IotDpsSkuInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsSkuInfoPtrOutput)
}

type IotDpsSkuInfoOutput struct{ *pulumi.OutputState }

func (IotDpsSkuInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsSkuInfo)(nil)).Elem()
}

func (o IotDpsSkuInfoOutput) ToIotDpsSkuInfoOutput() IotDpsSkuInfoOutput {
	return o
}

func (o IotDpsSkuInfoOutput) ToIotDpsSkuInfoOutputWithContext(ctx context.Context) IotDpsSkuInfoOutput {
	return o
}

func (o IotDpsSkuInfoOutput) ToIotDpsSkuInfoPtrOutput() IotDpsSkuInfoPtrOutput {
	return o.ToIotDpsSkuInfoPtrOutputWithContext(context.Background())
}

func (o IotDpsSkuInfoOutput) ToIotDpsSkuInfoPtrOutputWithContext(ctx context.Context) IotDpsSkuInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IotDpsSkuInfo) *IotDpsSkuInfo {
		return &v
	}).(IotDpsSkuInfoPtrOutput)
}

func (o IotDpsSkuInfoOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v IotDpsSkuInfo) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o IotDpsSkuInfoOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsSkuInfo) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type IotDpsSkuInfoPtrOutput struct{ *pulumi.OutputState }

func (IotDpsSkuInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotDpsSkuInfo)(nil)).Elem()
}

func (o IotDpsSkuInfoPtrOutput) ToIotDpsSkuInfoPtrOutput() IotDpsSkuInfoPtrOutput {
	return o
}

func (o IotDpsSkuInfoPtrOutput) ToIotDpsSkuInfoPtrOutputWithContext(ctx context.Context) IotDpsSkuInfoPtrOutput {
	return o
}

func (o IotDpsSkuInfoPtrOutput) Elem() IotDpsSkuInfoOutput {
	return o.ApplyT(func(v *IotDpsSkuInfo) IotDpsSkuInfo {
		if v != nil {
			return *v
		}
		var ret IotDpsSkuInfo
		return ret
	}).(IotDpsSkuInfoOutput)
}

func (o IotDpsSkuInfoPtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *IotDpsSkuInfo) *float64 {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o IotDpsSkuInfoPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsSkuInfo) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type IotDpsSkuInfoResponse struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
	Tier     string   `pulumi:"tier"`
}





type IotDpsSkuInfoResponseInput interface {
	pulumi.Input

	ToIotDpsSkuInfoResponseOutput() IotDpsSkuInfoResponseOutput
	ToIotDpsSkuInfoResponseOutputWithContext(context.Context) IotDpsSkuInfoResponseOutput
}

type IotDpsSkuInfoResponseArgs struct {
	Capacity pulumi.Float64PtrInput `pulumi:"capacity"`
	Name     pulumi.StringPtrInput  `pulumi:"name"`
	Tier     pulumi.StringInput     `pulumi:"tier"`
}

func (IotDpsSkuInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsSkuInfoResponse)(nil)).Elem()
}

func (i IotDpsSkuInfoResponseArgs) ToIotDpsSkuInfoResponseOutput() IotDpsSkuInfoResponseOutput {
	return i.ToIotDpsSkuInfoResponseOutputWithContext(context.Background())
}

func (i IotDpsSkuInfoResponseArgs) ToIotDpsSkuInfoResponseOutputWithContext(ctx context.Context) IotDpsSkuInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsSkuInfoResponseOutput)
}

func (i IotDpsSkuInfoResponseArgs) ToIotDpsSkuInfoResponsePtrOutput() IotDpsSkuInfoResponsePtrOutput {
	return i.ToIotDpsSkuInfoResponsePtrOutputWithContext(context.Background())
}

func (i IotDpsSkuInfoResponseArgs) ToIotDpsSkuInfoResponsePtrOutputWithContext(ctx context.Context) IotDpsSkuInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsSkuInfoResponseOutput).ToIotDpsSkuInfoResponsePtrOutputWithContext(ctx)
}









type IotDpsSkuInfoResponsePtrInput interface {
	pulumi.Input

	ToIotDpsSkuInfoResponsePtrOutput() IotDpsSkuInfoResponsePtrOutput
	ToIotDpsSkuInfoResponsePtrOutputWithContext(context.Context) IotDpsSkuInfoResponsePtrOutput
}

type iotDpsSkuInfoResponsePtrType IotDpsSkuInfoResponseArgs

func IotDpsSkuInfoResponsePtr(v *IotDpsSkuInfoResponseArgs) IotDpsSkuInfoResponsePtrInput {
	return (*iotDpsSkuInfoResponsePtrType)(v)
}

func (*iotDpsSkuInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IotDpsSkuInfoResponse)(nil)).Elem()
}

func (i *iotDpsSkuInfoResponsePtrType) ToIotDpsSkuInfoResponsePtrOutput() IotDpsSkuInfoResponsePtrOutput {
	return i.ToIotDpsSkuInfoResponsePtrOutputWithContext(context.Background())
}

func (i *iotDpsSkuInfoResponsePtrType) ToIotDpsSkuInfoResponsePtrOutputWithContext(ctx context.Context) IotDpsSkuInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsSkuInfoResponsePtrOutput)
}

type IotDpsSkuInfoResponseOutput struct{ *pulumi.OutputState }

func (IotDpsSkuInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsSkuInfoResponse)(nil)).Elem()
}

func (o IotDpsSkuInfoResponseOutput) ToIotDpsSkuInfoResponseOutput() IotDpsSkuInfoResponseOutput {
	return o
}

func (o IotDpsSkuInfoResponseOutput) ToIotDpsSkuInfoResponseOutputWithContext(ctx context.Context) IotDpsSkuInfoResponseOutput {
	return o
}

func (o IotDpsSkuInfoResponseOutput) ToIotDpsSkuInfoResponsePtrOutput() IotDpsSkuInfoResponsePtrOutput {
	return o.ToIotDpsSkuInfoResponsePtrOutputWithContext(context.Background())
}

func (o IotDpsSkuInfoResponseOutput) ToIotDpsSkuInfoResponsePtrOutputWithContext(ctx context.Context) IotDpsSkuInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IotDpsSkuInfoResponse) *IotDpsSkuInfoResponse {
		return &v
	}).(IotDpsSkuInfoResponsePtrOutput)
}

func (o IotDpsSkuInfoResponseOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v IotDpsSkuInfoResponse) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o IotDpsSkuInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsSkuInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IotDpsSkuInfoResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v IotDpsSkuInfoResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type IotDpsSkuInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (IotDpsSkuInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotDpsSkuInfoResponse)(nil)).Elem()
}

func (o IotDpsSkuInfoResponsePtrOutput) ToIotDpsSkuInfoResponsePtrOutput() IotDpsSkuInfoResponsePtrOutput {
	return o
}

func (o IotDpsSkuInfoResponsePtrOutput) ToIotDpsSkuInfoResponsePtrOutputWithContext(ctx context.Context) IotDpsSkuInfoResponsePtrOutput {
	return o
}

func (o IotDpsSkuInfoResponsePtrOutput) Elem() IotDpsSkuInfoResponseOutput {
	return o.ApplyT(func(v *IotDpsSkuInfoResponse) IotDpsSkuInfoResponse {
		if v != nil {
			return *v
		}
		var ret IotDpsSkuInfoResponse
		return ret
	}).(IotDpsSkuInfoResponseOutput)
}

func (o IotDpsSkuInfoResponsePtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *IotDpsSkuInfoResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o IotDpsSkuInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsSkuInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o IotDpsSkuInfoResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsSkuInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type IotHubDefinitionDescription struct {
	AllocationWeight      *int   `pulumi:"allocationWeight"`
	ApplyAllocationPolicy *bool  `pulumi:"applyAllocationPolicy"`
	ConnectionString      string `pulumi:"connectionString"`
	Location              string `pulumi:"location"`
}





type IotHubDefinitionDescriptionInput interface {
	pulumi.Input

	ToIotHubDefinitionDescriptionOutput() IotHubDefinitionDescriptionOutput
	ToIotHubDefinitionDescriptionOutputWithContext(context.Context) IotHubDefinitionDescriptionOutput
}

type IotHubDefinitionDescriptionArgs struct {
	AllocationWeight      pulumi.IntPtrInput  `pulumi:"allocationWeight"`
	ApplyAllocationPolicy pulumi.BoolPtrInput `pulumi:"applyAllocationPolicy"`
	ConnectionString      pulumi.StringInput  `pulumi:"connectionString"`
	Location              pulumi.StringInput  `pulumi:"location"`
}

func (IotHubDefinitionDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubDefinitionDescription)(nil)).Elem()
}

func (i IotHubDefinitionDescriptionArgs) ToIotHubDefinitionDescriptionOutput() IotHubDefinitionDescriptionOutput {
	return i.ToIotHubDefinitionDescriptionOutputWithContext(context.Background())
}

func (i IotHubDefinitionDescriptionArgs) ToIotHubDefinitionDescriptionOutputWithContext(ctx context.Context) IotHubDefinitionDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubDefinitionDescriptionOutput)
}





type IotHubDefinitionDescriptionArrayInput interface {
	pulumi.Input

	ToIotHubDefinitionDescriptionArrayOutput() IotHubDefinitionDescriptionArrayOutput
	ToIotHubDefinitionDescriptionArrayOutputWithContext(context.Context) IotHubDefinitionDescriptionArrayOutput
}

type IotHubDefinitionDescriptionArray []IotHubDefinitionDescriptionInput

func (IotHubDefinitionDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubDefinitionDescription)(nil)).Elem()
}

func (i IotHubDefinitionDescriptionArray) ToIotHubDefinitionDescriptionArrayOutput() IotHubDefinitionDescriptionArrayOutput {
	return i.ToIotHubDefinitionDescriptionArrayOutputWithContext(context.Background())
}

func (i IotHubDefinitionDescriptionArray) ToIotHubDefinitionDescriptionArrayOutputWithContext(ctx context.Context) IotHubDefinitionDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubDefinitionDescriptionArrayOutput)
}

type IotHubDefinitionDescriptionOutput struct{ *pulumi.OutputState }

func (IotHubDefinitionDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubDefinitionDescription)(nil)).Elem()
}

func (o IotHubDefinitionDescriptionOutput) ToIotHubDefinitionDescriptionOutput() IotHubDefinitionDescriptionOutput {
	return o
}

func (o IotHubDefinitionDescriptionOutput) ToIotHubDefinitionDescriptionOutputWithContext(ctx context.Context) IotHubDefinitionDescriptionOutput {
	return o
}

func (o IotHubDefinitionDescriptionOutput) AllocationWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IotHubDefinitionDescription) *int { return v.AllocationWeight }).(pulumi.IntPtrOutput)
}

func (o IotHubDefinitionDescriptionOutput) ApplyAllocationPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IotHubDefinitionDescription) *bool { return v.ApplyAllocationPolicy }).(pulumi.BoolPtrOutput)
}

func (o IotHubDefinitionDescriptionOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubDefinitionDescription) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o IotHubDefinitionDescriptionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubDefinitionDescription) string { return v.Location }).(pulumi.StringOutput)
}

type IotHubDefinitionDescriptionArrayOutput struct{ *pulumi.OutputState }

func (IotHubDefinitionDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubDefinitionDescription)(nil)).Elem()
}

func (o IotHubDefinitionDescriptionArrayOutput) ToIotHubDefinitionDescriptionArrayOutput() IotHubDefinitionDescriptionArrayOutput {
	return o
}

func (o IotHubDefinitionDescriptionArrayOutput) ToIotHubDefinitionDescriptionArrayOutputWithContext(ctx context.Context) IotHubDefinitionDescriptionArrayOutput {
	return o
}

func (o IotHubDefinitionDescriptionArrayOutput) Index(i pulumi.IntInput) IotHubDefinitionDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IotHubDefinitionDescription {
		return vs[0].([]IotHubDefinitionDescription)[vs[1].(int)]
	}).(IotHubDefinitionDescriptionOutput)
}

type IotHubDefinitionDescriptionResponse struct {
	AllocationWeight      *int   `pulumi:"allocationWeight"`
	ApplyAllocationPolicy *bool  `pulumi:"applyAllocationPolicy"`
	ConnectionString      string `pulumi:"connectionString"`
	Location              string `pulumi:"location"`
	Name                  string `pulumi:"name"`
}





type IotHubDefinitionDescriptionResponseInput interface {
	pulumi.Input

	ToIotHubDefinitionDescriptionResponseOutput() IotHubDefinitionDescriptionResponseOutput
	ToIotHubDefinitionDescriptionResponseOutputWithContext(context.Context) IotHubDefinitionDescriptionResponseOutput
}

type IotHubDefinitionDescriptionResponseArgs struct {
	AllocationWeight      pulumi.IntPtrInput  `pulumi:"allocationWeight"`
	ApplyAllocationPolicy pulumi.BoolPtrInput `pulumi:"applyAllocationPolicy"`
	ConnectionString      pulumi.StringInput  `pulumi:"connectionString"`
	Location              pulumi.StringInput  `pulumi:"location"`
	Name                  pulumi.StringInput  `pulumi:"name"`
}

func (IotHubDefinitionDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubDefinitionDescriptionResponse)(nil)).Elem()
}

func (i IotHubDefinitionDescriptionResponseArgs) ToIotHubDefinitionDescriptionResponseOutput() IotHubDefinitionDescriptionResponseOutput {
	return i.ToIotHubDefinitionDescriptionResponseOutputWithContext(context.Background())
}

func (i IotHubDefinitionDescriptionResponseArgs) ToIotHubDefinitionDescriptionResponseOutputWithContext(ctx context.Context) IotHubDefinitionDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubDefinitionDescriptionResponseOutput)
}





type IotHubDefinitionDescriptionResponseArrayInput interface {
	pulumi.Input

	ToIotHubDefinitionDescriptionResponseArrayOutput() IotHubDefinitionDescriptionResponseArrayOutput
	ToIotHubDefinitionDescriptionResponseArrayOutputWithContext(context.Context) IotHubDefinitionDescriptionResponseArrayOutput
}

type IotHubDefinitionDescriptionResponseArray []IotHubDefinitionDescriptionResponseInput

func (IotHubDefinitionDescriptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubDefinitionDescriptionResponse)(nil)).Elem()
}

func (i IotHubDefinitionDescriptionResponseArray) ToIotHubDefinitionDescriptionResponseArrayOutput() IotHubDefinitionDescriptionResponseArrayOutput {
	return i.ToIotHubDefinitionDescriptionResponseArrayOutputWithContext(context.Background())
}

func (i IotHubDefinitionDescriptionResponseArray) ToIotHubDefinitionDescriptionResponseArrayOutputWithContext(ctx context.Context) IotHubDefinitionDescriptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubDefinitionDescriptionResponseArrayOutput)
}

type IotHubDefinitionDescriptionResponseOutput struct{ *pulumi.OutputState }

func (IotHubDefinitionDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubDefinitionDescriptionResponse)(nil)).Elem()
}

func (o IotHubDefinitionDescriptionResponseOutput) ToIotHubDefinitionDescriptionResponseOutput() IotHubDefinitionDescriptionResponseOutput {
	return o
}

func (o IotHubDefinitionDescriptionResponseOutput) ToIotHubDefinitionDescriptionResponseOutputWithContext(ctx context.Context) IotHubDefinitionDescriptionResponseOutput {
	return o
}

func (o IotHubDefinitionDescriptionResponseOutput) AllocationWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IotHubDefinitionDescriptionResponse) *int { return v.AllocationWeight }).(pulumi.IntPtrOutput)
}

func (o IotHubDefinitionDescriptionResponseOutput) ApplyAllocationPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IotHubDefinitionDescriptionResponse) *bool { return v.ApplyAllocationPolicy }).(pulumi.BoolPtrOutput)
}

func (o IotHubDefinitionDescriptionResponseOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubDefinitionDescriptionResponse) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o IotHubDefinitionDescriptionResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubDefinitionDescriptionResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o IotHubDefinitionDescriptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubDefinitionDescriptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

type IotHubDefinitionDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (IotHubDefinitionDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubDefinitionDescriptionResponse)(nil)).Elem()
}

func (o IotHubDefinitionDescriptionResponseArrayOutput) ToIotHubDefinitionDescriptionResponseArrayOutput() IotHubDefinitionDescriptionResponseArrayOutput {
	return o
}

func (o IotHubDefinitionDescriptionResponseArrayOutput) ToIotHubDefinitionDescriptionResponseArrayOutputWithContext(ctx context.Context) IotHubDefinitionDescriptionResponseArrayOutput {
	return o
}

func (o IotHubDefinitionDescriptionResponseArrayOutput) Index(i pulumi.IntInput) IotHubDefinitionDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IotHubDefinitionDescriptionResponse {
		return vs[0].([]IotHubDefinitionDescriptionResponse)[vs[1].(int)]
	}).(IotHubDefinitionDescriptionResponseOutput)
}

type PrivateEndpointConnection struct {
	Properties PrivateEndpointConnectionProperties `pulumi:"properties"`
}





type PrivateEndpointConnectionInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput
	ToPrivateEndpointConnectionOutputWithContext(context.Context) PrivateEndpointConnectionOutput
}

type PrivateEndpointConnectionArgs struct {
	Properties PrivateEndpointConnectionPropertiesInput `pulumi:"properties"`
}

func (PrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnection)(nil)).Elem()
}

func (i PrivateEndpointConnectionArgs) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return i.ToPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionArgs) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionOutput)
}





type PrivateEndpointConnectionArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionArrayOutput() PrivateEndpointConnectionArrayOutput
	ToPrivateEndpointConnectionArrayOutputWithContext(context.Context) PrivateEndpointConnectionArrayOutput
}

type PrivateEndpointConnectionArray []PrivateEndpointConnectionInput

func (PrivateEndpointConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnection)(nil)).Elem()
}

func (i PrivateEndpointConnectionArray) ToPrivateEndpointConnectionArrayOutput() PrivateEndpointConnectionArrayOutput {
	return i.ToPrivateEndpointConnectionArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionArray) ToPrivateEndpointConnectionArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionArrayOutput)
}

type PrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnection)(nil)).Elem()
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) Properties() PrivateEndpointConnectionPropertiesOutput {
	return o.ApplyT(func(v PrivateEndpointConnection) PrivateEndpointConnectionProperties { return v.Properties }).(PrivateEndpointConnectionPropertiesOutput)
}

type PrivateEndpointConnectionArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnection)(nil)).Elem()
}

func (o PrivateEndpointConnectionArrayOutput) ToPrivateEndpointConnectionArrayOutput() PrivateEndpointConnectionArrayOutput {
	return o
}

func (o PrivateEndpointConnectionArrayOutput) ToPrivateEndpointConnectionArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionArrayOutput {
	return o
}

func (o PrivateEndpointConnectionArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnection {
		return vs[0].([]PrivateEndpointConnection)[vs[1].(int)]
	}).(PrivateEndpointConnectionOutput)
}

type PrivateEndpointConnectionProperties struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionPropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput
	ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesOutput
}

type PrivateEndpointConnectionPropertiesArgs struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput `pulumi:"privateLinkServiceConnectionState"`
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

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput).ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesPtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput
	ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPtrOutput
}

type privateEndpointConnectionPropertiesPtrType PrivateEndpointConnectionPropertiesArgs

func PrivateEndpointConnectionPropertiesPtr(v *PrivateEndpointConnectionPropertiesArgs) PrivateEndpointConnectionPropertiesPtrInput {
	return (*privateEndpointConnectionPropertiesPtrType)(v)
}

func (*privateEndpointConnectionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPtrOutput)
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

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionProperties) *PrivateEndpointConnectionProperties {
		return &v
	}).(PrivateEndpointConnectionPropertiesPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateEndpointConnectionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) Elem() PrivateEndpointConnectionPropertiesOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) PrivateEndpointConnectionProperties {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionProperties
		return ret
	}).(PrivateEndpointConnectionPropertiesOutput)
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) *PrivateLinkServiceConnectionState {
		if v == nil {
			return nil
		}
		return &v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponse struct {
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionPropertiesResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput
	ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponseOutput
}

type PrivateEndpointConnectionPropertiesResponseArgs struct {
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponseOutput).ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput
	ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput
}

type privateEndpointConnectionPropertiesResponsePtrType PrivateEndpointConnectionPropertiesResponseArgs

func PrivateEndpointConnectionPropertiesResponsePtr(v *PrivateEndpointConnectionPropertiesResponseArgs) PrivateEndpointConnectionPropertiesResponsePtrInput {
	return (*privateEndpointConnectionPropertiesResponsePtrType)(v)
}

func (*privateEndpointConnectionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesResponsePtrType) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesResponsePtrType) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointConnectionPropertiesResponse {
		return &v
	}).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

type PrivateEndpointConnectionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) Elem() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) PrivateEndpointConnectionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesResponse
		return ret
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateLinkServiceConnectionStateResponse {
		if v == nil {
			return nil
		}
		return &v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                          `pulumi:"systemData"`
	Type       string                                      `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Id         pulumi.StringInput                               `pulumi:"id"`
	Name       pulumi.StringInput                               `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponseInput `pulumi:"properties"`
	SystemData SystemDataResponseInput                          `pulumi:"systemData"`
	Type       pulumi.StringInput                               `pulumi:"type"`
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

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
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





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
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

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     string  `pulumi:"description"`
	Status          string  `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringInput    `pulumi:"description"`
	Status          pulumi.StringInput    `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     string  `pulumi:"description"`
	Status          string  `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringInput    `pulumi:"description"`
	Status          pulumi.StringInput    `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescription struct {
	KeyName      string  `pulumi:"keyName"`
	PrimaryKey   *string `pulumi:"primaryKey"`
	Rights       string  `pulumi:"rights"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}





type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionInput interface {
	pulumi.Input

	ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput
	ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutputWithContext(context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArgs struct {
	KeyName      pulumi.StringInput    `pulumi:"keyName"`
	PrimaryKey   pulumi.StringPtrInput `pulumi:"primaryKey"`
	Rights       pulumi.StringInput    `pulumi:"rights"`
	SecondaryKey pulumi.StringPtrInput `pulumi:"secondaryKey"`
}

func (SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleAccessRightsDescription)(nil)).Elem()
}

func (i SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArgs) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput {
	return i.ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutputWithContext(context.Background())
}

func (i SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArgs) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput)
}





type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayInput interface {
	pulumi.Input

	ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput
	ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutputWithContext(context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArray []SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionInput

func (SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessSignatureAuthorizationRuleAccessRightsDescription)(nil)).Elem()
}

func (i SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArray) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput {
	return i.ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutputWithContext(context.Background())
}

func (i SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArray) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput)
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleAccessRightsDescription)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescription) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescription) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput) Rights() pulumi.StringOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescription) string { return v.Rights }).(pulumi.StringOutput)
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescription) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessSignatureAuthorizationRuleAccessRightsDescription)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput) Index(i pulumi.IntInput) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedAccessSignatureAuthorizationRuleAccessRightsDescription {
		return vs[0].([]SharedAccessSignatureAuthorizationRuleAccessRightsDescription)[vs[1].(int)]
	}).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput)
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse struct {
	KeyName      string  `pulumi:"keyName"`
	PrimaryKey   *string `pulumi:"primaryKey"`
	Rights       string  `pulumi:"rights"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}





type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseInput interface {
	pulumi.Input

	ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput
	ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutputWithContext(context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArgs struct {
	KeyName      pulumi.StringInput    `pulumi:"keyName"`
	PrimaryKey   pulumi.StringPtrInput `pulumi:"primaryKey"`
	Rights       pulumi.StringInput    `pulumi:"rights"`
	SecondaryKey pulumi.StringPtrInput `pulumi:"secondaryKey"`
}

func (SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse)(nil)).Elem()
}

func (i SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArgs) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput {
	return i.ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutputWithContext(context.Background())
}

func (i SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArgs) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput)
}





type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayInput interface {
	pulumi.Input

	ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput
	ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutputWithContext(context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArray []SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseInput

func (SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse)(nil)).Elem()
}

func (i SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArray) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput {
	return i.ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutputWithContext(context.Background())
}

func (i SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArray) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput)
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse) *string {
		return v.PrimaryKey
	}).(pulumi.StringPtrOutput)
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput) Rights() pulumi.StringOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse) string { return v.Rights }).(pulumi.StringOutput)
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse) *string {
		return v.SecondaryKey
	}).(pulumi.StringPtrOutput)
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput) Index(i pulumi.IntInput) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse {
		return vs[0].([]SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse)[vs[1].(int)]
	}).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput)
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

type TargetIpFilterRule struct {
	Action     IpFilterActionType  `pulumi:"action"`
	FilterName string              `pulumi:"filterName"`
	IpMask     string              `pulumi:"ipMask"`
	Target     *IpFilterTargetType `pulumi:"target"`
}





type TargetIpFilterRuleInput interface {
	pulumi.Input

	ToTargetIpFilterRuleOutput() TargetIpFilterRuleOutput
	ToTargetIpFilterRuleOutputWithContext(context.Context) TargetIpFilterRuleOutput
}

type TargetIpFilterRuleArgs struct {
	Action     IpFilterActionTypeInput    `pulumi:"action"`
	FilterName pulumi.StringInput         `pulumi:"filterName"`
	IpMask     pulumi.StringInput         `pulumi:"ipMask"`
	Target     IpFilterTargetTypePtrInput `pulumi:"target"`
}

func (TargetIpFilterRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetIpFilterRule)(nil)).Elem()
}

func (i TargetIpFilterRuleArgs) ToTargetIpFilterRuleOutput() TargetIpFilterRuleOutput {
	return i.ToTargetIpFilterRuleOutputWithContext(context.Background())
}

func (i TargetIpFilterRuleArgs) ToTargetIpFilterRuleOutputWithContext(ctx context.Context) TargetIpFilterRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetIpFilterRuleOutput)
}





type TargetIpFilterRuleArrayInput interface {
	pulumi.Input

	ToTargetIpFilterRuleArrayOutput() TargetIpFilterRuleArrayOutput
	ToTargetIpFilterRuleArrayOutputWithContext(context.Context) TargetIpFilterRuleArrayOutput
}

type TargetIpFilterRuleArray []TargetIpFilterRuleInput

func (TargetIpFilterRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetIpFilterRule)(nil)).Elem()
}

func (i TargetIpFilterRuleArray) ToTargetIpFilterRuleArrayOutput() TargetIpFilterRuleArrayOutput {
	return i.ToTargetIpFilterRuleArrayOutputWithContext(context.Background())
}

func (i TargetIpFilterRuleArray) ToTargetIpFilterRuleArrayOutputWithContext(ctx context.Context) TargetIpFilterRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetIpFilterRuleArrayOutput)
}

type TargetIpFilterRuleOutput struct{ *pulumi.OutputState }

func (TargetIpFilterRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetIpFilterRule)(nil)).Elem()
}

func (o TargetIpFilterRuleOutput) ToTargetIpFilterRuleOutput() TargetIpFilterRuleOutput {
	return o
}

func (o TargetIpFilterRuleOutput) ToTargetIpFilterRuleOutputWithContext(ctx context.Context) TargetIpFilterRuleOutput {
	return o
}

func (o TargetIpFilterRuleOutput) Action() IpFilterActionTypeOutput {
	return o.ApplyT(func(v TargetIpFilterRule) IpFilterActionType { return v.Action }).(IpFilterActionTypeOutput)
}

func (o TargetIpFilterRuleOutput) FilterName() pulumi.StringOutput {
	return o.ApplyT(func(v TargetIpFilterRule) string { return v.FilterName }).(pulumi.StringOutput)
}

func (o TargetIpFilterRuleOutput) IpMask() pulumi.StringOutput {
	return o.ApplyT(func(v TargetIpFilterRule) string { return v.IpMask }).(pulumi.StringOutput)
}

func (o TargetIpFilterRuleOutput) Target() IpFilterTargetTypePtrOutput {
	return o.ApplyT(func(v TargetIpFilterRule) *IpFilterTargetType { return v.Target }).(IpFilterTargetTypePtrOutput)
}

type TargetIpFilterRuleArrayOutput struct{ *pulumi.OutputState }

func (TargetIpFilterRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetIpFilterRule)(nil)).Elem()
}

func (o TargetIpFilterRuleArrayOutput) ToTargetIpFilterRuleArrayOutput() TargetIpFilterRuleArrayOutput {
	return o
}

func (o TargetIpFilterRuleArrayOutput) ToTargetIpFilterRuleArrayOutputWithContext(ctx context.Context) TargetIpFilterRuleArrayOutput {
	return o
}

func (o TargetIpFilterRuleArrayOutput) Index(i pulumi.IntInput) TargetIpFilterRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetIpFilterRule {
		return vs[0].([]TargetIpFilterRule)[vs[1].(int)]
	}).(TargetIpFilterRuleOutput)
}

type TargetIpFilterRuleResponse struct {
	Action     string  `pulumi:"action"`
	FilterName string  `pulumi:"filterName"`
	IpMask     string  `pulumi:"ipMask"`
	Target     *string `pulumi:"target"`
}





type TargetIpFilterRuleResponseInput interface {
	pulumi.Input

	ToTargetIpFilterRuleResponseOutput() TargetIpFilterRuleResponseOutput
	ToTargetIpFilterRuleResponseOutputWithContext(context.Context) TargetIpFilterRuleResponseOutput
}

type TargetIpFilterRuleResponseArgs struct {
	Action     pulumi.StringInput    `pulumi:"action"`
	FilterName pulumi.StringInput    `pulumi:"filterName"`
	IpMask     pulumi.StringInput    `pulumi:"ipMask"`
	Target     pulumi.StringPtrInput `pulumi:"target"`
}

func (TargetIpFilterRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetIpFilterRuleResponse)(nil)).Elem()
}

func (i TargetIpFilterRuleResponseArgs) ToTargetIpFilterRuleResponseOutput() TargetIpFilterRuleResponseOutput {
	return i.ToTargetIpFilterRuleResponseOutputWithContext(context.Background())
}

func (i TargetIpFilterRuleResponseArgs) ToTargetIpFilterRuleResponseOutputWithContext(ctx context.Context) TargetIpFilterRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetIpFilterRuleResponseOutput)
}





type TargetIpFilterRuleResponseArrayInput interface {
	pulumi.Input

	ToTargetIpFilterRuleResponseArrayOutput() TargetIpFilterRuleResponseArrayOutput
	ToTargetIpFilterRuleResponseArrayOutputWithContext(context.Context) TargetIpFilterRuleResponseArrayOutput
}

type TargetIpFilterRuleResponseArray []TargetIpFilterRuleResponseInput

func (TargetIpFilterRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetIpFilterRuleResponse)(nil)).Elem()
}

func (i TargetIpFilterRuleResponseArray) ToTargetIpFilterRuleResponseArrayOutput() TargetIpFilterRuleResponseArrayOutput {
	return i.ToTargetIpFilterRuleResponseArrayOutputWithContext(context.Background())
}

func (i TargetIpFilterRuleResponseArray) ToTargetIpFilterRuleResponseArrayOutputWithContext(ctx context.Context) TargetIpFilterRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetIpFilterRuleResponseArrayOutput)
}

type TargetIpFilterRuleResponseOutput struct{ *pulumi.OutputState }

func (TargetIpFilterRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetIpFilterRuleResponse)(nil)).Elem()
}

func (o TargetIpFilterRuleResponseOutput) ToTargetIpFilterRuleResponseOutput() TargetIpFilterRuleResponseOutput {
	return o
}

func (o TargetIpFilterRuleResponseOutput) ToTargetIpFilterRuleResponseOutputWithContext(ctx context.Context) TargetIpFilterRuleResponseOutput {
	return o
}

func (o TargetIpFilterRuleResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v TargetIpFilterRuleResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o TargetIpFilterRuleResponseOutput) FilterName() pulumi.StringOutput {
	return o.ApplyT(func(v TargetIpFilterRuleResponse) string { return v.FilterName }).(pulumi.StringOutput)
}

func (o TargetIpFilterRuleResponseOutput) IpMask() pulumi.StringOutput {
	return o.ApplyT(func(v TargetIpFilterRuleResponse) string { return v.IpMask }).(pulumi.StringOutput)
}

func (o TargetIpFilterRuleResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetIpFilterRuleResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type TargetIpFilterRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (TargetIpFilterRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetIpFilterRuleResponse)(nil)).Elem()
}

func (o TargetIpFilterRuleResponseArrayOutput) ToTargetIpFilterRuleResponseArrayOutput() TargetIpFilterRuleResponseArrayOutput {
	return o
}

func (o TargetIpFilterRuleResponseArrayOutput) ToTargetIpFilterRuleResponseArrayOutputWithContext(ctx context.Context) TargetIpFilterRuleResponseArrayOutput {
	return o
}

func (o TargetIpFilterRuleResponseArrayOutput) Index(i pulumi.IntInput) TargetIpFilterRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetIpFilterRuleResponse {
		return vs[0].([]TargetIpFilterRuleResponse)[vs[1].(int)]
	}).(TargetIpFilterRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificatePropertiesResponseOutput{})
	pulumi.RegisterOutputType(CertificatePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(IotDpsPropertiesDescriptionOutput{})
	pulumi.RegisterOutputType(IotDpsPropertiesDescriptionPtrOutput{})
	pulumi.RegisterOutputType(IotDpsPropertiesDescriptionResponseOutput{})
	pulumi.RegisterOutputType(IotDpsPropertiesDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(IotDpsSkuInfoOutput{})
	pulumi.RegisterOutputType(IotDpsSkuInfoPtrOutput{})
	pulumi.RegisterOutputType(IotDpsSkuInfoResponseOutput{})
	pulumi.RegisterOutputType(IotDpsSkuInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(IotHubDefinitionDescriptionOutput{})
	pulumi.RegisterOutputType(IotHubDefinitionDescriptionArrayOutput{})
	pulumi.RegisterOutputType(IotHubDefinitionDescriptionResponseOutput{})
	pulumi.RegisterOutputType(IotHubDefinitionDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TargetIpFilterRuleOutput{})
	pulumi.RegisterOutputType(TargetIpFilterRuleArrayOutput{})
	pulumi.RegisterOutputType(TargetIpFilterRuleResponseOutput{})
	pulumi.RegisterOutputType(TargetIpFilterRuleResponseArrayOutput{})
}

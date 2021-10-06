


package v20170821preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CertificatePropertiesResponse struct {
	Created    string `pulumi:"created"`
	Expiry     string `pulumi:"expiry"`
	IsVerified bool   `pulumi:"isVerified"`
	Subject    string `pulumi:"subject"`
	Thumbprint string `pulumi:"thumbprint"`
	Updated    string `pulumi:"updated"`
}





type CertificatePropertiesResponseInput interface {
	pulumi.Input

	ToCertificatePropertiesResponseOutput() CertificatePropertiesResponseOutput
	ToCertificatePropertiesResponseOutputWithContext(context.Context) CertificatePropertiesResponseOutput
}

type CertificatePropertiesResponseArgs struct {
	Created    pulumi.StringInput `pulumi:"created"`
	Expiry     pulumi.StringInput `pulumi:"expiry"`
	IsVerified pulumi.BoolInput   `pulumi:"isVerified"`
	Subject    pulumi.StringInput `pulumi:"subject"`
	Thumbprint pulumi.StringInput `pulumi:"thumbprint"`
	Updated    pulumi.StringInput `pulumi:"updated"`
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
	AllocationPolicy      *string                                                         `pulumi:"allocationPolicy"`
	AuthorizationPolicies []SharedAccessSignatureAuthorizationRuleAccessRightsDescription `pulumi:"authorizationPolicies"`
	IotHubs               []IotHubDefinitionDescription                                   `pulumi:"iotHubs"`
	ProvisioningState     *string                                                         `pulumi:"provisioningState"`
	State                 *string                                                         `pulumi:"state"`
}





type IotDpsPropertiesDescriptionInput interface {
	pulumi.Input

	ToIotDpsPropertiesDescriptionOutput() IotDpsPropertiesDescriptionOutput
	ToIotDpsPropertiesDescriptionOutputWithContext(context.Context) IotDpsPropertiesDescriptionOutput
}

type IotDpsPropertiesDescriptionArgs struct {
	AllocationPolicy      pulumi.StringPtrInput                                                   `pulumi:"allocationPolicy"`
	AuthorizationPolicies SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayInput `pulumi:"authorizationPolicies"`
	IotHubs               IotHubDefinitionDescriptionArrayInput                                   `pulumi:"iotHubs"`
	ProvisioningState     pulumi.StringPtrInput                                                   `pulumi:"provisioningState"`
	State                 pulumi.StringPtrInput                                                   `pulumi:"state"`
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

func (o IotDpsPropertiesDescriptionOutput) IotHubs() IotHubDefinitionDescriptionArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) []IotHubDefinitionDescription { return v.IotHubs }).(IotHubDefinitionDescriptionArrayOutput)
}

func (o IotDpsPropertiesDescriptionOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
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

func (o IotDpsPropertiesDescriptionPtrOutput) IotHubs() IotHubDefinitionDescriptionArrayOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescription) []IotHubDefinitionDescription {
		if v == nil {
			return nil
		}
		return v.IotHubs
	}).(IotHubDefinitionDescriptionArrayOutput)
}

func (o IotDpsPropertiesDescriptionPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescription) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
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
	IdScope                    string                                                                  `pulumi:"idScope"`
	IotHubs                    []IotHubDefinitionDescriptionResponse                                   `pulumi:"iotHubs"`
	ProvisioningState          *string                                                                 `pulumi:"provisioningState"`
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
	IdScope                    pulumi.StringInput                                                              `pulumi:"idScope"`
	IotHubs                    IotHubDefinitionDescriptionResponseArrayInput                                   `pulumi:"iotHubs"`
	ProvisioningState          pulumi.StringPtrInput                                                           `pulumi:"provisioningState"`
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

func (o IotDpsPropertiesDescriptionResponseOutput) IdScope() pulumi.StringOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) string { return v.IdScope }).(pulumi.StringOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) IotHubs() IotHubDefinitionDescriptionResponseArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) []IotHubDefinitionDescriptionResponse { return v.IotHubs }).(IotHubDefinitionDescriptionResponseArrayOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
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

func (o IotDpsPropertiesDescriptionResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsPropertiesDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
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
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput{})
}

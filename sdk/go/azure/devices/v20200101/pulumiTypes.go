


package v20200101

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

type IotDpsPropertiesDescription struct {
	AllocationPolicy      *string                                                         `pulumi:"allocationPolicy"`
	AuthorizationPolicies []SharedAccessSignatureAuthorizationRuleAccessRightsDescription `pulumi:"authorizationPolicies"`
	IotHubs               []IotHubDefinitionDescription                                   `pulumi:"iotHubs"`
	IpFilterRules         []TargetIpFilterRule                                            `pulumi:"ipFilterRules"`
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
	IpFilterRules         TargetIpFilterRuleArrayInput                                            `pulumi:"ipFilterRules"`
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

func (o IotDpsPropertiesDescriptionOutput) IpFilterRules() TargetIpFilterRuleArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) []TargetIpFilterRule { return v.IpFilterRules }).(TargetIpFilterRuleArrayOutput)
}

func (o IotDpsPropertiesDescriptionOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) *string { return v.State }).(pulumi.StringPtrOutput)
}

type IotDpsPropertiesDescriptionResponse struct {
	AllocationPolicy           *string                                                                 `pulumi:"allocationPolicy"`
	AuthorizationPolicies      []SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse `pulumi:"authorizationPolicies"`
	DeviceProvisioningHostName string                                                                  `pulumi:"deviceProvisioningHostName"`
	IdScope                    string                                                                  `pulumi:"idScope"`
	IotHubs                    []IotHubDefinitionDescriptionResponse                                   `pulumi:"iotHubs"`
	IpFilterRules              []TargetIpFilterRuleResponse                                            `pulumi:"ipFilterRules"`
	ProvisioningState          *string                                                                 `pulumi:"provisioningState"`
	ServiceOperationsHostName  string                                                                  `pulumi:"serviceOperationsHostName"`
	State                      *string                                                                 `pulumi:"state"`
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

func (o IotDpsPropertiesDescriptionResponseOutput) IpFilterRules() TargetIpFilterRuleResponseArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) []TargetIpFilterRuleResponse { return v.IpFilterRules }).(TargetIpFilterRuleResponseArrayOutput)
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

func (o IotDpsSkuInfoOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v IotDpsSkuInfo) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o IotDpsSkuInfoOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsSkuInfo) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type IotDpsSkuInfoResponse struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
	Tier     string   `pulumi:"tier"`
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

func (o IotDpsSkuInfoResponseOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v IotDpsSkuInfoResponse) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o IotDpsSkuInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsSkuInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IotDpsSkuInfoResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v IotDpsSkuInfoResponse) string { return v.Tier }).(pulumi.StringOutput)
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
	pulumi.RegisterOutputType(IotDpsPropertiesDescriptionOutput{})
	pulumi.RegisterOutputType(IotDpsPropertiesDescriptionResponseOutput{})
	pulumi.RegisterOutputType(IotDpsSkuInfoOutput{})
	pulumi.RegisterOutputType(IotDpsSkuInfoResponseOutput{})
	pulumi.RegisterOutputType(IotHubDefinitionDescriptionOutput{})
	pulumi.RegisterOutputType(IotHubDefinitionDescriptionArrayOutput{})
	pulumi.RegisterOutputType(IotHubDefinitionDescriptionResponseOutput{})
	pulumi.RegisterOutputType(IotHubDefinitionDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(TargetIpFilterRuleOutput{})
	pulumi.RegisterOutputType(TargetIpFilterRuleArrayOutput{})
	pulumi.RegisterOutputType(TargetIpFilterRuleResponseOutput{})
	pulumi.RegisterOutputType(TargetIpFilterRuleResponseArrayOutput{})
}

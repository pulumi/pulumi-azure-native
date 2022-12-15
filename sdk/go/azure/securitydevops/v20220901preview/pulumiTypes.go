


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthorizationInfo struct {
	Code *string `pulumi:"code"`
}





type AuthorizationInfoInput interface {
	pulumi.Input

	ToAuthorizationInfoOutput() AuthorizationInfoOutput
	ToAuthorizationInfoOutputWithContext(context.Context) AuthorizationInfoOutput
}

type AuthorizationInfoArgs struct {
	Code pulumi.StringPtrInput `pulumi:"code"`
}

func (AuthorizationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationInfo)(nil)).Elem()
}

func (i AuthorizationInfoArgs) ToAuthorizationInfoOutput() AuthorizationInfoOutput {
	return i.ToAuthorizationInfoOutputWithContext(context.Background())
}

func (i AuthorizationInfoArgs) ToAuthorizationInfoOutputWithContext(ctx context.Context) AuthorizationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationInfoOutput)
}

func (i AuthorizationInfoArgs) ToAuthorizationInfoPtrOutput() AuthorizationInfoPtrOutput {
	return i.ToAuthorizationInfoPtrOutputWithContext(context.Background())
}

func (i AuthorizationInfoArgs) ToAuthorizationInfoPtrOutputWithContext(ctx context.Context) AuthorizationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationInfoOutput).ToAuthorizationInfoPtrOutputWithContext(ctx)
}









type AuthorizationInfoPtrInput interface {
	pulumi.Input

	ToAuthorizationInfoPtrOutput() AuthorizationInfoPtrOutput
	ToAuthorizationInfoPtrOutputWithContext(context.Context) AuthorizationInfoPtrOutput
}

type authorizationInfoPtrType AuthorizationInfoArgs

func AuthorizationInfoPtr(v *AuthorizationInfoArgs) AuthorizationInfoPtrInput {
	return (*authorizationInfoPtrType)(v)
}

func (*authorizationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationInfo)(nil)).Elem()
}

func (i *authorizationInfoPtrType) ToAuthorizationInfoPtrOutput() AuthorizationInfoPtrOutput {
	return i.ToAuthorizationInfoPtrOutputWithContext(context.Background())
}

func (i *authorizationInfoPtrType) ToAuthorizationInfoPtrOutputWithContext(ctx context.Context) AuthorizationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationInfoPtrOutput)
}

type AuthorizationInfoOutput struct{ *pulumi.OutputState }

func (AuthorizationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationInfo)(nil)).Elem()
}

func (o AuthorizationInfoOutput) ToAuthorizationInfoOutput() AuthorizationInfoOutput {
	return o
}

func (o AuthorizationInfoOutput) ToAuthorizationInfoOutputWithContext(ctx context.Context) AuthorizationInfoOutput {
	return o
}

func (o AuthorizationInfoOutput) ToAuthorizationInfoPtrOutput() AuthorizationInfoPtrOutput {
	return o.ToAuthorizationInfoPtrOutputWithContext(context.Background())
}

func (o AuthorizationInfoOutput) ToAuthorizationInfoPtrOutputWithContext(ctx context.Context) AuthorizationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthorizationInfo) *AuthorizationInfo {
		return &v
	}).(AuthorizationInfoPtrOutput)
}

func (o AuthorizationInfoOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthorizationInfo) *string { return v.Code }).(pulumi.StringPtrOutput)
}

type AuthorizationInfoPtrOutput struct{ *pulumi.OutputState }

func (AuthorizationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationInfo)(nil)).Elem()
}

func (o AuthorizationInfoPtrOutput) ToAuthorizationInfoPtrOutput() AuthorizationInfoPtrOutput {
	return o
}

func (o AuthorizationInfoPtrOutput) ToAuthorizationInfoPtrOutputWithContext(ctx context.Context) AuthorizationInfoPtrOutput {
	return o
}

func (o AuthorizationInfoPtrOutput) Elem() AuthorizationInfoOutput {
	return o.ApplyT(func(v *AuthorizationInfo) AuthorizationInfo {
		if v != nil {
			return *v
		}
		var ret AuthorizationInfo
		return ret
	}).(AuthorizationInfoOutput)
}

func (o AuthorizationInfoPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationInfo) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

type AuthorizationInfoResponse struct {
	Code *string `pulumi:"code"`
}

type AuthorizationInfoResponseOutput struct{ *pulumi.OutputState }

func (AuthorizationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationInfoResponse)(nil)).Elem()
}

func (o AuthorizationInfoResponseOutput) ToAuthorizationInfoResponseOutput() AuthorizationInfoResponseOutput {
	return o
}

func (o AuthorizationInfoResponseOutput) ToAuthorizationInfoResponseOutputWithContext(ctx context.Context) AuthorizationInfoResponseOutput {
	return o
}

func (o AuthorizationInfoResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthorizationInfoResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

type AuthorizationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (AuthorizationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationInfoResponse)(nil)).Elem()
}

func (o AuthorizationInfoResponsePtrOutput) ToAuthorizationInfoResponsePtrOutput() AuthorizationInfoResponsePtrOutput {
	return o
}

func (o AuthorizationInfoResponsePtrOutput) ToAuthorizationInfoResponsePtrOutputWithContext(ctx context.Context) AuthorizationInfoResponsePtrOutput {
	return o
}

func (o AuthorizationInfoResponsePtrOutput) Elem() AuthorizationInfoResponseOutput {
	return o.ApplyT(func(v *AuthorizationInfoResponse) AuthorizationInfoResponse {
		if v != nil {
			return *v
		}
		var ret AuthorizationInfoResponse
		return ret
	}).(AuthorizationInfoResponseOutput)
}

func (o AuthorizationInfoResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

type AzureDevOpsConnectorProperties struct {
	Authorization *AuthorizationInfo       `pulumi:"authorization"`
	Orgs          []AzureDevOpsOrgMetadata `pulumi:"orgs"`
}





type AzureDevOpsConnectorPropertiesInput interface {
	pulumi.Input

	ToAzureDevOpsConnectorPropertiesOutput() AzureDevOpsConnectorPropertiesOutput
	ToAzureDevOpsConnectorPropertiesOutputWithContext(context.Context) AzureDevOpsConnectorPropertiesOutput
}

type AzureDevOpsConnectorPropertiesArgs struct {
	Authorization AuthorizationInfoPtrInput        `pulumi:"authorization"`
	Orgs          AzureDevOpsOrgMetadataArrayInput `pulumi:"orgs"`
}

func (AzureDevOpsConnectorPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDevOpsConnectorProperties)(nil)).Elem()
}

func (i AzureDevOpsConnectorPropertiesArgs) ToAzureDevOpsConnectorPropertiesOutput() AzureDevOpsConnectorPropertiesOutput {
	return i.ToAzureDevOpsConnectorPropertiesOutputWithContext(context.Background())
}

func (i AzureDevOpsConnectorPropertiesArgs) ToAzureDevOpsConnectorPropertiesOutputWithContext(ctx context.Context) AzureDevOpsConnectorPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDevOpsConnectorPropertiesOutput)
}

func (i AzureDevOpsConnectorPropertiesArgs) ToAzureDevOpsConnectorPropertiesPtrOutput() AzureDevOpsConnectorPropertiesPtrOutput {
	return i.ToAzureDevOpsConnectorPropertiesPtrOutputWithContext(context.Background())
}

func (i AzureDevOpsConnectorPropertiesArgs) ToAzureDevOpsConnectorPropertiesPtrOutputWithContext(ctx context.Context) AzureDevOpsConnectorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDevOpsConnectorPropertiesOutput).ToAzureDevOpsConnectorPropertiesPtrOutputWithContext(ctx)
}









type AzureDevOpsConnectorPropertiesPtrInput interface {
	pulumi.Input

	ToAzureDevOpsConnectorPropertiesPtrOutput() AzureDevOpsConnectorPropertiesPtrOutput
	ToAzureDevOpsConnectorPropertiesPtrOutputWithContext(context.Context) AzureDevOpsConnectorPropertiesPtrOutput
}

type azureDevOpsConnectorPropertiesPtrType AzureDevOpsConnectorPropertiesArgs

func AzureDevOpsConnectorPropertiesPtr(v *AzureDevOpsConnectorPropertiesArgs) AzureDevOpsConnectorPropertiesPtrInput {
	return (*azureDevOpsConnectorPropertiesPtrType)(v)
}

func (*azureDevOpsConnectorPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureDevOpsConnectorProperties)(nil)).Elem()
}

func (i *azureDevOpsConnectorPropertiesPtrType) ToAzureDevOpsConnectorPropertiesPtrOutput() AzureDevOpsConnectorPropertiesPtrOutput {
	return i.ToAzureDevOpsConnectorPropertiesPtrOutputWithContext(context.Background())
}

func (i *azureDevOpsConnectorPropertiesPtrType) ToAzureDevOpsConnectorPropertiesPtrOutputWithContext(ctx context.Context) AzureDevOpsConnectorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDevOpsConnectorPropertiesPtrOutput)
}

type AzureDevOpsConnectorPropertiesOutput struct{ *pulumi.OutputState }

func (AzureDevOpsConnectorPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDevOpsConnectorProperties)(nil)).Elem()
}

func (o AzureDevOpsConnectorPropertiesOutput) ToAzureDevOpsConnectorPropertiesOutput() AzureDevOpsConnectorPropertiesOutput {
	return o
}

func (o AzureDevOpsConnectorPropertiesOutput) ToAzureDevOpsConnectorPropertiesOutputWithContext(ctx context.Context) AzureDevOpsConnectorPropertiesOutput {
	return o
}

func (o AzureDevOpsConnectorPropertiesOutput) ToAzureDevOpsConnectorPropertiesPtrOutput() AzureDevOpsConnectorPropertiesPtrOutput {
	return o.ToAzureDevOpsConnectorPropertiesPtrOutputWithContext(context.Background())
}

func (o AzureDevOpsConnectorPropertiesOutput) ToAzureDevOpsConnectorPropertiesPtrOutputWithContext(ctx context.Context) AzureDevOpsConnectorPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureDevOpsConnectorProperties) *AzureDevOpsConnectorProperties {
		return &v
	}).(AzureDevOpsConnectorPropertiesPtrOutput)
}

func (o AzureDevOpsConnectorPropertiesOutput) Authorization() AuthorizationInfoPtrOutput {
	return o.ApplyT(func(v AzureDevOpsConnectorProperties) *AuthorizationInfo { return v.Authorization }).(AuthorizationInfoPtrOutput)
}

func (o AzureDevOpsConnectorPropertiesOutput) Orgs() AzureDevOpsOrgMetadataArrayOutput {
	return o.ApplyT(func(v AzureDevOpsConnectorProperties) []AzureDevOpsOrgMetadata { return v.Orgs }).(AzureDevOpsOrgMetadataArrayOutput)
}

type AzureDevOpsConnectorPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AzureDevOpsConnectorPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureDevOpsConnectorProperties)(nil)).Elem()
}

func (o AzureDevOpsConnectorPropertiesPtrOutput) ToAzureDevOpsConnectorPropertiesPtrOutput() AzureDevOpsConnectorPropertiesPtrOutput {
	return o
}

func (o AzureDevOpsConnectorPropertiesPtrOutput) ToAzureDevOpsConnectorPropertiesPtrOutputWithContext(ctx context.Context) AzureDevOpsConnectorPropertiesPtrOutput {
	return o
}

func (o AzureDevOpsConnectorPropertiesPtrOutput) Elem() AzureDevOpsConnectorPropertiesOutput {
	return o.ApplyT(func(v *AzureDevOpsConnectorProperties) AzureDevOpsConnectorProperties {
		if v != nil {
			return *v
		}
		var ret AzureDevOpsConnectorProperties
		return ret
	}).(AzureDevOpsConnectorPropertiesOutput)
}

func (o AzureDevOpsConnectorPropertiesPtrOutput) Authorization() AuthorizationInfoPtrOutput {
	return o.ApplyT(func(v *AzureDevOpsConnectorProperties) *AuthorizationInfo {
		if v == nil {
			return nil
		}
		return v.Authorization
	}).(AuthorizationInfoPtrOutput)
}

func (o AzureDevOpsConnectorPropertiesPtrOutput) Orgs() AzureDevOpsOrgMetadataArrayOutput {
	return o.ApplyT(func(v *AzureDevOpsConnectorProperties) []AzureDevOpsOrgMetadata {
		if v == nil {
			return nil
		}
		return v.Orgs
	}).(AzureDevOpsOrgMetadataArrayOutput)
}

type AzureDevOpsConnectorPropertiesResponse struct {
	Authorization     *AuthorizationInfoResponse       `pulumi:"authorization"`
	Orgs              []AzureDevOpsOrgMetadataResponse `pulumi:"orgs"`
	ProvisioningState string                           `pulumi:"provisioningState"`
}

type AzureDevOpsConnectorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AzureDevOpsConnectorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDevOpsConnectorPropertiesResponse)(nil)).Elem()
}

func (o AzureDevOpsConnectorPropertiesResponseOutput) ToAzureDevOpsConnectorPropertiesResponseOutput() AzureDevOpsConnectorPropertiesResponseOutput {
	return o
}

func (o AzureDevOpsConnectorPropertiesResponseOutput) ToAzureDevOpsConnectorPropertiesResponseOutputWithContext(ctx context.Context) AzureDevOpsConnectorPropertiesResponseOutput {
	return o
}

func (o AzureDevOpsConnectorPropertiesResponseOutput) Authorization() AuthorizationInfoResponsePtrOutput {
	return o.ApplyT(func(v AzureDevOpsConnectorPropertiesResponse) *AuthorizationInfoResponse { return v.Authorization }).(AuthorizationInfoResponsePtrOutput)
}

func (o AzureDevOpsConnectorPropertiesResponseOutput) Orgs() AzureDevOpsOrgMetadataResponseArrayOutput {
	return o.ApplyT(func(v AzureDevOpsConnectorPropertiesResponse) []AzureDevOpsOrgMetadataResponse { return v.Orgs }).(AzureDevOpsOrgMetadataResponseArrayOutput)
}

func (o AzureDevOpsConnectorPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDevOpsConnectorPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type AzureDevOpsOrgMetadata struct {
	AutoDiscovery *string                      `pulumi:"autoDiscovery"`
	Name          *string                      `pulumi:"name"`
	Projects      []AzureDevOpsProjectMetadata `pulumi:"projects"`
}





type AzureDevOpsOrgMetadataInput interface {
	pulumi.Input

	ToAzureDevOpsOrgMetadataOutput() AzureDevOpsOrgMetadataOutput
	ToAzureDevOpsOrgMetadataOutputWithContext(context.Context) AzureDevOpsOrgMetadataOutput
}

type AzureDevOpsOrgMetadataArgs struct {
	AutoDiscovery pulumi.StringPtrInput                `pulumi:"autoDiscovery"`
	Name          pulumi.StringPtrInput                `pulumi:"name"`
	Projects      AzureDevOpsProjectMetadataArrayInput `pulumi:"projects"`
}

func (AzureDevOpsOrgMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDevOpsOrgMetadata)(nil)).Elem()
}

func (i AzureDevOpsOrgMetadataArgs) ToAzureDevOpsOrgMetadataOutput() AzureDevOpsOrgMetadataOutput {
	return i.ToAzureDevOpsOrgMetadataOutputWithContext(context.Background())
}

func (i AzureDevOpsOrgMetadataArgs) ToAzureDevOpsOrgMetadataOutputWithContext(ctx context.Context) AzureDevOpsOrgMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDevOpsOrgMetadataOutput)
}





type AzureDevOpsOrgMetadataArrayInput interface {
	pulumi.Input

	ToAzureDevOpsOrgMetadataArrayOutput() AzureDevOpsOrgMetadataArrayOutput
	ToAzureDevOpsOrgMetadataArrayOutputWithContext(context.Context) AzureDevOpsOrgMetadataArrayOutput
}

type AzureDevOpsOrgMetadataArray []AzureDevOpsOrgMetadataInput

func (AzureDevOpsOrgMetadataArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureDevOpsOrgMetadata)(nil)).Elem()
}

func (i AzureDevOpsOrgMetadataArray) ToAzureDevOpsOrgMetadataArrayOutput() AzureDevOpsOrgMetadataArrayOutput {
	return i.ToAzureDevOpsOrgMetadataArrayOutputWithContext(context.Background())
}

func (i AzureDevOpsOrgMetadataArray) ToAzureDevOpsOrgMetadataArrayOutputWithContext(ctx context.Context) AzureDevOpsOrgMetadataArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDevOpsOrgMetadataArrayOutput)
}

type AzureDevOpsOrgMetadataOutput struct{ *pulumi.OutputState }

func (AzureDevOpsOrgMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDevOpsOrgMetadata)(nil)).Elem()
}

func (o AzureDevOpsOrgMetadataOutput) ToAzureDevOpsOrgMetadataOutput() AzureDevOpsOrgMetadataOutput {
	return o
}

func (o AzureDevOpsOrgMetadataOutput) ToAzureDevOpsOrgMetadataOutputWithContext(ctx context.Context) AzureDevOpsOrgMetadataOutput {
	return o
}

func (o AzureDevOpsOrgMetadataOutput) AutoDiscovery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDevOpsOrgMetadata) *string { return v.AutoDiscovery }).(pulumi.StringPtrOutput)
}

func (o AzureDevOpsOrgMetadataOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDevOpsOrgMetadata) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AzureDevOpsOrgMetadataOutput) Projects() AzureDevOpsProjectMetadataArrayOutput {
	return o.ApplyT(func(v AzureDevOpsOrgMetadata) []AzureDevOpsProjectMetadata { return v.Projects }).(AzureDevOpsProjectMetadataArrayOutput)
}

type AzureDevOpsOrgMetadataArrayOutput struct{ *pulumi.OutputState }

func (AzureDevOpsOrgMetadataArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureDevOpsOrgMetadata)(nil)).Elem()
}

func (o AzureDevOpsOrgMetadataArrayOutput) ToAzureDevOpsOrgMetadataArrayOutput() AzureDevOpsOrgMetadataArrayOutput {
	return o
}

func (o AzureDevOpsOrgMetadataArrayOutput) ToAzureDevOpsOrgMetadataArrayOutputWithContext(ctx context.Context) AzureDevOpsOrgMetadataArrayOutput {
	return o
}

func (o AzureDevOpsOrgMetadataArrayOutput) Index(i pulumi.IntInput) AzureDevOpsOrgMetadataOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureDevOpsOrgMetadata {
		return vs[0].([]AzureDevOpsOrgMetadata)[vs[1].(int)]
	}).(AzureDevOpsOrgMetadataOutput)
}

type AzureDevOpsOrgMetadataResponse struct {
	AutoDiscovery *string                              `pulumi:"autoDiscovery"`
	Name          *string                              `pulumi:"name"`
	Projects      []AzureDevOpsProjectMetadataResponse `pulumi:"projects"`
}

type AzureDevOpsOrgMetadataResponseOutput struct{ *pulumi.OutputState }

func (AzureDevOpsOrgMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDevOpsOrgMetadataResponse)(nil)).Elem()
}

func (o AzureDevOpsOrgMetadataResponseOutput) ToAzureDevOpsOrgMetadataResponseOutput() AzureDevOpsOrgMetadataResponseOutput {
	return o
}

func (o AzureDevOpsOrgMetadataResponseOutput) ToAzureDevOpsOrgMetadataResponseOutputWithContext(ctx context.Context) AzureDevOpsOrgMetadataResponseOutput {
	return o
}

func (o AzureDevOpsOrgMetadataResponseOutput) AutoDiscovery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDevOpsOrgMetadataResponse) *string { return v.AutoDiscovery }).(pulumi.StringPtrOutput)
}

func (o AzureDevOpsOrgMetadataResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDevOpsOrgMetadataResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AzureDevOpsOrgMetadataResponseOutput) Projects() AzureDevOpsProjectMetadataResponseArrayOutput {
	return o.ApplyT(func(v AzureDevOpsOrgMetadataResponse) []AzureDevOpsProjectMetadataResponse { return v.Projects }).(AzureDevOpsProjectMetadataResponseArrayOutput)
}

type AzureDevOpsOrgMetadataResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureDevOpsOrgMetadataResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureDevOpsOrgMetadataResponse)(nil)).Elem()
}

func (o AzureDevOpsOrgMetadataResponseArrayOutput) ToAzureDevOpsOrgMetadataResponseArrayOutput() AzureDevOpsOrgMetadataResponseArrayOutput {
	return o
}

func (o AzureDevOpsOrgMetadataResponseArrayOutput) ToAzureDevOpsOrgMetadataResponseArrayOutputWithContext(ctx context.Context) AzureDevOpsOrgMetadataResponseArrayOutput {
	return o
}

func (o AzureDevOpsOrgMetadataResponseArrayOutput) Index(i pulumi.IntInput) AzureDevOpsOrgMetadataResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureDevOpsOrgMetadataResponse {
		return vs[0].([]AzureDevOpsOrgMetadataResponse)[vs[1].(int)]
	}).(AzureDevOpsOrgMetadataResponseOutput)
}

type AzureDevOpsProjectMetadata struct {
	AutoDiscovery *string  `pulumi:"autoDiscovery"`
	Name          *string  `pulumi:"name"`
	Repos         []string `pulumi:"repos"`
}





type AzureDevOpsProjectMetadataInput interface {
	pulumi.Input

	ToAzureDevOpsProjectMetadataOutput() AzureDevOpsProjectMetadataOutput
	ToAzureDevOpsProjectMetadataOutputWithContext(context.Context) AzureDevOpsProjectMetadataOutput
}

type AzureDevOpsProjectMetadataArgs struct {
	AutoDiscovery pulumi.StringPtrInput   `pulumi:"autoDiscovery"`
	Name          pulumi.StringPtrInput   `pulumi:"name"`
	Repos         pulumi.StringArrayInput `pulumi:"repos"`
}

func (AzureDevOpsProjectMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDevOpsProjectMetadata)(nil)).Elem()
}

func (i AzureDevOpsProjectMetadataArgs) ToAzureDevOpsProjectMetadataOutput() AzureDevOpsProjectMetadataOutput {
	return i.ToAzureDevOpsProjectMetadataOutputWithContext(context.Background())
}

func (i AzureDevOpsProjectMetadataArgs) ToAzureDevOpsProjectMetadataOutputWithContext(ctx context.Context) AzureDevOpsProjectMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDevOpsProjectMetadataOutput)
}





type AzureDevOpsProjectMetadataArrayInput interface {
	pulumi.Input

	ToAzureDevOpsProjectMetadataArrayOutput() AzureDevOpsProjectMetadataArrayOutput
	ToAzureDevOpsProjectMetadataArrayOutputWithContext(context.Context) AzureDevOpsProjectMetadataArrayOutput
}

type AzureDevOpsProjectMetadataArray []AzureDevOpsProjectMetadataInput

func (AzureDevOpsProjectMetadataArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureDevOpsProjectMetadata)(nil)).Elem()
}

func (i AzureDevOpsProjectMetadataArray) ToAzureDevOpsProjectMetadataArrayOutput() AzureDevOpsProjectMetadataArrayOutput {
	return i.ToAzureDevOpsProjectMetadataArrayOutputWithContext(context.Background())
}

func (i AzureDevOpsProjectMetadataArray) ToAzureDevOpsProjectMetadataArrayOutputWithContext(ctx context.Context) AzureDevOpsProjectMetadataArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDevOpsProjectMetadataArrayOutput)
}

type AzureDevOpsProjectMetadataOutput struct{ *pulumi.OutputState }

func (AzureDevOpsProjectMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDevOpsProjectMetadata)(nil)).Elem()
}

func (o AzureDevOpsProjectMetadataOutput) ToAzureDevOpsProjectMetadataOutput() AzureDevOpsProjectMetadataOutput {
	return o
}

func (o AzureDevOpsProjectMetadataOutput) ToAzureDevOpsProjectMetadataOutputWithContext(ctx context.Context) AzureDevOpsProjectMetadataOutput {
	return o
}

func (o AzureDevOpsProjectMetadataOutput) AutoDiscovery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDevOpsProjectMetadata) *string { return v.AutoDiscovery }).(pulumi.StringPtrOutput)
}

func (o AzureDevOpsProjectMetadataOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDevOpsProjectMetadata) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AzureDevOpsProjectMetadataOutput) Repos() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureDevOpsProjectMetadata) []string { return v.Repos }).(pulumi.StringArrayOutput)
}

type AzureDevOpsProjectMetadataArrayOutput struct{ *pulumi.OutputState }

func (AzureDevOpsProjectMetadataArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureDevOpsProjectMetadata)(nil)).Elem()
}

func (o AzureDevOpsProjectMetadataArrayOutput) ToAzureDevOpsProjectMetadataArrayOutput() AzureDevOpsProjectMetadataArrayOutput {
	return o
}

func (o AzureDevOpsProjectMetadataArrayOutput) ToAzureDevOpsProjectMetadataArrayOutputWithContext(ctx context.Context) AzureDevOpsProjectMetadataArrayOutput {
	return o
}

func (o AzureDevOpsProjectMetadataArrayOutput) Index(i pulumi.IntInput) AzureDevOpsProjectMetadataOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureDevOpsProjectMetadata {
		return vs[0].([]AzureDevOpsProjectMetadata)[vs[1].(int)]
	}).(AzureDevOpsProjectMetadataOutput)
}

type AzureDevOpsProjectMetadataResponse struct {
	AutoDiscovery *string  `pulumi:"autoDiscovery"`
	Name          *string  `pulumi:"name"`
	Repos         []string `pulumi:"repos"`
}

type AzureDevOpsProjectMetadataResponseOutput struct{ *pulumi.OutputState }

func (AzureDevOpsProjectMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDevOpsProjectMetadataResponse)(nil)).Elem()
}

func (o AzureDevOpsProjectMetadataResponseOutput) ToAzureDevOpsProjectMetadataResponseOutput() AzureDevOpsProjectMetadataResponseOutput {
	return o
}

func (o AzureDevOpsProjectMetadataResponseOutput) ToAzureDevOpsProjectMetadataResponseOutputWithContext(ctx context.Context) AzureDevOpsProjectMetadataResponseOutput {
	return o
}

func (o AzureDevOpsProjectMetadataResponseOutput) AutoDiscovery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDevOpsProjectMetadataResponse) *string { return v.AutoDiscovery }).(pulumi.StringPtrOutput)
}

func (o AzureDevOpsProjectMetadataResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDevOpsProjectMetadataResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AzureDevOpsProjectMetadataResponseOutput) Repos() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureDevOpsProjectMetadataResponse) []string { return v.Repos }).(pulumi.StringArrayOutput)
}

type AzureDevOpsProjectMetadataResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureDevOpsProjectMetadataResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureDevOpsProjectMetadataResponse)(nil)).Elem()
}

func (o AzureDevOpsProjectMetadataResponseArrayOutput) ToAzureDevOpsProjectMetadataResponseArrayOutput() AzureDevOpsProjectMetadataResponseArrayOutput {
	return o
}

func (o AzureDevOpsProjectMetadataResponseArrayOutput) ToAzureDevOpsProjectMetadataResponseArrayOutputWithContext(ctx context.Context) AzureDevOpsProjectMetadataResponseArrayOutput {
	return o
}

func (o AzureDevOpsProjectMetadataResponseArrayOutput) Index(i pulumi.IntInput) AzureDevOpsProjectMetadataResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureDevOpsProjectMetadataResponse {
		return vs[0].([]AzureDevOpsProjectMetadataResponse)[vs[1].(int)]
	}).(AzureDevOpsProjectMetadataResponseOutput)
}

type GitHubConnectorProperties struct {
	Code *string `pulumi:"code"`
}





type GitHubConnectorPropertiesInput interface {
	pulumi.Input

	ToGitHubConnectorPropertiesOutput() GitHubConnectorPropertiesOutput
	ToGitHubConnectorPropertiesOutputWithContext(context.Context) GitHubConnectorPropertiesOutput
}

type GitHubConnectorPropertiesArgs struct {
	Code pulumi.StringPtrInput `pulumi:"code"`
}

func (GitHubConnectorPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubConnectorProperties)(nil)).Elem()
}

func (i GitHubConnectorPropertiesArgs) ToGitHubConnectorPropertiesOutput() GitHubConnectorPropertiesOutput {
	return i.ToGitHubConnectorPropertiesOutputWithContext(context.Background())
}

func (i GitHubConnectorPropertiesArgs) ToGitHubConnectorPropertiesOutputWithContext(ctx context.Context) GitHubConnectorPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubConnectorPropertiesOutput)
}

func (i GitHubConnectorPropertiesArgs) ToGitHubConnectorPropertiesPtrOutput() GitHubConnectorPropertiesPtrOutput {
	return i.ToGitHubConnectorPropertiesPtrOutputWithContext(context.Background())
}

func (i GitHubConnectorPropertiesArgs) ToGitHubConnectorPropertiesPtrOutputWithContext(ctx context.Context) GitHubConnectorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubConnectorPropertiesOutput).ToGitHubConnectorPropertiesPtrOutputWithContext(ctx)
}









type GitHubConnectorPropertiesPtrInput interface {
	pulumi.Input

	ToGitHubConnectorPropertiesPtrOutput() GitHubConnectorPropertiesPtrOutput
	ToGitHubConnectorPropertiesPtrOutputWithContext(context.Context) GitHubConnectorPropertiesPtrOutput
}

type gitHubConnectorPropertiesPtrType GitHubConnectorPropertiesArgs

func GitHubConnectorPropertiesPtr(v *GitHubConnectorPropertiesArgs) GitHubConnectorPropertiesPtrInput {
	return (*gitHubConnectorPropertiesPtrType)(v)
}

func (*gitHubConnectorPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubConnectorProperties)(nil)).Elem()
}

func (i *gitHubConnectorPropertiesPtrType) ToGitHubConnectorPropertiesPtrOutput() GitHubConnectorPropertiesPtrOutput {
	return i.ToGitHubConnectorPropertiesPtrOutputWithContext(context.Background())
}

func (i *gitHubConnectorPropertiesPtrType) ToGitHubConnectorPropertiesPtrOutputWithContext(ctx context.Context) GitHubConnectorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubConnectorPropertiesPtrOutput)
}

type GitHubConnectorPropertiesOutput struct{ *pulumi.OutputState }

func (GitHubConnectorPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubConnectorProperties)(nil)).Elem()
}

func (o GitHubConnectorPropertiesOutput) ToGitHubConnectorPropertiesOutput() GitHubConnectorPropertiesOutput {
	return o
}

func (o GitHubConnectorPropertiesOutput) ToGitHubConnectorPropertiesOutputWithContext(ctx context.Context) GitHubConnectorPropertiesOutput {
	return o
}

func (o GitHubConnectorPropertiesOutput) ToGitHubConnectorPropertiesPtrOutput() GitHubConnectorPropertiesPtrOutput {
	return o.ToGitHubConnectorPropertiesPtrOutputWithContext(context.Background())
}

func (o GitHubConnectorPropertiesOutput) ToGitHubConnectorPropertiesPtrOutputWithContext(ctx context.Context) GitHubConnectorPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitHubConnectorProperties) *GitHubConnectorProperties {
		return &v
	}).(GitHubConnectorPropertiesPtrOutput)
}

func (o GitHubConnectorPropertiesOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubConnectorProperties) *string { return v.Code }).(pulumi.StringPtrOutput)
}

type GitHubConnectorPropertiesPtrOutput struct{ *pulumi.OutputState }

func (GitHubConnectorPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubConnectorProperties)(nil)).Elem()
}

func (o GitHubConnectorPropertiesPtrOutput) ToGitHubConnectorPropertiesPtrOutput() GitHubConnectorPropertiesPtrOutput {
	return o
}

func (o GitHubConnectorPropertiesPtrOutput) ToGitHubConnectorPropertiesPtrOutputWithContext(ctx context.Context) GitHubConnectorPropertiesPtrOutput {
	return o
}

func (o GitHubConnectorPropertiesPtrOutput) Elem() GitHubConnectorPropertiesOutput {
	return o.ApplyT(func(v *GitHubConnectorProperties) GitHubConnectorProperties {
		if v != nil {
			return *v
		}
		var ret GitHubConnectorProperties
		return ret
	}).(GitHubConnectorPropertiesOutput)
}

func (o GitHubConnectorPropertiesPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubConnectorProperties) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

type GitHubConnectorPropertiesResponse struct {
	Code              *string `pulumi:"code"`
	ProvisioningState string  `pulumi:"provisioningState"`
}

type GitHubConnectorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (GitHubConnectorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubConnectorPropertiesResponse)(nil)).Elem()
}

func (o GitHubConnectorPropertiesResponseOutput) ToGitHubConnectorPropertiesResponseOutput() GitHubConnectorPropertiesResponseOutput {
	return o
}

func (o GitHubConnectorPropertiesResponseOutput) ToGitHubConnectorPropertiesResponseOutputWithContext(ctx context.Context) GitHubConnectorPropertiesResponseOutput {
	return o
}

func (o GitHubConnectorPropertiesResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubConnectorPropertiesResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o GitHubConnectorPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GitHubConnectorPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
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

func init() {
	pulumi.RegisterOutputType(AuthorizationInfoOutput{})
	pulumi.RegisterOutputType(AuthorizationInfoPtrOutput{})
	pulumi.RegisterOutputType(AuthorizationInfoResponseOutput{})
	pulumi.RegisterOutputType(AuthorizationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureDevOpsConnectorPropertiesOutput{})
	pulumi.RegisterOutputType(AzureDevOpsConnectorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AzureDevOpsConnectorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AzureDevOpsOrgMetadataOutput{})
	pulumi.RegisterOutputType(AzureDevOpsOrgMetadataArrayOutput{})
	pulumi.RegisterOutputType(AzureDevOpsOrgMetadataResponseOutput{})
	pulumi.RegisterOutputType(AzureDevOpsOrgMetadataResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureDevOpsProjectMetadataOutput{})
	pulumi.RegisterOutputType(AzureDevOpsProjectMetadataArrayOutput{})
	pulumi.RegisterOutputType(AzureDevOpsProjectMetadataResponseOutput{})
	pulumi.RegisterOutputType(AzureDevOpsProjectMetadataResponseArrayOutput{})
	pulumi.RegisterOutputType(GitHubConnectorPropertiesOutput{})
	pulumi.RegisterOutputType(GitHubConnectorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(GitHubConnectorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}




package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthorizationProfileResponse struct {
	ApprovedTime      string `pulumi:"approvedTime"`
	Approver          string `pulumi:"approver"`
	RequestedTime     string `pulumi:"requestedTime"`
	Requester         string `pulumi:"requester"`
	RequesterObjectId string `pulumi:"requesterObjectId"`
}





type AuthorizationProfileResponseInput interface {
	pulumi.Input

	ToAuthorizationProfileResponseOutput() AuthorizationProfileResponseOutput
	ToAuthorizationProfileResponseOutputWithContext(context.Context) AuthorizationProfileResponseOutput
}

type AuthorizationProfileResponseArgs struct {
	ApprovedTime      pulumi.StringInput `pulumi:"approvedTime"`
	Approver          pulumi.StringInput `pulumi:"approver"`
	RequestedTime     pulumi.StringInput `pulumi:"requestedTime"`
	Requester         pulumi.StringInput `pulumi:"requester"`
	RequesterObjectId pulumi.StringInput `pulumi:"requesterObjectId"`
}

func (AuthorizationProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationProfileResponse)(nil)).Elem()
}

func (i AuthorizationProfileResponseArgs) ToAuthorizationProfileResponseOutput() AuthorizationProfileResponseOutput {
	return i.ToAuthorizationProfileResponseOutputWithContext(context.Background())
}

func (i AuthorizationProfileResponseArgs) ToAuthorizationProfileResponseOutputWithContext(ctx context.Context) AuthorizationProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationProfileResponseOutput)
}

func (i AuthorizationProfileResponseArgs) ToAuthorizationProfileResponsePtrOutput() AuthorizationProfileResponsePtrOutput {
	return i.ToAuthorizationProfileResponsePtrOutputWithContext(context.Background())
}

func (i AuthorizationProfileResponseArgs) ToAuthorizationProfileResponsePtrOutputWithContext(ctx context.Context) AuthorizationProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationProfileResponseOutput).ToAuthorizationProfileResponsePtrOutputWithContext(ctx)
}









type AuthorizationProfileResponsePtrInput interface {
	pulumi.Input

	ToAuthorizationProfileResponsePtrOutput() AuthorizationProfileResponsePtrOutput
	ToAuthorizationProfileResponsePtrOutputWithContext(context.Context) AuthorizationProfileResponsePtrOutput
}

type authorizationProfileResponsePtrType AuthorizationProfileResponseArgs

func AuthorizationProfileResponsePtr(v *AuthorizationProfileResponseArgs) AuthorizationProfileResponsePtrInput {
	return (*authorizationProfileResponsePtrType)(v)
}

func (*authorizationProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationProfileResponse)(nil)).Elem()
}

func (i *authorizationProfileResponsePtrType) ToAuthorizationProfileResponsePtrOutput() AuthorizationProfileResponsePtrOutput {
	return i.ToAuthorizationProfileResponsePtrOutputWithContext(context.Background())
}

func (i *authorizationProfileResponsePtrType) ToAuthorizationProfileResponsePtrOutputWithContext(ctx context.Context) AuthorizationProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationProfileResponsePtrOutput)
}

type AuthorizationProfileResponseOutput struct{ *pulumi.OutputState }

func (AuthorizationProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationProfileResponse)(nil)).Elem()
}

func (o AuthorizationProfileResponseOutput) ToAuthorizationProfileResponseOutput() AuthorizationProfileResponseOutput {
	return o
}

func (o AuthorizationProfileResponseOutput) ToAuthorizationProfileResponseOutputWithContext(ctx context.Context) AuthorizationProfileResponseOutput {
	return o
}

func (o AuthorizationProfileResponseOutput) ToAuthorizationProfileResponsePtrOutput() AuthorizationProfileResponsePtrOutput {
	return o.ToAuthorizationProfileResponsePtrOutputWithContext(context.Background())
}

func (o AuthorizationProfileResponseOutput) ToAuthorizationProfileResponsePtrOutputWithContext(ctx context.Context) AuthorizationProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthorizationProfileResponse) *AuthorizationProfileResponse {
		return &v
	}).(AuthorizationProfileResponsePtrOutput)
}

func (o AuthorizationProfileResponseOutput) ApprovedTime() pulumi.StringOutput {
	return o.ApplyT(func(v AuthorizationProfileResponse) string { return v.ApprovedTime }).(pulumi.StringOutput)
}

func (o AuthorizationProfileResponseOutput) Approver() pulumi.StringOutput {
	return o.ApplyT(func(v AuthorizationProfileResponse) string { return v.Approver }).(pulumi.StringOutput)
}

func (o AuthorizationProfileResponseOutput) RequestedTime() pulumi.StringOutput {
	return o.ApplyT(func(v AuthorizationProfileResponse) string { return v.RequestedTime }).(pulumi.StringOutput)
}

func (o AuthorizationProfileResponseOutput) Requester() pulumi.StringOutput {
	return o.ApplyT(func(v AuthorizationProfileResponse) string { return v.Requester }).(pulumi.StringOutput)
}

func (o AuthorizationProfileResponseOutput) RequesterObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v AuthorizationProfileResponse) string { return v.RequesterObjectId }).(pulumi.StringOutput)
}

type AuthorizationProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (AuthorizationProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationProfileResponse)(nil)).Elem()
}

func (o AuthorizationProfileResponsePtrOutput) ToAuthorizationProfileResponsePtrOutput() AuthorizationProfileResponsePtrOutput {
	return o
}

func (o AuthorizationProfileResponsePtrOutput) ToAuthorizationProfileResponsePtrOutputWithContext(ctx context.Context) AuthorizationProfileResponsePtrOutput {
	return o
}

func (o AuthorizationProfileResponsePtrOutput) Elem() AuthorizationProfileResponseOutput {
	return o.ApplyT(func(v *AuthorizationProfileResponse) AuthorizationProfileResponse {
		if v != nil {
			return *v
		}
		var ret AuthorizationProfileResponse
		return ret
	}).(AuthorizationProfileResponseOutput)
}

func (o AuthorizationProfileResponsePtrOutput) ApprovedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ApprovedTime
	}).(pulumi.StringPtrOutput)
}

func (o AuthorizationProfileResponsePtrOutput) Approver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Approver
	}).(pulumi.StringPtrOutput)
}

func (o AuthorizationProfileResponsePtrOutput) RequestedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RequestedTime
	}).(pulumi.StringPtrOutput)
}

func (o AuthorizationProfileResponsePtrOutput) Requester() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Requester
	}).(pulumi.StringPtrOutput)
}

func (o AuthorizationProfileResponsePtrOutput) RequesterObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RequesterObjectId
	}).(pulumi.StringPtrOutput)
}

type SubscriptionFeatureRegistrationProperties struct {
	Description                  *string           `pulumi:"description"`
	Metadata                     map[string]string `pulumi:"metadata"`
	ShouldFeatureDisplayInPortal *bool             `pulumi:"shouldFeatureDisplayInPortal"`
	State                        *string           `pulumi:"state"`
}


func (val *SubscriptionFeatureRegistrationProperties) Defaults() *SubscriptionFeatureRegistrationProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ShouldFeatureDisplayInPortal) {
		shouldFeatureDisplayInPortal_ := false
		tmp.ShouldFeatureDisplayInPortal = &shouldFeatureDisplayInPortal_
	}
	return &tmp
}





type SubscriptionFeatureRegistrationPropertiesInput interface {
	pulumi.Input

	ToSubscriptionFeatureRegistrationPropertiesOutput() SubscriptionFeatureRegistrationPropertiesOutput
	ToSubscriptionFeatureRegistrationPropertiesOutputWithContext(context.Context) SubscriptionFeatureRegistrationPropertiesOutput
}

type SubscriptionFeatureRegistrationPropertiesArgs struct {
	Description                  pulumi.StringPtrInput `pulumi:"description"`
	Metadata                     pulumi.StringMapInput `pulumi:"metadata"`
	ShouldFeatureDisplayInPortal pulumi.BoolPtrInput   `pulumi:"shouldFeatureDisplayInPortal"`
	State                        pulumi.StringPtrInput `pulumi:"state"`
}

func (SubscriptionFeatureRegistrationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionFeatureRegistrationProperties)(nil)).Elem()
}

func (i SubscriptionFeatureRegistrationPropertiesArgs) ToSubscriptionFeatureRegistrationPropertiesOutput() SubscriptionFeatureRegistrationPropertiesOutput {
	return i.ToSubscriptionFeatureRegistrationPropertiesOutputWithContext(context.Background())
}

func (i SubscriptionFeatureRegistrationPropertiesArgs) ToSubscriptionFeatureRegistrationPropertiesOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionFeatureRegistrationPropertiesOutput)
}

func (i SubscriptionFeatureRegistrationPropertiesArgs) ToSubscriptionFeatureRegistrationPropertiesPtrOutput() SubscriptionFeatureRegistrationPropertiesPtrOutput {
	return i.ToSubscriptionFeatureRegistrationPropertiesPtrOutputWithContext(context.Background())
}

func (i SubscriptionFeatureRegistrationPropertiesArgs) ToSubscriptionFeatureRegistrationPropertiesPtrOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionFeatureRegistrationPropertiesOutput).ToSubscriptionFeatureRegistrationPropertiesPtrOutputWithContext(ctx)
}









type SubscriptionFeatureRegistrationPropertiesPtrInput interface {
	pulumi.Input

	ToSubscriptionFeatureRegistrationPropertiesPtrOutput() SubscriptionFeatureRegistrationPropertiesPtrOutput
	ToSubscriptionFeatureRegistrationPropertiesPtrOutputWithContext(context.Context) SubscriptionFeatureRegistrationPropertiesPtrOutput
}

type subscriptionFeatureRegistrationPropertiesPtrType SubscriptionFeatureRegistrationPropertiesArgs

func SubscriptionFeatureRegistrationPropertiesPtr(v *SubscriptionFeatureRegistrationPropertiesArgs) SubscriptionFeatureRegistrationPropertiesPtrInput {
	return (*subscriptionFeatureRegistrationPropertiesPtrType)(v)
}

func (*subscriptionFeatureRegistrationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionFeatureRegistrationProperties)(nil)).Elem()
}

func (i *subscriptionFeatureRegistrationPropertiesPtrType) ToSubscriptionFeatureRegistrationPropertiesPtrOutput() SubscriptionFeatureRegistrationPropertiesPtrOutput {
	return i.ToSubscriptionFeatureRegistrationPropertiesPtrOutputWithContext(context.Background())
}

func (i *subscriptionFeatureRegistrationPropertiesPtrType) ToSubscriptionFeatureRegistrationPropertiesPtrOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionFeatureRegistrationPropertiesPtrOutput)
}

type SubscriptionFeatureRegistrationPropertiesOutput struct{ *pulumi.OutputState }

func (SubscriptionFeatureRegistrationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionFeatureRegistrationProperties)(nil)).Elem()
}

func (o SubscriptionFeatureRegistrationPropertiesOutput) ToSubscriptionFeatureRegistrationPropertiesOutput() SubscriptionFeatureRegistrationPropertiesOutput {
	return o
}

func (o SubscriptionFeatureRegistrationPropertiesOutput) ToSubscriptionFeatureRegistrationPropertiesOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationPropertiesOutput {
	return o
}

func (o SubscriptionFeatureRegistrationPropertiesOutput) ToSubscriptionFeatureRegistrationPropertiesPtrOutput() SubscriptionFeatureRegistrationPropertiesPtrOutput {
	return o.ToSubscriptionFeatureRegistrationPropertiesPtrOutputWithContext(context.Background())
}

func (o SubscriptionFeatureRegistrationPropertiesOutput) ToSubscriptionFeatureRegistrationPropertiesPtrOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubscriptionFeatureRegistrationProperties) *SubscriptionFeatureRegistrationProperties {
		return &v
	}).(SubscriptionFeatureRegistrationPropertiesPtrOutput)
}

func (o SubscriptionFeatureRegistrationPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SubscriptionFeatureRegistrationPropertiesOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationProperties) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o SubscriptionFeatureRegistrationPropertiesOutput) ShouldFeatureDisplayInPortal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationProperties) *bool { return v.ShouldFeatureDisplayInPortal }).(pulumi.BoolPtrOutput)
}

func (o SubscriptionFeatureRegistrationPropertiesOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationProperties) *string { return v.State }).(pulumi.StringPtrOutput)
}

type SubscriptionFeatureRegistrationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SubscriptionFeatureRegistrationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionFeatureRegistrationProperties)(nil)).Elem()
}

func (o SubscriptionFeatureRegistrationPropertiesPtrOutput) ToSubscriptionFeatureRegistrationPropertiesPtrOutput() SubscriptionFeatureRegistrationPropertiesPtrOutput {
	return o
}

func (o SubscriptionFeatureRegistrationPropertiesPtrOutput) ToSubscriptionFeatureRegistrationPropertiesPtrOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationPropertiesPtrOutput {
	return o
}

func (o SubscriptionFeatureRegistrationPropertiesPtrOutput) Elem() SubscriptionFeatureRegistrationPropertiesOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationProperties) SubscriptionFeatureRegistrationProperties {
		if v != nil {
			return *v
		}
		var ret SubscriptionFeatureRegistrationProperties
		return ret
	}).(SubscriptionFeatureRegistrationPropertiesOutput)
}

func (o SubscriptionFeatureRegistrationPropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionFeatureRegistrationPropertiesPtrOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.StringMapOutput)
}

func (o SubscriptionFeatureRegistrationPropertiesPtrOutput) ShouldFeatureDisplayInPortal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationProperties) *bool {
		if v == nil {
			return nil
		}
		return v.ShouldFeatureDisplayInPortal
	}).(pulumi.BoolPtrOutput)
}

func (o SubscriptionFeatureRegistrationPropertiesPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationProperties) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type SubscriptionFeatureRegistrationResponseProperties struct {
	ApprovalType                 string                        `pulumi:"approvalType"`
	AuthorizationProfile         *AuthorizationProfileResponse `pulumi:"authorizationProfile"`
	Description                  *string                       `pulumi:"description"`
	DisplayName                  string                        `pulumi:"displayName"`
	DocumentationLink            string                        `pulumi:"documentationLink"`
	FeatureName                  string                        `pulumi:"featureName"`
	Metadata                     map[string]string             `pulumi:"metadata"`
	ProviderNamespace            string                        `pulumi:"providerNamespace"`
	RegistrationDate             string                        `pulumi:"registrationDate"`
	ReleaseDate                  string                        `pulumi:"releaseDate"`
	ShouldFeatureDisplayInPortal *bool                         `pulumi:"shouldFeatureDisplayInPortal"`
	State                        *string                       `pulumi:"state"`
	SubscriptionId               string                        `pulumi:"subscriptionId"`
	TenantId                     string                        `pulumi:"tenantId"`
}


func (val *SubscriptionFeatureRegistrationResponseProperties) Defaults() *SubscriptionFeatureRegistrationResponseProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ShouldFeatureDisplayInPortal) {
		shouldFeatureDisplayInPortal_ := false
		tmp.ShouldFeatureDisplayInPortal = &shouldFeatureDisplayInPortal_
	}
	return &tmp
}





type SubscriptionFeatureRegistrationResponsePropertiesInput interface {
	pulumi.Input

	ToSubscriptionFeatureRegistrationResponsePropertiesOutput() SubscriptionFeatureRegistrationResponsePropertiesOutput
	ToSubscriptionFeatureRegistrationResponsePropertiesOutputWithContext(context.Context) SubscriptionFeatureRegistrationResponsePropertiesOutput
}

type SubscriptionFeatureRegistrationResponsePropertiesArgs struct {
	ApprovalType                 pulumi.StringInput                   `pulumi:"approvalType"`
	AuthorizationProfile         AuthorizationProfileResponsePtrInput `pulumi:"authorizationProfile"`
	Description                  pulumi.StringPtrInput                `pulumi:"description"`
	DisplayName                  pulumi.StringInput                   `pulumi:"displayName"`
	DocumentationLink            pulumi.StringInput                   `pulumi:"documentationLink"`
	FeatureName                  pulumi.StringInput                   `pulumi:"featureName"`
	Metadata                     pulumi.StringMapInput                `pulumi:"metadata"`
	ProviderNamespace            pulumi.StringInput                   `pulumi:"providerNamespace"`
	RegistrationDate             pulumi.StringInput                   `pulumi:"registrationDate"`
	ReleaseDate                  pulumi.StringInput                   `pulumi:"releaseDate"`
	ShouldFeatureDisplayInPortal pulumi.BoolPtrInput                  `pulumi:"shouldFeatureDisplayInPortal"`
	State                        pulumi.StringPtrInput                `pulumi:"state"`
	SubscriptionId               pulumi.StringInput                   `pulumi:"subscriptionId"`
	TenantId                     pulumi.StringInput                   `pulumi:"tenantId"`
}

func (SubscriptionFeatureRegistrationResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionFeatureRegistrationResponseProperties)(nil)).Elem()
}

func (i SubscriptionFeatureRegistrationResponsePropertiesArgs) ToSubscriptionFeatureRegistrationResponsePropertiesOutput() SubscriptionFeatureRegistrationResponsePropertiesOutput {
	return i.ToSubscriptionFeatureRegistrationResponsePropertiesOutputWithContext(context.Background())
}

func (i SubscriptionFeatureRegistrationResponsePropertiesArgs) ToSubscriptionFeatureRegistrationResponsePropertiesOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionFeatureRegistrationResponsePropertiesOutput)
}

func (i SubscriptionFeatureRegistrationResponsePropertiesArgs) ToSubscriptionFeatureRegistrationResponsePropertiesPtrOutput() SubscriptionFeatureRegistrationResponsePropertiesPtrOutput {
	return i.ToSubscriptionFeatureRegistrationResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i SubscriptionFeatureRegistrationResponsePropertiesArgs) ToSubscriptionFeatureRegistrationResponsePropertiesPtrOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionFeatureRegistrationResponsePropertiesOutput).ToSubscriptionFeatureRegistrationResponsePropertiesPtrOutputWithContext(ctx)
}









type SubscriptionFeatureRegistrationResponsePropertiesPtrInput interface {
	pulumi.Input

	ToSubscriptionFeatureRegistrationResponsePropertiesPtrOutput() SubscriptionFeatureRegistrationResponsePropertiesPtrOutput
	ToSubscriptionFeatureRegistrationResponsePropertiesPtrOutputWithContext(context.Context) SubscriptionFeatureRegistrationResponsePropertiesPtrOutput
}

type subscriptionFeatureRegistrationResponsePropertiesPtrType SubscriptionFeatureRegistrationResponsePropertiesArgs

func SubscriptionFeatureRegistrationResponsePropertiesPtr(v *SubscriptionFeatureRegistrationResponsePropertiesArgs) SubscriptionFeatureRegistrationResponsePropertiesPtrInput {
	return (*subscriptionFeatureRegistrationResponsePropertiesPtrType)(v)
}

func (*subscriptionFeatureRegistrationResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionFeatureRegistrationResponseProperties)(nil)).Elem()
}

func (i *subscriptionFeatureRegistrationResponsePropertiesPtrType) ToSubscriptionFeatureRegistrationResponsePropertiesPtrOutput() SubscriptionFeatureRegistrationResponsePropertiesPtrOutput {
	return i.ToSubscriptionFeatureRegistrationResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *subscriptionFeatureRegistrationResponsePropertiesPtrType) ToSubscriptionFeatureRegistrationResponsePropertiesPtrOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionFeatureRegistrationResponsePropertiesPtrOutput)
}

type SubscriptionFeatureRegistrationResponsePropertiesOutput struct{ *pulumi.OutputState }

func (SubscriptionFeatureRegistrationResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionFeatureRegistrationResponseProperties)(nil)).Elem()
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) ToSubscriptionFeatureRegistrationResponsePropertiesOutput() SubscriptionFeatureRegistrationResponsePropertiesOutput {
	return o
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) ToSubscriptionFeatureRegistrationResponsePropertiesOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationResponsePropertiesOutput {
	return o
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) ToSubscriptionFeatureRegistrationResponsePropertiesPtrOutput() SubscriptionFeatureRegistrationResponsePropertiesPtrOutput {
	return o.ToSubscriptionFeatureRegistrationResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) ToSubscriptionFeatureRegistrationResponsePropertiesPtrOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubscriptionFeatureRegistrationResponseProperties) *SubscriptionFeatureRegistrationResponseProperties {
		return &v
	}).(SubscriptionFeatureRegistrationResponsePropertiesPtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) ApprovalType() pulumi.StringOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationResponseProperties) string { return v.ApprovalType }).(pulumi.StringOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) AuthorizationProfile() AuthorizationProfileResponsePtrOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationResponseProperties) *AuthorizationProfileResponse {
		return v.AuthorizationProfile
	}).(AuthorizationProfileResponsePtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationResponseProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationResponseProperties) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) DocumentationLink() pulumi.StringOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationResponseProperties) string { return v.DocumentationLink }).(pulumi.StringOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) FeatureName() pulumi.StringOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationResponseProperties) string { return v.FeatureName }).(pulumi.StringOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationResponseProperties) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) ProviderNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationResponseProperties) string { return v.ProviderNamespace }).(pulumi.StringOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) RegistrationDate() pulumi.StringOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationResponseProperties) string { return v.RegistrationDate }).(pulumi.StringOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) ReleaseDate() pulumi.StringOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationResponseProperties) string { return v.ReleaseDate }).(pulumi.StringOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) ShouldFeatureDisplayInPortal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationResponseProperties) *bool { return v.ShouldFeatureDisplayInPortal }).(pulumi.BoolPtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationResponseProperties) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationResponseProperties) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v SubscriptionFeatureRegistrationResponseProperties) string { return v.TenantId }).(pulumi.StringOutput)
}

type SubscriptionFeatureRegistrationResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionFeatureRegistrationResponseProperties)(nil)).Elem()
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) ToSubscriptionFeatureRegistrationResponsePropertiesPtrOutput() SubscriptionFeatureRegistrationResponsePropertiesPtrOutput {
	return o
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) ToSubscriptionFeatureRegistrationResponsePropertiesPtrOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationResponsePropertiesPtrOutput {
	return o
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) Elem() SubscriptionFeatureRegistrationResponsePropertiesOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationResponseProperties) SubscriptionFeatureRegistrationResponseProperties {
		if v != nil {
			return *v
		}
		var ret SubscriptionFeatureRegistrationResponseProperties
		return ret
	}).(SubscriptionFeatureRegistrationResponsePropertiesOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) ApprovalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ApprovalType
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) AuthorizationProfile() AuthorizationProfileResponsePtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationResponseProperties) *AuthorizationProfileResponse {
		if v == nil {
			return nil
		}
		return v.AuthorizationProfile
	}).(AuthorizationProfileResponsePtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) DocumentationLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.DocumentationLink
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) FeatureName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.FeatureName
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationResponseProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.StringMapOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) ProviderNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ProviderNamespace
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) RegistrationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.RegistrationDate
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) ReleaseDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ReleaseDate
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) ShouldFeatureDisplayInPortal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationResponseProperties) *bool {
		if v == nil {
			return nil
		}
		return v.ShouldFeatureDisplayInPortal
	}).(pulumi.BoolPtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionFeatureRegistrationResponsePropertiesPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationProfileResponseOutput{})
	pulumi.RegisterOutputType(AuthorizationProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SubscriptionFeatureRegistrationPropertiesOutput{})
	pulumi.RegisterOutputType(SubscriptionFeatureRegistrationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SubscriptionFeatureRegistrationResponsePropertiesOutput{})
	pulumi.RegisterOutputType(SubscriptionFeatureRegistrationResponsePropertiesPtrOutput{})
}

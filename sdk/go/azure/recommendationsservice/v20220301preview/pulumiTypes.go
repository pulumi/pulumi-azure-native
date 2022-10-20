


package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccountResourceProperties struct {
	Configuration           *string                  `pulumi:"configuration"`
	Cors                    []CorsRule               `pulumi:"cors"`
	EndpointAuthentications []EndpointAuthentication `pulumi:"endpointAuthentications"`
	ReportsConnectionString *string                  `pulumi:"reportsConnectionString"`
}





type AccountResourcePropertiesInput interface {
	pulumi.Input

	ToAccountResourcePropertiesOutput() AccountResourcePropertiesOutput
	ToAccountResourcePropertiesOutputWithContext(context.Context) AccountResourcePropertiesOutput
}

type AccountResourcePropertiesArgs struct {
	Configuration           pulumi.StringPtrInput            `pulumi:"configuration"`
	Cors                    CorsRuleArrayInput               `pulumi:"cors"`
	EndpointAuthentications EndpointAuthenticationArrayInput `pulumi:"endpointAuthentications"`
	ReportsConnectionString pulumi.StringPtrInput            `pulumi:"reportsConnectionString"`
}

func (AccountResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountResourceProperties)(nil)).Elem()
}

func (i AccountResourcePropertiesArgs) ToAccountResourcePropertiesOutput() AccountResourcePropertiesOutput {
	return i.ToAccountResourcePropertiesOutputWithContext(context.Background())
}

func (i AccountResourcePropertiesArgs) ToAccountResourcePropertiesOutputWithContext(ctx context.Context) AccountResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountResourcePropertiesOutput)
}

func (i AccountResourcePropertiesArgs) ToAccountResourcePropertiesPtrOutput() AccountResourcePropertiesPtrOutput {
	return i.ToAccountResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i AccountResourcePropertiesArgs) ToAccountResourcePropertiesPtrOutputWithContext(ctx context.Context) AccountResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountResourcePropertiesOutput).ToAccountResourcePropertiesPtrOutputWithContext(ctx)
}









type AccountResourcePropertiesPtrInput interface {
	pulumi.Input

	ToAccountResourcePropertiesPtrOutput() AccountResourcePropertiesPtrOutput
	ToAccountResourcePropertiesPtrOutputWithContext(context.Context) AccountResourcePropertiesPtrOutput
}

type accountResourcePropertiesPtrType AccountResourcePropertiesArgs

func AccountResourcePropertiesPtr(v *AccountResourcePropertiesArgs) AccountResourcePropertiesPtrInput {
	return (*accountResourcePropertiesPtrType)(v)
}

func (*accountResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountResourceProperties)(nil)).Elem()
}

func (i *accountResourcePropertiesPtrType) ToAccountResourcePropertiesPtrOutput() AccountResourcePropertiesPtrOutput {
	return i.ToAccountResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *accountResourcePropertiesPtrType) ToAccountResourcePropertiesPtrOutputWithContext(ctx context.Context) AccountResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountResourcePropertiesPtrOutput)
}

type AccountResourcePropertiesOutput struct{ *pulumi.OutputState }

func (AccountResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountResourceProperties)(nil)).Elem()
}

func (o AccountResourcePropertiesOutput) ToAccountResourcePropertiesOutput() AccountResourcePropertiesOutput {
	return o
}

func (o AccountResourcePropertiesOutput) ToAccountResourcePropertiesOutputWithContext(ctx context.Context) AccountResourcePropertiesOutput {
	return o
}

func (o AccountResourcePropertiesOutput) ToAccountResourcePropertiesPtrOutput() AccountResourcePropertiesPtrOutput {
	return o.ToAccountResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o AccountResourcePropertiesOutput) ToAccountResourcePropertiesPtrOutputWithContext(ctx context.Context) AccountResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountResourceProperties) *AccountResourceProperties {
		return &v
	}).(AccountResourcePropertiesPtrOutput)
}

func (o AccountResourcePropertiesOutput) Configuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountResourceProperties) *string { return v.Configuration }).(pulumi.StringPtrOutput)
}

func (o AccountResourcePropertiesOutput) Cors() CorsRuleArrayOutput {
	return o.ApplyT(func(v AccountResourceProperties) []CorsRule { return v.Cors }).(CorsRuleArrayOutput)
}

func (o AccountResourcePropertiesOutput) EndpointAuthentications() EndpointAuthenticationArrayOutput {
	return o.ApplyT(func(v AccountResourceProperties) []EndpointAuthentication { return v.EndpointAuthentications }).(EndpointAuthenticationArrayOutput)
}

func (o AccountResourcePropertiesOutput) ReportsConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountResourceProperties) *string { return v.ReportsConnectionString }).(pulumi.StringPtrOutput)
}

type AccountResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AccountResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountResourceProperties)(nil)).Elem()
}

func (o AccountResourcePropertiesPtrOutput) ToAccountResourcePropertiesPtrOutput() AccountResourcePropertiesPtrOutput {
	return o
}

func (o AccountResourcePropertiesPtrOutput) ToAccountResourcePropertiesPtrOutputWithContext(ctx context.Context) AccountResourcePropertiesPtrOutput {
	return o
}

func (o AccountResourcePropertiesPtrOutput) Elem() AccountResourcePropertiesOutput {
	return o.ApplyT(func(v *AccountResourceProperties) AccountResourceProperties {
		if v != nil {
			return *v
		}
		var ret AccountResourceProperties
		return ret
	}).(AccountResourcePropertiesOutput)
}

func (o AccountResourcePropertiesPtrOutput) Configuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(pulumi.StringPtrOutput)
}

func (o AccountResourcePropertiesPtrOutput) Cors() CorsRuleArrayOutput {
	return o.ApplyT(func(v *AccountResourceProperties) []CorsRule {
		if v == nil {
			return nil
		}
		return v.Cors
	}).(CorsRuleArrayOutput)
}

func (o AccountResourcePropertiesPtrOutput) EndpointAuthentications() EndpointAuthenticationArrayOutput {
	return o.ApplyT(func(v *AccountResourceProperties) []EndpointAuthentication {
		if v == nil {
			return nil
		}
		return v.EndpointAuthentications
	}).(EndpointAuthenticationArrayOutput)
}

func (o AccountResourcePropertiesPtrOutput) ReportsConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ReportsConnectionString
	}).(pulumi.StringPtrOutput)
}

type AccountResourceResponseProperties struct {
	Configuration           *string                          `pulumi:"configuration"`
	Cors                    []CorsRuleResponse               `pulumi:"cors"`
	EndpointAuthentications []EndpointAuthenticationResponse `pulumi:"endpointAuthentications"`
	ProvisioningState       string                           `pulumi:"provisioningState"`
	ReportsConnectionString *string                          `pulumi:"reportsConnectionString"`
}

type AccountResourceResponsePropertiesOutput struct{ *pulumi.OutputState }

func (AccountResourceResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountResourceResponseProperties)(nil)).Elem()
}

func (o AccountResourceResponsePropertiesOutput) ToAccountResourceResponsePropertiesOutput() AccountResourceResponsePropertiesOutput {
	return o
}

func (o AccountResourceResponsePropertiesOutput) ToAccountResourceResponsePropertiesOutputWithContext(ctx context.Context) AccountResourceResponsePropertiesOutput {
	return o
}

func (o AccountResourceResponsePropertiesOutput) Configuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountResourceResponseProperties) *string { return v.Configuration }).(pulumi.StringPtrOutput)
}

func (o AccountResourceResponsePropertiesOutput) Cors() CorsRuleResponseArrayOutput {
	return o.ApplyT(func(v AccountResourceResponseProperties) []CorsRuleResponse { return v.Cors }).(CorsRuleResponseArrayOutput)
}

func (o AccountResourceResponsePropertiesOutput) EndpointAuthentications() EndpointAuthenticationResponseArrayOutput {
	return o.ApplyT(func(v AccountResourceResponseProperties) []EndpointAuthenticationResponse {
		return v.EndpointAuthentications
	}).(EndpointAuthenticationResponseArrayOutput)
}

func (o AccountResourceResponsePropertiesOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AccountResourceResponseProperties) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AccountResourceResponsePropertiesOutput) ReportsConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountResourceResponseProperties) *string { return v.ReportsConnectionString }).(pulumi.StringPtrOutput)
}

type CorsRule struct {
	AllowedHeaders  []string `pulumi:"allowedHeaders"`
	AllowedMethods  []string `pulumi:"allowedMethods"`
	AllowedOrigins  []string `pulumi:"allowedOrigins"`
	ExposedHeaders  []string `pulumi:"exposedHeaders"`
	MaxAgeInSeconds *int     `pulumi:"maxAgeInSeconds"`
}





type CorsRuleInput interface {
	pulumi.Input

	ToCorsRuleOutput() CorsRuleOutput
	ToCorsRuleOutputWithContext(context.Context) CorsRuleOutput
}

type CorsRuleArgs struct {
	AllowedHeaders  pulumi.StringArrayInput `pulumi:"allowedHeaders"`
	AllowedMethods  pulumi.StringArrayInput `pulumi:"allowedMethods"`
	AllowedOrigins  pulumi.StringArrayInput `pulumi:"allowedOrigins"`
	ExposedHeaders  pulumi.StringArrayInput `pulumi:"exposedHeaders"`
	MaxAgeInSeconds pulumi.IntPtrInput      `pulumi:"maxAgeInSeconds"`
}

func (CorsRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRule)(nil)).Elem()
}

func (i CorsRuleArgs) ToCorsRuleOutput() CorsRuleOutput {
	return i.ToCorsRuleOutputWithContext(context.Background())
}

func (i CorsRuleArgs) ToCorsRuleOutputWithContext(ctx context.Context) CorsRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRuleOutput)
}





type CorsRuleArrayInput interface {
	pulumi.Input

	ToCorsRuleArrayOutput() CorsRuleArrayOutput
	ToCorsRuleArrayOutputWithContext(context.Context) CorsRuleArrayOutput
}

type CorsRuleArray []CorsRuleInput

func (CorsRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsRule)(nil)).Elem()
}

func (i CorsRuleArray) ToCorsRuleArrayOutput() CorsRuleArrayOutput {
	return i.ToCorsRuleArrayOutputWithContext(context.Background())
}

func (i CorsRuleArray) ToCorsRuleArrayOutputWithContext(ctx context.Context) CorsRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRuleArrayOutput)
}

type CorsRuleOutput struct{ *pulumi.OutputState }

func (CorsRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRule)(nil)).Elem()
}

func (o CorsRuleOutput) ToCorsRuleOutput() CorsRuleOutput {
	return o
}

func (o CorsRuleOutput) ToCorsRuleOutputWithContext(ctx context.Context) CorsRuleOutput {
	return o
}

func (o CorsRuleOutput) AllowedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRule) []string { return v.AllowedHeaders }).(pulumi.StringArrayOutput)
}

func (o CorsRuleOutput) AllowedMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRule) []string { return v.AllowedMethods }).(pulumi.StringArrayOutput)
}

func (o CorsRuleOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRule) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

func (o CorsRuleOutput) ExposedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRule) []string { return v.ExposedHeaders }).(pulumi.StringArrayOutput)
}

func (o CorsRuleOutput) MaxAgeInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CorsRule) *int { return v.MaxAgeInSeconds }).(pulumi.IntPtrOutput)
}

type CorsRuleArrayOutput struct{ *pulumi.OutputState }

func (CorsRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsRule)(nil)).Elem()
}

func (o CorsRuleArrayOutput) ToCorsRuleArrayOutput() CorsRuleArrayOutput {
	return o
}

func (o CorsRuleArrayOutput) ToCorsRuleArrayOutputWithContext(ctx context.Context) CorsRuleArrayOutput {
	return o
}

func (o CorsRuleArrayOutput) Index(i pulumi.IntInput) CorsRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CorsRule {
		return vs[0].([]CorsRule)[vs[1].(int)]
	}).(CorsRuleOutput)
}

type CorsRuleResponse struct {
	AllowedHeaders  []string `pulumi:"allowedHeaders"`
	AllowedMethods  []string `pulumi:"allowedMethods"`
	AllowedOrigins  []string `pulumi:"allowedOrigins"`
	ExposedHeaders  []string `pulumi:"exposedHeaders"`
	MaxAgeInSeconds *int     `pulumi:"maxAgeInSeconds"`
}

type CorsRuleResponseOutput struct{ *pulumi.OutputState }

func (CorsRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRuleResponse)(nil)).Elem()
}

func (o CorsRuleResponseOutput) ToCorsRuleResponseOutput() CorsRuleResponseOutput {
	return o
}

func (o CorsRuleResponseOutput) ToCorsRuleResponseOutputWithContext(ctx context.Context) CorsRuleResponseOutput {
	return o
}

func (o CorsRuleResponseOutput) AllowedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRuleResponse) []string { return v.AllowedHeaders }).(pulumi.StringArrayOutput)
}

func (o CorsRuleResponseOutput) AllowedMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRuleResponse) []string { return v.AllowedMethods }).(pulumi.StringArrayOutput)
}

func (o CorsRuleResponseOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRuleResponse) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

func (o CorsRuleResponseOutput) ExposedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRuleResponse) []string { return v.ExposedHeaders }).(pulumi.StringArrayOutput)
}

func (o CorsRuleResponseOutput) MaxAgeInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CorsRuleResponse) *int { return v.MaxAgeInSeconds }).(pulumi.IntPtrOutput)
}

type CorsRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (CorsRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsRuleResponse)(nil)).Elem()
}

func (o CorsRuleResponseArrayOutput) ToCorsRuleResponseArrayOutput() CorsRuleResponseArrayOutput {
	return o
}

func (o CorsRuleResponseArrayOutput) ToCorsRuleResponseArrayOutputWithContext(ctx context.Context) CorsRuleResponseArrayOutput {
	return o
}

func (o CorsRuleResponseArrayOutput) Index(i pulumi.IntInput) CorsRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CorsRuleResponse {
		return vs[0].([]CorsRuleResponse)[vs[1].(int)]
	}).(CorsRuleResponseOutput)
}

type EndpointAuthentication struct {
	AadTenantID   *string `pulumi:"aadTenantID"`
	PrincipalID   *string `pulumi:"principalID"`
	PrincipalType *string `pulumi:"principalType"`
}





type EndpointAuthenticationInput interface {
	pulumi.Input

	ToEndpointAuthenticationOutput() EndpointAuthenticationOutput
	ToEndpointAuthenticationOutputWithContext(context.Context) EndpointAuthenticationOutput
}

type EndpointAuthenticationArgs struct {
	AadTenantID   pulumi.StringPtrInput `pulumi:"aadTenantID"`
	PrincipalID   pulumi.StringPtrInput `pulumi:"principalID"`
	PrincipalType pulumi.StringPtrInput `pulumi:"principalType"`
}

func (EndpointAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointAuthentication)(nil)).Elem()
}

func (i EndpointAuthenticationArgs) ToEndpointAuthenticationOutput() EndpointAuthenticationOutput {
	return i.ToEndpointAuthenticationOutputWithContext(context.Background())
}

func (i EndpointAuthenticationArgs) ToEndpointAuthenticationOutputWithContext(ctx context.Context) EndpointAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAuthenticationOutput)
}





type EndpointAuthenticationArrayInput interface {
	pulumi.Input

	ToEndpointAuthenticationArrayOutput() EndpointAuthenticationArrayOutput
	ToEndpointAuthenticationArrayOutputWithContext(context.Context) EndpointAuthenticationArrayOutput
}

type EndpointAuthenticationArray []EndpointAuthenticationInput

func (EndpointAuthenticationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointAuthentication)(nil)).Elem()
}

func (i EndpointAuthenticationArray) ToEndpointAuthenticationArrayOutput() EndpointAuthenticationArrayOutput {
	return i.ToEndpointAuthenticationArrayOutputWithContext(context.Background())
}

func (i EndpointAuthenticationArray) ToEndpointAuthenticationArrayOutputWithContext(ctx context.Context) EndpointAuthenticationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAuthenticationArrayOutput)
}

type EndpointAuthenticationOutput struct{ *pulumi.OutputState }

func (EndpointAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointAuthentication)(nil)).Elem()
}

func (o EndpointAuthenticationOutput) ToEndpointAuthenticationOutput() EndpointAuthenticationOutput {
	return o
}

func (o EndpointAuthenticationOutput) ToEndpointAuthenticationOutputWithContext(ctx context.Context) EndpointAuthenticationOutput {
	return o
}

func (o EndpointAuthenticationOutput) AadTenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointAuthentication) *string { return v.AadTenantID }).(pulumi.StringPtrOutput)
}

func (o EndpointAuthenticationOutput) PrincipalID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointAuthentication) *string { return v.PrincipalID }).(pulumi.StringPtrOutput)
}

func (o EndpointAuthenticationOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointAuthentication) *string { return v.PrincipalType }).(pulumi.StringPtrOutput)
}

type EndpointAuthenticationArrayOutput struct{ *pulumi.OutputState }

func (EndpointAuthenticationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointAuthentication)(nil)).Elem()
}

func (o EndpointAuthenticationArrayOutput) ToEndpointAuthenticationArrayOutput() EndpointAuthenticationArrayOutput {
	return o
}

func (o EndpointAuthenticationArrayOutput) ToEndpointAuthenticationArrayOutputWithContext(ctx context.Context) EndpointAuthenticationArrayOutput {
	return o
}

func (o EndpointAuthenticationArrayOutput) Index(i pulumi.IntInput) EndpointAuthenticationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointAuthentication {
		return vs[0].([]EndpointAuthentication)[vs[1].(int)]
	}).(EndpointAuthenticationOutput)
}

type EndpointAuthenticationResponse struct {
	AadTenantID   *string `pulumi:"aadTenantID"`
	PrincipalID   *string `pulumi:"principalID"`
	PrincipalType *string `pulumi:"principalType"`
}

type EndpointAuthenticationResponseOutput struct{ *pulumi.OutputState }

func (EndpointAuthenticationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointAuthenticationResponse)(nil)).Elem()
}

func (o EndpointAuthenticationResponseOutput) ToEndpointAuthenticationResponseOutput() EndpointAuthenticationResponseOutput {
	return o
}

func (o EndpointAuthenticationResponseOutput) ToEndpointAuthenticationResponseOutputWithContext(ctx context.Context) EndpointAuthenticationResponseOutput {
	return o
}

func (o EndpointAuthenticationResponseOutput) AadTenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointAuthenticationResponse) *string { return v.AadTenantID }).(pulumi.StringPtrOutput)
}

func (o EndpointAuthenticationResponseOutput) PrincipalID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointAuthenticationResponse) *string { return v.PrincipalID }).(pulumi.StringPtrOutput)
}

func (o EndpointAuthenticationResponseOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointAuthenticationResponse) *string { return v.PrincipalType }).(pulumi.StringPtrOutput)
}

type EndpointAuthenticationResponseArrayOutput struct{ *pulumi.OutputState }

func (EndpointAuthenticationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointAuthenticationResponse)(nil)).Elem()
}

func (o EndpointAuthenticationResponseArrayOutput) ToEndpointAuthenticationResponseArrayOutput() EndpointAuthenticationResponseArrayOutput {
	return o
}

func (o EndpointAuthenticationResponseArrayOutput) ToEndpointAuthenticationResponseArrayOutputWithContext(ctx context.Context) EndpointAuthenticationResponseArrayOutput {
	return o
}

func (o EndpointAuthenticationResponseArrayOutput) Index(i pulumi.IntInput) EndpointAuthenticationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointAuthenticationResponse {
		return vs[0].([]EndpointAuthenticationResponse)[vs[1].(int)]
	}).(EndpointAuthenticationResponseOutput)
}

type ManagedServiceIdentity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return i.ToManagedServiceIdentityOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput)
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput).ToManagedServiceIdentityPtrOutputWithContext(ctx)
}









type ManagedServiceIdentityPtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput
	ToManagedServiceIdentityPtrOutputWithContext(context.Context) ManagedServiceIdentityPtrOutput
}

type managedServiceIdentityPtrType ManagedServiceIdentityArgs

func ManagedServiceIdentityPtr(v *ManagedServiceIdentityArgs) ManagedServiceIdentityPtrInput {
	return (*managedServiceIdentityPtrType)(v)
}

func (*managedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityPtrOutput)
}

type ManagedServiceIdentityOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentity) *ManagedServiceIdentity {
		return &v
	}).(ManagedServiceIdentityPtrOutput)
}

func (o ManagedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) Elem() ManagedServiceIdentityOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) ManagedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentity
		return ret
	}).(ManagedServiceIdentityOutput)
}

func (o ManagedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedServiceIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   string                                  `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ManagedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ManagedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) Elem() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) ManagedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityResponse
		return ret
	}).(ManagedServiceIdentityResponseOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ModelingInputData struct {
	ConnectionString *string `pulumi:"connectionString"`
}





type ModelingInputDataInput interface {
	pulumi.Input

	ToModelingInputDataOutput() ModelingInputDataOutput
	ToModelingInputDataOutputWithContext(context.Context) ModelingInputDataOutput
}

type ModelingInputDataArgs struct {
	ConnectionString pulumi.StringPtrInput `pulumi:"connectionString"`
}

func (ModelingInputDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelingInputData)(nil)).Elem()
}

func (i ModelingInputDataArgs) ToModelingInputDataOutput() ModelingInputDataOutput {
	return i.ToModelingInputDataOutputWithContext(context.Background())
}

func (i ModelingInputDataArgs) ToModelingInputDataOutputWithContext(ctx context.Context) ModelingInputDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelingInputDataOutput)
}

func (i ModelingInputDataArgs) ToModelingInputDataPtrOutput() ModelingInputDataPtrOutput {
	return i.ToModelingInputDataPtrOutputWithContext(context.Background())
}

func (i ModelingInputDataArgs) ToModelingInputDataPtrOutputWithContext(ctx context.Context) ModelingInputDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelingInputDataOutput).ToModelingInputDataPtrOutputWithContext(ctx)
}









type ModelingInputDataPtrInput interface {
	pulumi.Input

	ToModelingInputDataPtrOutput() ModelingInputDataPtrOutput
	ToModelingInputDataPtrOutputWithContext(context.Context) ModelingInputDataPtrOutput
}

type modelingInputDataPtrType ModelingInputDataArgs

func ModelingInputDataPtr(v *ModelingInputDataArgs) ModelingInputDataPtrInput {
	return (*modelingInputDataPtrType)(v)
}

func (*modelingInputDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelingInputData)(nil)).Elem()
}

func (i *modelingInputDataPtrType) ToModelingInputDataPtrOutput() ModelingInputDataPtrOutput {
	return i.ToModelingInputDataPtrOutputWithContext(context.Background())
}

func (i *modelingInputDataPtrType) ToModelingInputDataPtrOutputWithContext(ctx context.Context) ModelingInputDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelingInputDataPtrOutput)
}

type ModelingInputDataOutput struct{ *pulumi.OutputState }

func (ModelingInputDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelingInputData)(nil)).Elem()
}

func (o ModelingInputDataOutput) ToModelingInputDataOutput() ModelingInputDataOutput {
	return o
}

func (o ModelingInputDataOutput) ToModelingInputDataOutputWithContext(ctx context.Context) ModelingInputDataOutput {
	return o
}

func (o ModelingInputDataOutput) ToModelingInputDataPtrOutput() ModelingInputDataPtrOutput {
	return o.ToModelingInputDataPtrOutputWithContext(context.Background())
}

func (o ModelingInputDataOutput) ToModelingInputDataPtrOutputWithContext(ctx context.Context) ModelingInputDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ModelingInputData) *ModelingInputData {
		return &v
	}).(ModelingInputDataPtrOutput)
}

func (o ModelingInputDataOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelingInputData) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

type ModelingInputDataPtrOutput struct{ *pulumi.OutputState }

func (ModelingInputDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelingInputData)(nil)).Elem()
}

func (o ModelingInputDataPtrOutput) ToModelingInputDataPtrOutput() ModelingInputDataPtrOutput {
	return o
}

func (o ModelingInputDataPtrOutput) ToModelingInputDataPtrOutputWithContext(ctx context.Context) ModelingInputDataPtrOutput {
	return o
}

func (o ModelingInputDataPtrOutput) Elem() ModelingInputDataOutput {
	return o.ApplyT(func(v *ModelingInputData) ModelingInputData {
		if v != nil {
			return *v
		}
		var ret ModelingInputData
		return ret
	}).(ModelingInputDataOutput)
}

func (o ModelingInputDataPtrOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelingInputData) *string {
		if v == nil {
			return nil
		}
		return v.ConnectionString
	}).(pulumi.StringPtrOutput)
}

type ModelingInputDataResponse struct {
	ConnectionString *string `pulumi:"connectionString"`
}

type ModelingInputDataResponseOutput struct{ *pulumi.OutputState }

func (ModelingInputDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelingInputDataResponse)(nil)).Elem()
}

func (o ModelingInputDataResponseOutput) ToModelingInputDataResponseOutput() ModelingInputDataResponseOutput {
	return o
}

func (o ModelingInputDataResponseOutput) ToModelingInputDataResponseOutputWithContext(ctx context.Context) ModelingInputDataResponseOutput {
	return o
}

func (o ModelingInputDataResponseOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelingInputDataResponse) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

type ModelingInputDataResponsePtrOutput struct{ *pulumi.OutputState }

func (ModelingInputDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelingInputDataResponse)(nil)).Elem()
}

func (o ModelingInputDataResponsePtrOutput) ToModelingInputDataResponsePtrOutput() ModelingInputDataResponsePtrOutput {
	return o
}

func (o ModelingInputDataResponsePtrOutput) ToModelingInputDataResponsePtrOutputWithContext(ctx context.Context) ModelingInputDataResponsePtrOutput {
	return o
}

func (o ModelingInputDataResponsePtrOutput) Elem() ModelingInputDataResponseOutput {
	return o.ApplyT(func(v *ModelingInputDataResponse) ModelingInputDataResponse {
		if v != nil {
			return *v
		}
		var ret ModelingInputDataResponse
		return ret
	}).(ModelingInputDataResponseOutput)
}

func (o ModelingInputDataResponsePtrOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelingInputDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConnectionString
	}).(pulumi.StringPtrOutput)
}

type ModelingResourceProperties struct {
	Features  *string            `pulumi:"features"`
	Frequency *string            `pulumi:"frequency"`
	InputData *ModelingInputData `pulumi:"inputData"`
	Size      *string            `pulumi:"size"`
}





type ModelingResourcePropertiesInput interface {
	pulumi.Input

	ToModelingResourcePropertiesOutput() ModelingResourcePropertiesOutput
	ToModelingResourcePropertiesOutputWithContext(context.Context) ModelingResourcePropertiesOutput
}

type ModelingResourcePropertiesArgs struct {
	Features  pulumi.StringPtrInput     `pulumi:"features"`
	Frequency pulumi.StringPtrInput     `pulumi:"frequency"`
	InputData ModelingInputDataPtrInput `pulumi:"inputData"`
	Size      pulumi.StringPtrInput     `pulumi:"size"`
}

func (ModelingResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelingResourceProperties)(nil)).Elem()
}

func (i ModelingResourcePropertiesArgs) ToModelingResourcePropertiesOutput() ModelingResourcePropertiesOutput {
	return i.ToModelingResourcePropertiesOutputWithContext(context.Background())
}

func (i ModelingResourcePropertiesArgs) ToModelingResourcePropertiesOutputWithContext(ctx context.Context) ModelingResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelingResourcePropertiesOutput)
}

func (i ModelingResourcePropertiesArgs) ToModelingResourcePropertiesPtrOutput() ModelingResourcePropertiesPtrOutput {
	return i.ToModelingResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i ModelingResourcePropertiesArgs) ToModelingResourcePropertiesPtrOutputWithContext(ctx context.Context) ModelingResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelingResourcePropertiesOutput).ToModelingResourcePropertiesPtrOutputWithContext(ctx)
}









type ModelingResourcePropertiesPtrInput interface {
	pulumi.Input

	ToModelingResourcePropertiesPtrOutput() ModelingResourcePropertiesPtrOutput
	ToModelingResourcePropertiesPtrOutputWithContext(context.Context) ModelingResourcePropertiesPtrOutput
}

type modelingResourcePropertiesPtrType ModelingResourcePropertiesArgs

func ModelingResourcePropertiesPtr(v *ModelingResourcePropertiesArgs) ModelingResourcePropertiesPtrInput {
	return (*modelingResourcePropertiesPtrType)(v)
}

func (*modelingResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelingResourceProperties)(nil)).Elem()
}

func (i *modelingResourcePropertiesPtrType) ToModelingResourcePropertiesPtrOutput() ModelingResourcePropertiesPtrOutput {
	return i.ToModelingResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *modelingResourcePropertiesPtrType) ToModelingResourcePropertiesPtrOutputWithContext(ctx context.Context) ModelingResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelingResourcePropertiesPtrOutput)
}

type ModelingResourcePropertiesOutput struct{ *pulumi.OutputState }

func (ModelingResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelingResourceProperties)(nil)).Elem()
}

func (o ModelingResourcePropertiesOutput) ToModelingResourcePropertiesOutput() ModelingResourcePropertiesOutput {
	return o
}

func (o ModelingResourcePropertiesOutput) ToModelingResourcePropertiesOutputWithContext(ctx context.Context) ModelingResourcePropertiesOutput {
	return o
}

func (o ModelingResourcePropertiesOutput) ToModelingResourcePropertiesPtrOutput() ModelingResourcePropertiesPtrOutput {
	return o.ToModelingResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o ModelingResourcePropertiesOutput) ToModelingResourcePropertiesPtrOutputWithContext(ctx context.Context) ModelingResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ModelingResourceProperties) *ModelingResourceProperties {
		return &v
	}).(ModelingResourcePropertiesPtrOutput)
}

func (o ModelingResourcePropertiesOutput) Features() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelingResourceProperties) *string { return v.Features }).(pulumi.StringPtrOutput)
}

func (o ModelingResourcePropertiesOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelingResourceProperties) *string { return v.Frequency }).(pulumi.StringPtrOutput)
}

func (o ModelingResourcePropertiesOutput) InputData() ModelingInputDataPtrOutput {
	return o.ApplyT(func(v ModelingResourceProperties) *ModelingInputData { return v.InputData }).(ModelingInputDataPtrOutput)
}

func (o ModelingResourcePropertiesOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelingResourceProperties) *string { return v.Size }).(pulumi.StringPtrOutput)
}

type ModelingResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ModelingResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelingResourceProperties)(nil)).Elem()
}

func (o ModelingResourcePropertiesPtrOutput) ToModelingResourcePropertiesPtrOutput() ModelingResourcePropertiesPtrOutput {
	return o
}

func (o ModelingResourcePropertiesPtrOutput) ToModelingResourcePropertiesPtrOutputWithContext(ctx context.Context) ModelingResourcePropertiesPtrOutput {
	return o
}

func (o ModelingResourcePropertiesPtrOutput) Elem() ModelingResourcePropertiesOutput {
	return o.ApplyT(func(v *ModelingResourceProperties) ModelingResourceProperties {
		if v != nil {
			return *v
		}
		var ret ModelingResourceProperties
		return ret
	}).(ModelingResourcePropertiesOutput)
}

func (o ModelingResourcePropertiesPtrOutput) Features() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelingResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Features
	}).(pulumi.StringPtrOutput)
}

func (o ModelingResourcePropertiesPtrOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelingResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Frequency
	}).(pulumi.StringPtrOutput)
}

func (o ModelingResourcePropertiesPtrOutput) InputData() ModelingInputDataPtrOutput {
	return o.ApplyT(func(v *ModelingResourceProperties) *ModelingInputData {
		if v == nil {
			return nil
		}
		return v.InputData
	}).(ModelingInputDataPtrOutput)
}

func (o ModelingResourcePropertiesPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelingResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

type ModelingResourceResponseProperties struct {
	Features          *string                    `pulumi:"features"`
	Frequency         *string                    `pulumi:"frequency"`
	InputData         *ModelingInputDataResponse `pulumi:"inputData"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	Size              *string                    `pulumi:"size"`
}

type ModelingResourceResponsePropertiesOutput struct{ *pulumi.OutputState }

func (ModelingResourceResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelingResourceResponseProperties)(nil)).Elem()
}

func (o ModelingResourceResponsePropertiesOutput) ToModelingResourceResponsePropertiesOutput() ModelingResourceResponsePropertiesOutput {
	return o
}

func (o ModelingResourceResponsePropertiesOutput) ToModelingResourceResponsePropertiesOutputWithContext(ctx context.Context) ModelingResourceResponsePropertiesOutput {
	return o
}

func (o ModelingResourceResponsePropertiesOutput) Features() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelingResourceResponseProperties) *string { return v.Features }).(pulumi.StringPtrOutput)
}

func (o ModelingResourceResponsePropertiesOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelingResourceResponseProperties) *string { return v.Frequency }).(pulumi.StringPtrOutput)
}

func (o ModelingResourceResponsePropertiesOutput) InputData() ModelingInputDataResponsePtrOutput {
	return o.ApplyT(func(v ModelingResourceResponseProperties) *ModelingInputDataResponse { return v.InputData }).(ModelingInputDataResponsePtrOutput)
}

func (o ModelingResourceResponsePropertiesOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ModelingResourceResponseProperties) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ModelingResourceResponsePropertiesOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelingResourceResponseProperties) *string { return v.Size }).(pulumi.StringPtrOutput)
}

type ServiceEndpointResourceProperties struct {
	PreAllocatedCapacity *int `pulumi:"preAllocatedCapacity"`
}





type ServiceEndpointResourcePropertiesInput interface {
	pulumi.Input

	ToServiceEndpointResourcePropertiesOutput() ServiceEndpointResourcePropertiesOutput
	ToServiceEndpointResourcePropertiesOutputWithContext(context.Context) ServiceEndpointResourcePropertiesOutput
}

type ServiceEndpointResourcePropertiesArgs struct {
	PreAllocatedCapacity pulumi.IntPtrInput `pulumi:"preAllocatedCapacity"`
}

func (ServiceEndpointResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointResourceProperties)(nil)).Elem()
}

func (i ServiceEndpointResourcePropertiesArgs) ToServiceEndpointResourcePropertiesOutput() ServiceEndpointResourcePropertiesOutput {
	return i.ToServiceEndpointResourcePropertiesOutputWithContext(context.Background())
}

func (i ServiceEndpointResourcePropertiesArgs) ToServiceEndpointResourcePropertiesOutputWithContext(ctx context.Context) ServiceEndpointResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointResourcePropertiesOutput)
}

func (i ServiceEndpointResourcePropertiesArgs) ToServiceEndpointResourcePropertiesPtrOutput() ServiceEndpointResourcePropertiesPtrOutput {
	return i.ToServiceEndpointResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i ServiceEndpointResourcePropertiesArgs) ToServiceEndpointResourcePropertiesPtrOutputWithContext(ctx context.Context) ServiceEndpointResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointResourcePropertiesOutput).ToServiceEndpointResourcePropertiesPtrOutputWithContext(ctx)
}









type ServiceEndpointResourcePropertiesPtrInput interface {
	pulumi.Input

	ToServiceEndpointResourcePropertiesPtrOutput() ServiceEndpointResourcePropertiesPtrOutput
	ToServiceEndpointResourcePropertiesPtrOutputWithContext(context.Context) ServiceEndpointResourcePropertiesPtrOutput
}

type serviceEndpointResourcePropertiesPtrType ServiceEndpointResourcePropertiesArgs

func ServiceEndpointResourcePropertiesPtr(v *ServiceEndpointResourcePropertiesArgs) ServiceEndpointResourcePropertiesPtrInput {
	return (*serviceEndpointResourcePropertiesPtrType)(v)
}

func (*serviceEndpointResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointResourceProperties)(nil)).Elem()
}

func (i *serviceEndpointResourcePropertiesPtrType) ToServiceEndpointResourcePropertiesPtrOutput() ServiceEndpointResourcePropertiesPtrOutput {
	return i.ToServiceEndpointResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *serviceEndpointResourcePropertiesPtrType) ToServiceEndpointResourcePropertiesPtrOutputWithContext(ctx context.Context) ServiceEndpointResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointResourcePropertiesPtrOutput)
}

type ServiceEndpointResourcePropertiesOutput struct{ *pulumi.OutputState }

func (ServiceEndpointResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointResourceProperties)(nil)).Elem()
}

func (o ServiceEndpointResourcePropertiesOutput) ToServiceEndpointResourcePropertiesOutput() ServiceEndpointResourcePropertiesOutput {
	return o
}

func (o ServiceEndpointResourcePropertiesOutput) ToServiceEndpointResourcePropertiesOutputWithContext(ctx context.Context) ServiceEndpointResourcePropertiesOutput {
	return o
}

func (o ServiceEndpointResourcePropertiesOutput) ToServiceEndpointResourcePropertiesPtrOutput() ServiceEndpointResourcePropertiesPtrOutput {
	return o.ToServiceEndpointResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o ServiceEndpointResourcePropertiesOutput) ToServiceEndpointResourcePropertiesPtrOutputWithContext(ctx context.Context) ServiceEndpointResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceEndpointResourceProperties) *ServiceEndpointResourceProperties {
		return &v
	}).(ServiceEndpointResourcePropertiesPtrOutput)
}

func (o ServiceEndpointResourcePropertiesOutput) PreAllocatedCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceEndpointResourceProperties) *int { return v.PreAllocatedCapacity }).(pulumi.IntPtrOutput)
}

type ServiceEndpointResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ServiceEndpointResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointResourceProperties)(nil)).Elem()
}

func (o ServiceEndpointResourcePropertiesPtrOutput) ToServiceEndpointResourcePropertiesPtrOutput() ServiceEndpointResourcePropertiesPtrOutput {
	return o
}

func (o ServiceEndpointResourcePropertiesPtrOutput) ToServiceEndpointResourcePropertiesPtrOutputWithContext(ctx context.Context) ServiceEndpointResourcePropertiesPtrOutput {
	return o
}

func (o ServiceEndpointResourcePropertiesPtrOutput) Elem() ServiceEndpointResourcePropertiesOutput {
	return o.ApplyT(func(v *ServiceEndpointResourceProperties) ServiceEndpointResourceProperties {
		if v != nil {
			return *v
		}
		var ret ServiceEndpointResourceProperties
		return ret
	}).(ServiceEndpointResourcePropertiesOutput)
}

func (o ServiceEndpointResourcePropertiesPtrOutput) PreAllocatedCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointResourceProperties) *int {
		if v == nil {
			return nil
		}
		return v.PreAllocatedCapacity
	}).(pulumi.IntPtrOutput)
}

type ServiceEndpointResourceResponseProperties struct {
	PairedLocation       string `pulumi:"pairedLocation"`
	PreAllocatedCapacity *int   `pulumi:"preAllocatedCapacity"`
	ProvisioningState    string `pulumi:"provisioningState"`
	Url                  string `pulumi:"url"`
}

type ServiceEndpointResourceResponsePropertiesOutput struct{ *pulumi.OutputState }

func (ServiceEndpointResourceResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointResourceResponseProperties)(nil)).Elem()
}

func (o ServiceEndpointResourceResponsePropertiesOutput) ToServiceEndpointResourceResponsePropertiesOutput() ServiceEndpointResourceResponsePropertiesOutput {
	return o
}

func (o ServiceEndpointResourceResponsePropertiesOutput) ToServiceEndpointResourceResponsePropertiesOutputWithContext(ctx context.Context) ServiceEndpointResourceResponsePropertiesOutput {
	return o
}

func (o ServiceEndpointResourceResponsePropertiesOutput) PairedLocation() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceEndpointResourceResponseProperties) string { return v.PairedLocation }).(pulumi.StringOutput)
}

func (o ServiceEndpointResourceResponsePropertiesOutput) PreAllocatedCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceEndpointResourceResponseProperties) *int { return v.PreAllocatedCapacity }).(pulumi.IntPtrOutput)
}

func (o ServiceEndpointResourceResponsePropertiesOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceEndpointResourceResponseProperties) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServiceEndpointResourceResponsePropertiesOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceEndpointResourceResponseProperties) string { return v.Url }).(pulumi.StringOutput)
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
	pulumi.RegisterOutputType(AccountResourcePropertiesOutput{})
	pulumi.RegisterOutputType(AccountResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AccountResourceResponsePropertiesOutput{})
	pulumi.RegisterOutputType(CorsRuleOutput{})
	pulumi.RegisterOutputType(CorsRuleArrayOutput{})
	pulumi.RegisterOutputType(CorsRuleResponseOutput{})
	pulumi.RegisterOutputType(CorsRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(EndpointAuthenticationOutput{})
	pulumi.RegisterOutputType(EndpointAuthenticationArrayOutput{})
	pulumi.RegisterOutputType(EndpointAuthenticationResponseOutput{})
	pulumi.RegisterOutputType(EndpointAuthenticationResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ModelingInputDataOutput{})
	pulumi.RegisterOutputType(ModelingInputDataPtrOutput{})
	pulumi.RegisterOutputType(ModelingInputDataResponseOutput{})
	pulumi.RegisterOutputType(ModelingInputDataResponsePtrOutput{})
	pulumi.RegisterOutputType(ModelingResourcePropertiesOutput{})
	pulumi.RegisterOutputType(ModelingResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ModelingResourceResponsePropertiesOutput{})
	pulumi.RegisterOutputType(ServiceEndpointResourcePropertiesOutput{})
	pulumi.RegisterOutputType(ServiceEndpointResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ServiceEndpointResourceResponsePropertiesOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}

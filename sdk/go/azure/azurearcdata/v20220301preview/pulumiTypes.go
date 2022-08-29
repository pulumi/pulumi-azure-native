


package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActiveDirectoryConnectorDNSDetails struct {
	DomainName                *string  `pulumi:"domainName"`
	NameserverIPAddresses     []string `pulumi:"nameserverIPAddresses"`
	PreferK8sDnsForPtrLookups *bool    `pulumi:"preferK8sDnsForPtrLookups"`
	Replicas                  *float64 `pulumi:"replicas"`
}


func (val *ActiveDirectoryConnectorDNSDetails) Defaults() *ActiveDirectoryConnectorDNSDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PreferK8sDnsForPtrLookups) {
		preferK8sDnsForPtrLookups_ := true
		tmp.PreferK8sDnsForPtrLookups = &preferK8sDnsForPtrLookups_
	}
	if isZero(tmp.Replicas) {
		replicas_ := 1.0
		tmp.Replicas = &replicas_
	}
	return &tmp
}





type ActiveDirectoryConnectorDNSDetailsInput interface {
	pulumi.Input

	ToActiveDirectoryConnectorDNSDetailsOutput() ActiveDirectoryConnectorDNSDetailsOutput
	ToActiveDirectoryConnectorDNSDetailsOutputWithContext(context.Context) ActiveDirectoryConnectorDNSDetailsOutput
}

type ActiveDirectoryConnectorDNSDetailsArgs struct {
	DomainName                pulumi.StringPtrInput   `pulumi:"domainName"`
	NameserverIPAddresses     pulumi.StringArrayInput `pulumi:"nameserverIPAddresses"`
	PreferK8sDnsForPtrLookups pulumi.BoolPtrInput     `pulumi:"preferK8sDnsForPtrLookups"`
	Replicas                  pulumi.Float64PtrInput  `pulumi:"replicas"`
}


func (val *ActiveDirectoryConnectorDNSDetailsArgs) Defaults() *ActiveDirectoryConnectorDNSDetailsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PreferK8sDnsForPtrLookups) {
		tmp.PreferK8sDnsForPtrLookups = pulumi.BoolPtr(true)
	}
	if isZero(tmp.Replicas) {
		tmp.Replicas = pulumi.Float64Ptr(1.0)
	}
	return &tmp
}
func (ActiveDirectoryConnectorDNSDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryConnectorDNSDetails)(nil)).Elem()
}

func (i ActiveDirectoryConnectorDNSDetailsArgs) ToActiveDirectoryConnectorDNSDetailsOutput() ActiveDirectoryConnectorDNSDetailsOutput {
	return i.ToActiveDirectoryConnectorDNSDetailsOutputWithContext(context.Background())
}

func (i ActiveDirectoryConnectorDNSDetailsArgs) ToActiveDirectoryConnectorDNSDetailsOutputWithContext(ctx context.Context) ActiveDirectoryConnectorDNSDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryConnectorDNSDetailsOutput)
}

type ActiveDirectoryConnectorDNSDetailsOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryConnectorDNSDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryConnectorDNSDetails)(nil)).Elem()
}

func (o ActiveDirectoryConnectorDNSDetailsOutput) ToActiveDirectoryConnectorDNSDetailsOutput() ActiveDirectoryConnectorDNSDetailsOutput {
	return o
}

func (o ActiveDirectoryConnectorDNSDetailsOutput) ToActiveDirectoryConnectorDNSDetailsOutputWithContext(ctx context.Context) ActiveDirectoryConnectorDNSDetailsOutput {
	return o
}

func (o ActiveDirectoryConnectorDNSDetailsOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDNSDetails) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryConnectorDNSDetailsOutput) NameserverIPAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDNSDetails) []string { return v.NameserverIPAddresses }).(pulumi.StringArrayOutput)
}

func (o ActiveDirectoryConnectorDNSDetailsOutput) PreferK8sDnsForPtrLookups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDNSDetails) *bool { return v.PreferK8sDnsForPtrLookups }).(pulumi.BoolPtrOutput)
}

func (o ActiveDirectoryConnectorDNSDetailsOutput) Replicas() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDNSDetails) *float64 { return v.Replicas }).(pulumi.Float64PtrOutput)
}

type ActiveDirectoryConnectorDNSDetailsResponse struct {
	DomainName                *string  `pulumi:"domainName"`
	NameserverIPAddresses     []string `pulumi:"nameserverIPAddresses"`
	PreferK8sDnsForPtrLookups *bool    `pulumi:"preferK8sDnsForPtrLookups"`
	Replicas                  *float64 `pulumi:"replicas"`
}


func (val *ActiveDirectoryConnectorDNSDetailsResponse) Defaults() *ActiveDirectoryConnectorDNSDetailsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PreferK8sDnsForPtrLookups) {
		preferK8sDnsForPtrLookups_ := true
		tmp.PreferK8sDnsForPtrLookups = &preferK8sDnsForPtrLookups_
	}
	if isZero(tmp.Replicas) {
		replicas_ := 1.0
		tmp.Replicas = &replicas_
	}
	return &tmp
}

type ActiveDirectoryConnectorDNSDetailsResponseOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryConnectorDNSDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryConnectorDNSDetailsResponse)(nil)).Elem()
}

func (o ActiveDirectoryConnectorDNSDetailsResponseOutput) ToActiveDirectoryConnectorDNSDetailsResponseOutput() ActiveDirectoryConnectorDNSDetailsResponseOutput {
	return o
}

func (o ActiveDirectoryConnectorDNSDetailsResponseOutput) ToActiveDirectoryConnectorDNSDetailsResponseOutputWithContext(ctx context.Context) ActiveDirectoryConnectorDNSDetailsResponseOutput {
	return o
}

func (o ActiveDirectoryConnectorDNSDetailsResponseOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDNSDetailsResponse) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryConnectorDNSDetailsResponseOutput) NameserverIPAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDNSDetailsResponse) []string { return v.NameserverIPAddresses }).(pulumi.StringArrayOutput)
}

func (o ActiveDirectoryConnectorDNSDetailsResponseOutput) PreferK8sDnsForPtrLookups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDNSDetailsResponse) *bool { return v.PreferK8sDnsForPtrLookups }).(pulumi.BoolPtrOutput)
}

func (o ActiveDirectoryConnectorDNSDetailsResponseOutput) Replicas() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDNSDetailsResponse) *float64 { return v.Replicas }).(pulumi.Float64PtrOutput)
}

type ActiveDirectoryConnectorDomainDetails struct {
	DomainControllers          ActiveDirectoryDomainControllers `pulumi:"domainControllers"`
	NetbiosDomainName          *string                          `pulumi:"netbiosDomainName"`
	OuDistinguishedName        *string                          `pulumi:"ouDistinguishedName"`
	Realm                      string                           `pulumi:"realm"`
	ServiceAccountProvisioning *string                          `pulumi:"serviceAccountProvisioning"`
}


func (val *ActiveDirectoryConnectorDomainDetails) Defaults() *ActiveDirectoryConnectorDomainDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceAccountProvisioning) {
		serviceAccountProvisioning_ := "manual"
		tmp.ServiceAccountProvisioning = &serviceAccountProvisioning_
	}
	return &tmp
}





type ActiveDirectoryConnectorDomainDetailsInput interface {
	pulumi.Input

	ToActiveDirectoryConnectorDomainDetailsOutput() ActiveDirectoryConnectorDomainDetailsOutput
	ToActiveDirectoryConnectorDomainDetailsOutputWithContext(context.Context) ActiveDirectoryConnectorDomainDetailsOutput
}

type ActiveDirectoryConnectorDomainDetailsArgs struct {
	DomainControllers          ActiveDirectoryDomainControllersInput `pulumi:"domainControllers"`
	NetbiosDomainName          pulumi.StringPtrInput                 `pulumi:"netbiosDomainName"`
	OuDistinguishedName        pulumi.StringPtrInput                 `pulumi:"ouDistinguishedName"`
	Realm                      pulumi.StringInput                    `pulumi:"realm"`
	ServiceAccountProvisioning pulumi.StringPtrInput                 `pulumi:"serviceAccountProvisioning"`
}


func (val *ActiveDirectoryConnectorDomainDetailsArgs) Defaults() *ActiveDirectoryConnectorDomainDetailsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceAccountProvisioning) {
		tmp.ServiceAccountProvisioning = pulumi.StringPtr("manual")
	}
	return &tmp
}
func (ActiveDirectoryConnectorDomainDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryConnectorDomainDetails)(nil)).Elem()
}

func (i ActiveDirectoryConnectorDomainDetailsArgs) ToActiveDirectoryConnectorDomainDetailsOutput() ActiveDirectoryConnectorDomainDetailsOutput {
	return i.ToActiveDirectoryConnectorDomainDetailsOutputWithContext(context.Background())
}

func (i ActiveDirectoryConnectorDomainDetailsArgs) ToActiveDirectoryConnectorDomainDetailsOutputWithContext(ctx context.Context) ActiveDirectoryConnectorDomainDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryConnectorDomainDetailsOutput)
}

type ActiveDirectoryConnectorDomainDetailsOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryConnectorDomainDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryConnectorDomainDetails)(nil)).Elem()
}

func (o ActiveDirectoryConnectorDomainDetailsOutput) ToActiveDirectoryConnectorDomainDetailsOutput() ActiveDirectoryConnectorDomainDetailsOutput {
	return o
}

func (o ActiveDirectoryConnectorDomainDetailsOutput) ToActiveDirectoryConnectorDomainDetailsOutputWithContext(ctx context.Context) ActiveDirectoryConnectorDomainDetailsOutput {
	return o
}

func (o ActiveDirectoryConnectorDomainDetailsOutput) DomainControllers() ActiveDirectoryDomainControllersOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDomainDetails) ActiveDirectoryDomainControllers {
		return v.DomainControllers
	}).(ActiveDirectoryDomainControllersOutput)
}

func (o ActiveDirectoryConnectorDomainDetailsOutput) NetbiosDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDomainDetails) *string { return v.NetbiosDomainName }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryConnectorDomainDetailsOutput) OuDistinguishedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDomainDetails) *string { return v.OuDistinguishedName }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryConnectorDomainDetailsOutput) Realm() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDomainDetails) string { return v.Realm }).(pulumi.StringOutput)
}

func (o ActiveDirectoryConnectorDomainDetailsOutput) ServiceAccountProvisioning() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDomainDetails) *string { return v.ServiceAccountProvisioning }).(pulumi.StringPtrOutput)
}

type ActiveDirectoryConnectorDomainDetailsResponse struct {
	DomainControllers          ActiveDirectoryDomainControllersResponse `pulumi:"domainControllers"`
	NetbiosDomainName          *string                                  `pulumi:"netbiosDomainName"`
	OuDistinguishedName        *string                                  `pulumi:"ouDistinguishedName"`
	Realm                      string                                   `pulumi:"realm"`
	ServiceAccountProvisioning *string                                  `pulumi:"serviceAccountProvisioning"`
}


func (val *ActiveDirectoryConnectorDomainDetailsResponse) Defaults() *ActiveDirectoryConnectorDomainDetailsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceAccountProvisioning) {
		serviceAccountProvisioning_ := "manual"
		tmp.ServiceAccountProvisioning = &serviceAccountProvisioning_
	}
	return &tmp
}

type ActiveDirectoryConnectorDomainDetailsResponseOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryConnectorDomainDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryConnectorDomainDetailsResponse)(nil)).Elem()
}

func (o ActiveDirectoryConnectorDomainDetailsResponseOutput) ToActiveDirectoryConnectorDomainDetailsResponseOutput() ActiveDirectoryConnectorDomainDetailsResponseOutput {
	return o
}

func (o ActiveDirectoryConnectorDomainDetailsResponseOutput) ToActiveDirectoryConnectorDomainDetailsResponseOutputWithContext(ctx context.Context) ActiveDirectoryConnectorDomainDetailsResponseOutput {
	return o
}

func (o ActiveDirectoryConnectorDomainDetailsResponseOutput) DomainControllers() ActiveDirectoryDomainControllersResponseOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDomainDetailsResponse) ActiveDirectoryDomainControllersResponse {
		return v.DomainControllers
	}).(ActiveDirectoryDomainControllersResponseOutput)
}

func (o ActiveDirectoryConnectorDomainDetailsResponseOutput) NetbiosDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDomainDetailsResponse) *string { return v.NetbiosDomainName }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryConnectorDomainDetailsResponseOutput) OuDistinguishedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDomainDetailsResponse) *string { return v.OuDistinguishedName }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryConnectorDomainDetailsResponseOutput) Realm() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDomainDetailsResponse) string { return v.Realm }).(pulumi.StringOutput)
}

func (o ActiveDirectoryConnectorDomainDetailsResponseOutput) ServiceAccountProvisioning() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorDomainDetailsResponse) *string { return v.ServiceAccountProvisioning }).(pulumi.StringPtrOutput)
}

type ActiveDirectoryConnectorProperties struct {
	DomainServiceAccountLoginInformation *BasicLoginInformation          `pulumi:"domainServiceAccountLoginInformation"`
	Spec                                 ActiveDirectoryConnectorSpec    `pulumi:"spec"`
	Status                               *ActiveDirectoryConnectorStatus `pulumi:"status"`
}


func (val *ActiveDirectoryConnectorProperties) Defaults() *ActiveDirectoryConnectorProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Spec = *tmp.Spec.Defaults()

	return &tmp
}





type ActiveDirectoryConnectorPropertiesInput interface {
	pulumi.Input

	ToActiveDirectoryConnectorPropertiesOutput() ActiveDirectoryConnectorPropertiesOutput
	ToActiveDirectoryConnectorPropertiesOutputWithContext(context.Context) ActiveDirectoryConnectorPropertiesOutput
}

type ActiveDirectoryConnectorPropertiesArgs struct {
	DomainServiceAccountLoginInformation BasicLoginInformationPtrInput          `pulumi:"domainServiceAccountLoginInformation"`
	Spec                                 ActiveDirectoryConnectorSpecInput      `pulumi:"spec"`
	Status                               ActiveDirectoryConnectorStatusPtrInput `pulumi:"status"`
}


func (val *ActiveDirectoryConnectorPropertiesArgs) Defaults() *ActiveDirectoryConnectorPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
}
func (ActiveDirectoryConnectorPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryConnectorProperties)(nil)).Elem()
}

func (i ActiveDirectoryConnectorPropertiesArgs) ToActiveDirectoryConnectorPropertiesOutput() ActiveDirectoryConnectorPropertiesOutput {
	return i.ToActiveDirectoryConnectorPropertiesOutputWithContext(context.Background())
}

func (i ActiveDirectoryConnectorPropertiesArgs) ToActiveDirectoryConnectorPropertiesOutputWithContext(ctx context.Context) ActiveDirectoryConnectorPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryConnectorPropertiesOutput)
}

type ActiveDirectoryConnectorPropertiesOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryConnectorPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryConnectorProperties)(nil)).Elem()
}

func (o ActiveDirectoryConnectorPropertiesOutput) ToActiveDirectoryConnectorPropertiesOutput() ActiveDirectoryConnectorPropertiesOutput {
	return o
}

func (o ActiveDirectoryConnectorPropertiesOutput) ToActiveDirectoryConnectorPropertiesOutputWithContext(ctx context.Context) ActiveDirectoryConnectorPropertiesOutput {
	return o
}

func (o ActiveDirectoryConnectorPropertiesOutput) DomainServiceAccountLoginInformation() BasicLoginInformationPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorProperties) *BasicLoginInformation {
		return v.DomainServiceAccountLoginInformation
	}).(BasicLoginInformationPtrOutput)
}

func (o ActiveDirectoryConnectorPropertiesOutput) Spec() ActiveDirectoryConnectorSpecOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorProperties) ActiveDirectoryConnectorSpec { return v.Spec }).(ActiveDirectoryConnectorSpecOutput)
}

func (o ActiveDirectoryConnectorPropertiesOutput) Status() ActiveDirectoryConnectorStatusPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorProperties) *ActiveDirectoryConnectorStatus { return v.Status }).(ActiveDirectoryConnectorStatusPtrOutput)
}

type ActiveDirectoryConnectorPropertiesResponse struct {
	DomainServiceAccountLoginInformation *BasicLoginInformationResponse          `pulumi:"domainServiceAccountLoginInformation"`
	ProvisioningState                    string                                  `pulumi:"provisioningState"`
	Spec                                 ActiveDirectoryConnectorSpecResponse    `pulumi:"spec"`
	Status                               *ActiveDirectoryConnectorStatusResponse `pulumi:"status"`
}


func (val *ActiveDirectoryConnectorPropertiesResponse) Defaults() *ActiveDirectoryConnectorPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Spec = *tmp.Spec.Defaults()

	return &tmp
}

type ActiveDirectoryConnectorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryConnectorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryConnectorPropertiesResponse)(nil)).Elem()
}

func (o ActiveDirectoryConnectorPropertiesResponseOutput) ToActiveDirectoryConnectorPropertiesResponseOutput() ActiveDirectoryConnectorPropertiesResponseOutput {
	return o
}

func (o ActiveDirectoryConnectorPropertiesResponseOutput) ToActiveDirectoryConnectorPropertiesResponseOutputWithContext(ctx context.Context) ActiveDirectoryConnectorPropertiesResponseOutput {
	return o
}

func (o ActiveDirectoryConnectorPropertiesResponseOutput) DomainServiceAccountLoginInformation() BasicLoginInformationResponsePtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorPropertiesResponse) *BasicLoginInformationResponse {
		return v.DomainServiceAccountLoginInformation
	}).(BasicLoginInformationResponsePtrOutput)
}

func (o ActiveDirectoryConnectorPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ActiveDirectoryConnectorPropertiesResponseOutput) Spec() ActiveDirectoryConnectorSpecResponseOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorPropertiesResponse) ActiveDirectoryConnectorSpecResponse { return v.Spec }).(ActiveDirectoryConnectorSpecResponseOutput)
}

func (o ActiveDirectoryConnectorPropertiesResponseOutput) Status() ActiveDirectoryConnectorStatusResponsePtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorPropertiesResponse) *ActiveDirectoryConnectorStatusResponse {
		return v.Status
	}).(ActiveDirectoryConnectorStatusResponsePtrOutput)
}

type ActiveDirectoryConnectorSpec struct {
	ActiveDirectory ActiveDirectoryConnectorDomainDetails `pulumi:"activeDirectory"`
	Dns             ActiveDirectoryConnectorDNSDetails    `pulumi:"dns"`
}


func (val *ActiveDirectoryConnectorSpec) Defaults() *ActiveDirectoryConnectorSpec {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ActiveDirectory = *tmp.ActiveDirectory.Defaults()

	tmp.Dns = *tmp.Dns.Defaults()

	return &tmp
}





type ActiveDirectoryConnectorSpecInput interface {
	pulumi.Input

	ToActiveDirectoryConnectorSpecOutput() ActiveDirectoryConnectorSpecOutput
	ToActiveDirectoryConnectorSpecOutputWithContext(context.Context) ActiveDirectoryConnectorSpecOutput
}

type ActiveDirectoryConnectorSpecArgs struct {
	ActiveDirectory ActiveDirectoryConnectorDomainDetailsInput `pulumi:"activeDirectory"`
	Dns             ActiveDirectoryConnectorDNSDetailsInput    `pulumi:"dns"`
}


func (val *ActiveDirectoryConnectorSpecArgs) Defaults() *ActiveDirectoryConnectorSpecArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
}
func (ActiveDirectoryConnectorSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryConnectorSpec)(nil)).Elem()
}

func (i ActiveDirectoryConnectorSpecArgs) ToActiveDirectoryConnectorSpecOutput() ActiveDirectoryConnectorSpecOutput {
	return i.ToActiveDirectoryConnectorSpecOutputWithContext(context.Background())
}

func (i ActiveDirectoryConnectorSpecArgs) ToActiveDirectoryConnectorSpecOutputWithContext(ctx context.Context) ActiveDirectoryConnectorSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryConnectorSpecOutput)
}

type ActiveDirectoryConnectorSpecOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryConnectorSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryConnectorSpec)(nil)).Elem()
}

func (o ActiveDirectoryConnectorSpecOutput) ToActiveDirectoryConnectorSpecOutput() ActiveDirectoryConnectorSpecOutput {
	return o
}

func (o ActiveDirectoryConnectorSpecOutput) ToActiveDirectoryConnectorSpecOutputWithContext(ctx context.Context) ActiveDirectoryConnectorSpecOutput {
	return o
}

func (o ActiveDirectoryConnectorSpecOutput) ActiveDirectory() ActiveDirectoryConnectorDomainDetailsOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorSpec) ActiveDirectoryConnectorDomainDetails { return v.ActiveDirectory }).(ActiveDirectoryConnectorDomainDetailsOutput)
}

func (o ActiveDirectoryConnectorSpecOutput) Dns() ActiveDirectoryConnectorDNSDetailsOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorSpec) ActiveDirectoryConnectorDNSDetails { return v.Dns }).(ActiveDirectoryConnectorDNSDetailsOutput)
}

type ActiveDirectoryConnectorSpecResponse struct {
	ActiveDirectory ActiveDirectoryConnectorDomainDetailsResponse `pulumi:"activeDirectory"`
	Dns             ActiveDirectoryConnectorDNSDetailsResponse    `pulumi:"dns"`
}


func (val *ActiveDirectoryConnectorSpecResponse) Defaults() *ActiveDirectoryConnectorSpecResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ActiveDirectory = *tmp.ActiveDirectory.Defaults()

	tmp.Dns = *tmp.Dns.Defaults()

	return &tmp
}

type ActiveDirectoryConnectorSpecResponseOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryConnectorSpecResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryConnectorSpecResponse)(nil)).Elem()
}

func (o ActiveDirectoryConnectorSpecResponseOutput) ToActiveDirectoryConnectorSpecResponseOutput() ActiveDirectoryConnectorSpecResponseOutput {
	return o
}

func (o ActiveDirectoryConnectorSpecResponseOutput) ToActiveDirectoryConnectorSpecResponseOutputWithContext(ctx context.Context) ActiveDirectoryConnectorSpecResponseOutput {
	return o
}

func (o ActiveDirectoryConnectorSpecResponseOutput) ActiveDirectory() ActiveDirectoryConnectorDomainDetailsResponseOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorSpecResponse) ActiveDirectoryConnectorDomainDetailsResponse {
		return v.ActiveDirectory
	}).(ActiveDirectoryConnectorDomainDetailsResponseOutput)
}

func (o ActiveDirectoryConnectorSpecResponseOutput) Dns() ActiveDirectoryConnectorDNSDetailsResponseOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorSpecResponse) ActiveDirectoryConnectorDNSDetailsResponse { return v.Dns }).(ActiveDirectoryConnectorDNSDetailsResponseOutput)
}

type ActiveDirectoryConnectorStatus struct {
	LastUpdateTime     *string  `pulumi:"lastUpdateTime"`
	ObservedGeneration *float64 `pulumi:"observedGeneration"`
	State              *string  `pulumi:"state"`
}





type ActiveDirectoryConnectorStatusInput interface {
	pulumi.Input

	ToActiveDirectoryConnectorStatusOutput() ActiveDirectoryConnectorStatusOutput
	ToActiveDirectoryConnectorStatusOutputWithContext(context.Context) ActiveDirectoryConnectorStatusOutput
}

type ActiveDirectoryConnectorStatusArgs struct {
	LastUpdateTime     pulumi.StringPtrInput  `pulumi:"lastUpdateTime"`
	ObservedGeneration pulumi.Float64PtrInput `pulumi:"observedGeneration"`
	State              pulumi.StringPtrInput  `pulumi:"state"`
}

func (ActiveDirectoryConnectorStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryConnectorStatus)(nil)).Elem()
}

func (i ActiveDirectoryConnectorStatusArgs) ToActiveDirectoryConnectorStatusOutput() ActiveDirectoryConnectorStatusOutput {
	return i.ToActiveDirectoryConnectorStatusOutputWithContext(context.Background())
}

func (i ActiveDirectoryConnectorStatusArgs) ToActiveDirectoryConnectorStatusOutputWithContext(ctx context.Context) ActiveDirectoryConnectorStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryConnectorStatusOutput)
}

func (i ActiveDirectoryConnectorStatusArgs) ToActiveDirectoryConnectorStatusPtrOutput() ActiveDirectoryConnectorStatusPtrOutput {
	return i.ToActiveDirectoryConnectorStatusPtrOutputWithContext(context.Background())
}

func (i ActiveDirectoryConnectorStatusArgs) ToActiveDirectoryConnectorStatusPtrOutputWithContext(ctx context.Context) ActiveDirectoryConnectorStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryConnectorStatusOutput).ToActiveDirectoryConnectorStatusPtrOutputWithContext(ctx)
}









type ActiveDirectoryConnectorStatusPtrInput interface {
	pulumi.Input

	ToActiveDirectoryConnectorStatusPtrOutput() ActiveDirectoryConnectorStatusPtrOutput
	ToActiveDirectoryConnectorStatusPtrOutputWithContext(context.Context) ActiveDirectoryConnectorStatusPtrOutput
}

type activeDirectoryConnectorStatusPtrType ActiveDirectoryConnectorStatusArgs

func ActiveDirectoryConnectorStatusPtr(v *ActiveDirectoryConnectorStatusArgs) ActiveDirectoryConnectorStatusPtrInput {
	return (*activeDirectoryConnectorStatusPtrType)(v)
}

func (*activeDirectoryConnectorStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryConnectorStatus)(nil)).Elem()
}

func (i *activeDirectoryConnectorStatusPtrType) ToActiveDirectoryConnectorStatusPtrOutput() ActiveDirectoryConnectorStatusPtrOutput {
	return i.ToActiveDirectoryConnectorStatusPtrOutputWithContext(context.Background())
}

func (i *activeDirectoryConnectorStatusPtrType) ToActiveDirectoryConnectorStatusPtrOutputWithContext(ctx context.Context) ActiveDirectoryConnectorStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryConnectorStatusPtrOutput)
}

type ActiveDirectoryConnectorStatusOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryConnectorStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryConnectorStatus)(nil)).Elem()
}

func (o ActiveDirectoryConnectorStatusOutput) ToActiveDirectoryConnectorStatusOutput() ActiveDirectoryConnectorStatusOutput {
	return o
}

func (o ActiveDirectoryConnectorStatusOutput) ToActiveDirectoryConnectorStatusOutputWithContext(ctx context.Context) ActiveDirectoryConnectorStatusOutput {
	return o
}

func (o ActiveDirectoryConnectorStatusOutput) ToActiveDirectoryConnectorStatusPtrOutput() ActiveDirectoryConnectorStatusPtrOutput {
	return o.ToActiveDirectoryConnectorStatusPtrOutputWithContext(context.Background())
}

func (o ActiveDirectoryConnectorStatusOutput) ToActiveDirectoryConnectorStatusPtrOutputWithContext(ctx context.Context) ActiveDirectoryConnectorStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActiveDirectoryConnectorStatus) *ActiveDirectoryConnectorStatus {
		return &v
	}).(ActiveDirectoryConnectorStatusPtrOutput)
}

func (o ActiveDirectoryConnectorStatusOutput) LastUpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorStatus) *string { return v.LastUpdateTime }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryConnectorStatusOutput) ObservedGeneration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorStatus) *float64 { return v.ObservedGeneration }).(pulumi.Float64PtrOutput)
}

func (o ActiveDirectoryConnectorStatusOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorStatus) *string { return v.State }).(pulumi.StringPtrOutput)
}

type ActiveDirectoryConnectorStatusPtrOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryConnectorStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryConnectorStatus)(nil)).Elem()
}

func (o ActiveDirectoryConnectorStatusPtrOutput) ToActiveDirectoryConnectorStatusPtrOutput() ActiveDirectoryConnectorStatusPtrOutput {
	return o
}

func (o ActiveDirectoryConnectorStatusPtrOutput) ToActiveDirectoryConnectorStatusPtrOutputWithContext(ctx context.Context) ActiveDirectoryConnectorStatusPtrOutput {
	return o
}

func (o ActiveDirectoryConnectorStatusPtrOutput) Elem() ActiveDirectoryConnectorStatusOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnectorStatus) ActiveDirectoryConnectorStatus {
		if v != nil {
			return *v
		}
		var ret ActiveDirectoryConnectorStatus
		return ret
	}).(ActiveDirectoryConnectorStatusOutput)
}

func (o ActiveDirectoryConnectorStatusPtrOutput) LastUpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnectorStatus) *string {
		if v == nil {
			return nil
		}
		return v.LastUpdateTime
	}).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryConnectorStatusPtrOutput) ObservedGeneration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnectorStatus) *float64 {
		if v == nil {
			return nil
		}
		return v.ObservedGeneration
	}).(pulumi.Float64PtrOutput)
}

func (o ActiveDirectoryConnectorStatusPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnectorStatus) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type ActiveDirectoryConnectorStatusResponse struct {
	LastUpdateTime     *string  `pulumi:"lastUpdateTime"`
	ObservedGeneration *float64 `pulumi:"observedGeneration"`
	State              *string  `pulumi:"state"`
}

type ActiveDirectoryConnectorStatusResponseOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryConnectorStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryConnectorStatusResponse)(nil)).Elem()
}

func (o ActiveDirectoryConnectorStatusResponseOutput) ToActiveDirectoryConnectorStatusResponseOutput() ActiveDirectoryConnectorStatusResponseOutput {
	return o
}

func (o ActiveDirectoryConnectorStatusResponseOutput) ToActiveDirectoryConnectorStatusResponseOutputWithContext(ctx context.Context) ActiveDirectoryConnectorStatusResponseOutput {
	return o
}

func (o ActiveDirectoryConnectorStatusResponseOutput) LastUpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorStatusResponse) *string { return v.LastUpdateTime }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryConnectorStatusResponseOutput) ObservedGeneration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorStatusResponse) *float64 { return v.ObservedGeneration }).(pulumi.Float64PtrOutput)
}

func (o ActiveDirectoryConnectorStatusResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryConnectorStatusResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type ActiveDirectoryConnectorStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryConnectorStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryConnectorStatusResponse)(nil)).Elem()
}

func (o ActiveDirectoryConnectorStatusResponsePtrOutput) ToActiveDirectoryConnectorStatusResponsePtrOutput() ActiveDirectoryConnectorStatusResponsePtrOutput {
	return o
}

func (o ActiveDirectoryConnectorStatusResponsePtrOutput) ToActiveDirectoryConnectorStatusResponsePtrOutputWithContext(ctx context.Context) ActiveDirectoryConnectorStatusResponsePtrOutput {
	return o
}

func (o ActiveDirectoryConnectorStatusResponsePtrOutput) Elem() ActiveDirectoryConnectorStatusResponseOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnectorStatusResponse) ActiveDirectoryConnectorStatusResponse {
		if v != nil {
			return *v
		}
		var ret ActiveDirectoryConnectorStatusResponse
		return ret
	}).(ActiveDirectoryConnectorStatusResponseOutput)
}

func (o ActiveDirectoryConnectorStatusResponsePtrOutput) LastUpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnectorStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastUpdateTime
	}).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryConnectorStatusResponsePtrOutput) ObservedGeneration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnectorStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.ObservedGeneration
	}).(pulumi.Float64PtrOutput)
}

func (o ActiveDirectoryConnectorStatusResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnectorStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type ActiveDirectoryDomainController struct {
	Hostname string `pulumi:"hostname"`
}





type ActiveDirectoryDomainControllerInput interface {
	pulumi.Input

	ToActiveDirectoryDomainControllerOutput() ActiveDirectoryDomainControllerOutput
	ToActiveDirectoryDomainControllerOutputWithContext(context.Context) ActiveDirectoryDomainControllerOutput
}

type ActiveDirectoryDomainControllerArgs struct {
	Hostname pulumi.StringInput `pulumi:"hostname"`
}

func (ActiveDirectoryDomainControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryDomainController)(nil)).Elem()
}

func (i ActiveDirectoryDomainControllerArgs) ToActiveDirectoryDomainControllerOutput() ActiveDirectoryDomainControllerOutput {
	return i.ToActiveDirectoryDomainControllerOutputWithContext(context.Background())
}

func (i ActiveDirectoryDomainControllerArgs) ToActiveDirectoryDomainControllerOutputWithContext(ctx context.Context) ActiveDirectoryDomainControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryDomainControllerOutput)
}

func (i ActiveDirectoryDomainControllerArgs) ToActiveDirectoryDomainControllerPtrOutput() ActiveDirectoryDomainControllerPtrOutput {
	return i.ToActiveDirectoryDomainControllerPtrOutputWithContext(context.Background())
}

func (i ActiveDirectoryDomainControllerArgs) ToActiveDirectoryDomainControllerPtrOutputWithContext(ctx context.Context) ActiveDirectoryDomainControllerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryDomainControllerOutput).ToActiveDirectoryDomainControllerPtrOutputWithContext(ctx)
}









type ActiveDirectoryDomainControllerPtrInput interface {
	pulumi.Input

	ToActiveDirectoryDomainControllerPtrOutput() ActiveDirectoryDomainControllerPtrOutput
	ToActiveDirectoryDomainControllerPtrOutputWithContext(context.Context) ActiveDirectoryDomainControllerPtrOutput
}

type activeDirectoryDomainControllerPtrType ActiveDirectoryDomainControllerArgs

func ActiveDirectoryDomainControllerPtr(v *ActiveDirectoryDomainControllerArgs) ActiveDirectoryDomainControllerPtrInput {
	return (*activeDirectoryDomainControllerPtrType)(v)
}

func (*activeDirectoryDomainControllerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryDomainController)(nil)).Elem()
}

func (i *activeDirectoryDomainControllerPtrType) ToActiveDirectoryDomainControllerPtrOutput() ActiveDirectoryDomainControllerPtrOutput {
	return i.ToActiveDirectoryDomainControllerPtrOutputWithContext(context.Background())
}

func (i *activeDirectoryDomainControllerPtrType) ToActiveDirectoryDomainControllerPtrOutputWithContext(ctx context.Context) ActiveDirectoryDomainControllerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryDomainControllerPtrOutput)
}





type ActiveDirectoryDomainControllerArrayInput interface {
	pulumi.Input

	ToActiveDirectoryDomainControllerArrayOutput() ActiveDirectoryDomainControllerArrayOutput
	ToActiveDirectoryDomainControllerArrayOutputWithContext(context.Context) ActiveDirectoryDomainControllerArrayOutput
}

type ActiveDirectoryDomainControllerArray []ActiveDirectoryDomainControllerInput

func (ActiveDirectoryDomainControllerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveDirectoryDomainController)(nil)).Elem()
}

func (i ActiveDirectoryDomainControllerArray) ToActiveDirectoryDomainControllerArrayOutput() ActiveDirectoryDomainControllerArrayOutput {
	return i.ToActiveDirectoryDomainControllerArrayOutputWithContext(context.Background())
}

func (i ActiveDirectoryDomainControllerArray) ToActiveDirectoryDomainControllerArrayOutputWithContext(ctx context.Context) ActiveDirectoryDomainControllerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryDomainControllerArrayOutput)
}

type ActiveDirectoryDomainControllerOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryDomainControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryDomainController)(nil)).Elem()
}

func (o ActiveDirectoryDomainControllerOutput) ToActiveDirectoryDomainControllerOutput() ActiveDirectoryDomainControllerOutput {
	return o
}

func (o ActiveDirectoryDomainControllerOutput) ToActiveDirectoryDomainControllerOutputWithContext(ctx context.Context) ActiveDirectoryDomainControllerOutput {
	return o
}

func (o ActiveDirectoryDomainControllerOutput) ToActiveDirectoryDomainControllerPtrOutput() ActiveDirectoryDomainControllerPtrOutput {
	return o.ToActiveDirectoryDomainControllerPtrOutputWithContext(context.Background())
}

func (o ActiveDirectoryDomainControllerOutput) ToActiveDirectoryDomainControllerPtrOutputWithContext(ctx context.Context) ActiveDirectoryDomainControllerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActiveDirectoryDomainController) *ActiveDirectoryDomainController {
		return &v
	}).(ActiveDirectoryDomainControllerPtrOutput)
}

func (o ActiveDirectoryDomainControllerOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryDomainController) string { return v.Hostname }).(pulumi.StringOutput)
}

type ActiveDirectoryDomainControllerPtrOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryDomainControllerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryDomainController)(nil)).Elem()
}

func (o ActiveDirectoryDomainControllerPtrOutput) ToActiveDirectoryDomainControllerPtrOutput() ActiveDirectoryDomainControllerPtrOutput {
	return o
}

func (o ActiveDirectoryDomainControllerPtrOutput) ToActiveDirectoryDomainControllerPtrOutputWithContext(ctx context.Context) ActiveDirectoryDomainControllerPtrOutput {
	return o
}

func (o ActiveDirectoryDomainControllerPtrOutput) Elem() ActiveDirectoryDomainControllerOutput {
	return o.ApplyT(func(v *ActiveDirectoryDomainController) ActiveDirectoryDomainController {
		if v != nil {
			return *v
		}
		var ret ActiveDirectoryDomainController
		return ret
	}).(ActiveDirectoryDomainControllerOutput)
}

func (o ActiveDirectoryDomainControllerPtrOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryDomainController) *string {
		if v == nil {
			return nil
		}
		return &v.Hostname
	}).(pulumi.StringPtrOutput)
}

type ActiveDirectoryDomainControllerArrayOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryDomainControllerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveDirectoryDomainController)(nil)).Elem()
}

func (o ActiveDirectoryDomainControllerArrayOutput) ToActiveDirectoryDomainControllerArrayOutput() ActiveDirectoryDomainControllerArrayOutput {
	return o
}

func (o ActiveDirectoryDomainControllerArrayOutput) ToActiveDirectoryDomainControllerArrayOutputWithContext(ctx context.Context) ActiveDirectoryDomainControllerArrayOutput {
	return o
}

func (o ActiveDirectoryDomainControllerArrayOutput) Index(i pulumi.IntInput) ActiveDirectoryDomainControllerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActiveDirectoryDomainController {
		return vs[0].([]ActiveDirectoryDomainController)[vs[1].(int)]
	}).(ActiveDirectoryDomainControllerOutput)
}

type ActiveDirectoryDomainControllerResponse struct {
	Hostname string `pulumi:"hostname"`
}

type ActiveDirectoryDomainControllerResponseOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryDomainControllerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryDomainControllerResponse)(nil)).Elem()
}

func (o ActiveDirectoryDomainControllerResponseOutput) ToActiveDirectoryDomainControllerResponseOutput() ActiveDirectoryDomainControllerResponseOutput {
	return o
}

func (o ActiveDirectoryDomainControllerResponseOutput) ToActiveDirectoryDomainControllerResponseOutputWithContext(ctx context.Context) ActiveDirectoryDomainControllerResponseOutput {
	return o
}

func (o ActiveDirectoryDomainControllerResponseOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryDomainControllerResponse) string { return v.Hostname }).(pulumi.StringOutput)
}

type ActiveDirectoryDomainControllerResponsePtrOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryDomainControllerResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryDomainControllerResponse)(nil)).Elem()
}

func (o ActiveDirectoryDomainControllerResponsePtrOutput) ToActiveDirectoryDomainControllerResponsePtrOutput() ActiveDirectoryDomainControllerResponsePtrOutput {
	return o
}

func (o ActiveDirectoryDomainControllerResponsePtrOutput) ToActiveDirectoryDomainControllerResponsePtrOutputWithContext(ctx context.Context) ActiveDirectoryDomainControllerResponsePtrOutput {
	return o
}

func (o ActiveDirectoryDomainControllerResponsePtrOutput) Elem() ActiveDirectoryDomainControllerResponseOutput {
	return o.ApplyT(func(v *ActiveDirectoryDomainControllerResponse) ActiveDirectoryDomainControllerResponse {
		if v != nil {
			return *v
		}
		var ret ActiveDirectoryDomainControllerResponse
		return ret
	}).(ActiveDirectoryDomainControllerResponseOutput)
}

func (o ActiveDirectoryDomainControllerResponsePtrOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryDomainControllerResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Hostname
	}).(pulumi.StringPtrOutput)
}

type ActiveDirectoryDomainControllerResponseArrayOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryDomainControllerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveDirectoryDomainControllerResponse)(nil)).Elem()
}

func (o ActiveDirectoryDomainControllerResponseArrayOutput) ToActiveDirectoryDomainControllerResponseArrayOutput() ActiveDirectoryDomainControllerResponseArrayOutput {
	return o
}

func (o ActiveDirectoryDomainControllerResponseArrayOutput) ToActiveDirectoryDomainControllerResponseArrayOutputWithContext(ctx context.Context) ActiveDirectoryDomainControllerResponseArrayOutput {
	return o
}

func (o ActiveDirectoryDomainControllerResponseArrayOutput) Index(i pulumi.IntInput) ActiveDirectoryDomainControllerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActiveDirectoryDomainControllerResponse {
		return vs[0].([]ActiveDirectoryDomainControllerResponse)[vs[1].(int)]
	}).(ActiveDirectoryDomainControllerResponseOutput)
}

type ActiveDirectoryDomainControllers struct {
	PrimaryDomainController    *ActiveDirectoryDomainController  `pulumi:"primaryDomainController"`
	SecondaryDomainControllers []ActiveDirectoryDomainController `pulumi:"secondaryDomainControllers"`
}





type ActiveDirectoryDomainControllersInput interface {
	pulumi.Input

	ToActiveDirectoryDomainControllersOutput() ActiveDirectoryDomainControllersOutput
	ToActiveDirectoryDomainControllersOutputWithContext(context.Context) ActiveDirectoryDomainControllersOutput
}

type ActiveDirectoryDomainControllersArgs struct {
	PrimaryDomainController    ActiveDirectoryDomainControllerPtrInput   `pulumi:"primaryDomainController"`
	SecondaryDomainControllers ActiveDirectoryDomainControllerArrayInput `pulumi:"secondaryDomainControllers"`
}

func (ActiveDirectoryDomainControllersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryDomainControllers)(nil)).Elem()
}

func (i ActiveDirectoryDomainControllersArgs) ToActiveDirectoryDomainControllersOutput() ActiveDirectoryDomainControllersOutput {
	return i.ToActiveDirectoryDomainControllersOutputWithContext(context.Background())
}

func (i ActiveDirectoryDomainControllersArgs) ToActiveDirectoryDomainControllersOutputWithContext(ctx context.Context) ActiveDirectoryDomainControllersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryDomainControllersOutput)
}

type ActiveDirectoryDomainControllersOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryDomainControllersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryDomainControllers)(nil)).Elem()
}

func (o ActiveDirectoryDomainControllersOutput) ToActiveDirectoryDomainControllersOutput() ActiveDirectoryDomainControllersOutput {
	return o
}

func (o ActiveDirectoryDomainControllersOutput) ToActiveDirectoryDomainControllersOutputWithContext(ctx context.Context) ActiveDirectoryDomainControllersOutput {
	return o
}

func (o ActiveDirectoryDomainControllersOutput) PrimaryDomainController() ActiveDirectoryDomainControllerPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryDomainControllers) *ActiveDirectoryDomainController {
		return v.PrimaryDomainController
	}).(ActiveDirectoryDomainControllerPtrOutput)
}

func (o ActiveDirectoryDomainControllersOutput) SecondaryDomainControllers() ActiveDirectoryDomainControllerArrayOutput {
	return o.ApplyT(func(v ActiveDirectoryDomainControllers) []ActiveDirectoryDomainController {
		return v.SecondaryDomainControllers
	}).(ActiveDirectoryDomainControllerArrayOutput)
}

type ActiveDirectoryDomainControllersResponse struct {
	PrimaryDomainController    *ActiveDirectoryDomainControllerResponse  `pulumi:"primaryDomainController"`
	SecondaryDomainControllers []ActiveDirectoryDomainControllerResponse `pulumi:"secondaryDomainControllers"`
}

type ActiveDirectoryDomainControllersResponseOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryDomainControllersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryDomainControllersResponse)(nil)).Elem()
}

func (o ActiveDirectoryDomainControllersResponseOutput) ToActiveDirectoryDomainControllersResponseOutput() ActiveDirectoryDomainControllersResponseOutput {
	return o
}

func (o ActiveDirectoryDomainControllersResponseOutput) ToActiveDirectoryDomainControllersResponseOutputWithContext(ctx context.Context) ActiveDirectoryDomainControllersResponseOutput {
	return o
}

func (o ActiveDirectoryDomainControllersResponseOutput) PrimaryDomainController() ActiveDirectoryDomainControllerResponsePtrOutput {
	return o.ApplyT(func(v ActiveDirectoryDomainControllersResponse) *ActiveDirectoryDomainControllerResponse {
		return v.PrimaryDomainController
	}).(ActiveDirectoryDomainControllerResponsePtrOutput)
}

func (o ActiveDirectoryDomainControllersResponseOutput) SecondaryDomainControllers() ActiveDirectoryDomainControllerResponseArrayOutput {
	return o.ApplyT(func(v ActiveDirectoryDomainControllersResponse) []ActiveDirectoryDomainControllerResponse {
		return v.SecondaryDomainControllers
	}).(ActiveDirectoryDomainControllerResponseArrayOutput)
}

type ActiveDirectoryInformation struct {
	KeytabInformation *KeytabInformation `pulumi:"keytabInformation"`
}





type ActiveDirectoryInformationInput interface {
	pulumi.Input

	ToActiveDirectoryInformationOutput() ActiveDirectoryInformationOutput
	ToActiveDirectoryInformationOutputWithContext(context.Context) ActiveDirectoryInformationOutput
}

type ActiveDirectoryInformationArgs struct {
	KeytabInformation KeytabInformationPtrInput `pulumi:"keytabInformation"`
}

func (ActiveDirectoryInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryInformation)(nil)).Elem()
}

func (i ActiveDirectoryInformationArgs) ToActiveDirectoryInformationOutput() ActiveDirectoryInformationOutput {
	return i.ToActiveDirectoryInformationOutputWithContext(context.Background())
}

func (i ActiveDirectoryInformationArgs) ToActiveDirectoryInformationOutputWithContext(ctx context.Context) ActiveDirectoryInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryInformationOutput)
}

func (i ActiveDirectoryInformationArgs) ToActiveDirectoryInformationPtrOutput() ActiveDirectoryInformationPtrOutput {
	return i.ToActiveDirectoryInformationPtrOutputWithContext(context.Background())
}

func (i ActiveDirectoryInformationArgs) ToActiveDirectoryInformationPtrOutputWithContext(ctx context.Context) ActiveDirectoryInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryInformationOutput).ToActiveDirectoryInformationPtrOutputWithContext(ctx)
}









type ActiveDirectoryInformationPtrInput interface {
	pulumi.Input

	ToActiveDirectoryInformationPtrOutput() ActiveDirectoryInformationPtrOutput
	ToActiveDirectoryInformationPtrOutputWithContext(context.Context) ActiveDirectoryInformationPtrOutput
}

type activeDirectoryInformationPtrType ActiveDirectoryInformationArgs

func ActiveDirectoryInformationPtr(v *ActiveDirectoryInformationArgs) ActiveDirectoryInformationPtrInput {
	return (*activeDirectoryInformationPtrType)(v)
}

func (*activeDirectoryInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryInformation)(nil)).Elem()
}

func (i *activeDirectoryInformationPtrType) ToActiveDirectoryInformationPtrOutput() ActiveDirectoryInformationPtrOutput {
	return i.ToActiveDirectoryInformationPtrOutputWithContext(context.Background())
}

func (i *activeDirectoryInformationPtrType) ToActiveDirectoryInformationPtrOutputWithContext(ctx context.Context) ActiveDirectoryInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryInformationPtrOutput)
}

type ActiveDirectoryInformationOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryInformation)(nil)).Elem()
}

func (o ActiveDirectoryInformationOutput) ToActiveDirectoryInformationOutput() ActiveDirectoryInformationOutput {
	return o
}

func (o ActiveDirectoryInformationOutput) ToActiveDirectoryInformationOutputWithContext(ctx context.Context) ActiveDirectoryInformationOutput {
	return o
}

func (o ActiveDirectoryInformationOutput) ToActiveDirectoryInformationPtrOutput() ActiveDirectoryInformationPtrOutput {
	return o.ToActiveDirectoryInformationPtrOutputWithContext(context.Background())
}

func (o ActiveDirectoryInformationOutput) ToActiveDirectoryInformationPtrOutputWithContext(ctx context.Context) ActiveDirectoryInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActiveDirectoryInformation) *ActiveDirectoryInformation {
		return &v
	}).(ActiveDirectoryInformationPtrOutput)
}

func (o ActiveDirectoryInformationOutput) KeytabInformation() KeytabInformationPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryInformation) *KeytabInformation { return v.KeytabInformation }).(KeytabInformationPtrOutput)
}

type ActiveDirectoryInformationPtrOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryInformation)(nil)).Elem()
}

func (o ActiveDirectoryInformationPtrOutput) ToActiveDirectoryInformationPtrOutput() ActiveDirectoryInformationPtrOutput {
	return o
}

func (o ActiveDirectoryInformationPtrOutput) ToActiveDirectoryInformationPtrOutputWithContext(ctx context.Context) ActiveDirectoryInformationPtrOutput {
	return o
}

func (o ActiveDirectoryInformationPtrOutput) Elem() ActiveDirectoryInformationOutput {
	return o.ApplyT(func(v *ActiveDirectoryInformation) ActiveDirectoryInformation {
		if v != nil {
			return *v
		}
		var ret ActiveDirectoryInformation
		return ret
	}).(ActiveDirectoryInformationOutput)
}

func (o ActiveDirectoryInformationPtrOutput) KeytabInformation() KeytabInformationPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryInformation) *KeytabInformation {
		if v == nil {
			return nil
		}
		return v.KeytabInformation
	}).(KeytabInformationPtrOutput)
}

type BasicLoginInformation struct {
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}





type BasicLoginInformationInput interface {
	pulumi.Input

	ToBasicLoginInformationOutput() BasicLoginInformationOutput
	ToBasicLoginInformationOutputWithContext(context.Context) BasicLoginInformationOutput
}

type BasicLoginInformationArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (BasicLoginInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BasicLoginInformation)(nil)).Elem()
}

func (i BasicLoginInformationArgs) ToBasicLoginInformationOutput() BasicLoginInformationOutput {
	return i.ToBasicLoginInformationOutputWithContext(context.Background())
}

func (i BasicLoginInformationArgs) ToBasicLoginInformationOutputWithContext(ctx context.Context) BasicLoginInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicLoginInformationOutput)
}

func (i BasicLoginInformationArgs) ToBasicLoginInformationPtrOutput() BasicLoginInformationPtrOutput {
	return i.ToBasicLoginInformationPtrOutputWithContext(context.Background())
}

func (i BasicLoginInformationArgs) ToBasicLoginInformationPtrOutputWithContext(ctx context.Context) BasicLoginInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicLoginInformationOutput).ToBasicLoginInformationPtrOutputWithContext(ctx)
}









type BasicLoginInformationPtrInput interface {
	pulumi.Input

	ToBasicLoginInformationPtrOutput() BasicLoginInformationPtrOutput
	ToBasicLoginInformationPtrOutputWithContext(context.Context) BasicLoginInformationPtrOutput
}

type basicLoginInformationPtrType BasicLoginInformationArgs

func BasicLoginInformationPtr(v *BasicLoginInformationArgs) BasicLoginInformationPtrInput {
	return (*basicLoginInformationPtrType)(v)
}

func (*basicLoginInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicLoginInformation)(nil)).Elem()
}

func (i *basicLoginInformationPtrType) ToBasicLoginInformationPtrOutput() BasicLoginInformationPtrOutput {
	return i.ToBasicLoginInformationPtrOutputWithContext(context.Background())
}

func (i *basicLoginInformationPtrType) ToBasicLoginInformationPtrOutputWithContext(ctx context.Context) BasicLoginInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicLoginInformationPtrOutput)
}

type BasicLoginInformationOutput struct{ *pulumi.OutputState }

func (BasicLoginInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BasicLoginInformation)(nil)).Elem()
}

func (o BasicLoginInformationOutput) ToBasicLoginInformationOutput() BasicLoginInformationOutput {
	return o
}

func (o BasicLoginInformationOutput) ToBasicLoginInformationOutputWithContext(ctx context.Context) BasicLoginInformationOutput {
	return o
}

func (o BasicLoginInformationOutput) ToBasicLoginInformationPtrOutput() BasicLoginInformationPtrOutput {
	return o.ToBasicLoginInformationPtrOutputWithContext(context.Background())
}

func (o BasicLoginInformationOutput) ToBasicLoginInformationPtrOutputWithContext(ctx context.Context) BasicLoginInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BasicLoginInformation) *BasicLoginInformation {
		return &v
	}).(BasicLoginInformationPtrOutput)
}

func (o BasicLoginInformationOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BasicLoginInformation) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o BasicLoginInformationOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BasicLoginInformation) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type BasicLoginInformationPtrOutput struct{ *pulumi.OutputState }

func (BasicLoginInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicLoginInformation)(nil)).Elem()
}

func (o BasicLoginInformationPtrOutput) ToBasicLoginInformationPtrOutput() BasicLoginInformationPtrOutput {
	return o
}

func (o BasicLoginInformationPtrOutput) ToBasicLoginInformationPtrOutputWithContext(ctx context.Context) BasicLoginInformationPtrOutput {
	return o
}

func (o BasicLoginInformationPtrOutput) Elem() BasicLoginInformationOutput {
	return o.ApplyT(func(v *BasicLoginInformation) BasicLoginInformation {
		if v != nil {
			return *v
		}
		var ret BasicLoginInformation
		return ret
	}).(BasicLoginInformationOutput)
}

func (o BasicLoginInformationPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicLoginInformation) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o BasicLoginInformationPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicLoginInformation) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type BasicLoginInformationResponse struct {
	Username *string `pulumi:"username"`
}

type BasicLoginInformationResponseOutput struct{ *pulumi.OutputState }

func (BasicLoginInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BasicLoginInformationResponse)(nil)).Elem()
}

func (o BasicLoginInformationResponseOutput) ToBasicLoginInformationResponseOutput() BasicLoginInformationResponseOutput {
	return o
}

func (o BasicLoginInformationResponseOutput) ToBasicLoginInformationResponseOutputWithContext(ctx context.Context) BasicLoginInformationResponseOutput {
	return o
}

func (o BasicLoginInformationResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BasicLoginInformationResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type BasicLoginInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (BasicLoginInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicLoginInformationResponse)(nil)).Elem()
}

func (o BasicLoginInformationResponsePtrOutput) ToBasicLoginInformationResponsePtrOutput() BasicLoginInformationResponsePtrOutput {
	return o
}

func (o BasicLoginInformationResponsePtrOutput) ToBasicLoginInformationResponsePtrOutputWithContext(ctx context.Context) BasicLoginInformationResponsePtrOutput {
	return o
}

func (o BasicLoginInformationResponsePtrOutput) Elem() BasicLoginInformationResponseOutput {
	return o.ApplyT(func(v *BasicLoginInformationResponse) BasicLoginInformationResponse {
		if v != nil {
			return *v
		}
		var ret BasicLoginInformationResponse
		return ret
	}).(BasicLoginInformationResponseOutput)
}

func (o BasicLoginInformationResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicLoginInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type DataControllerProperties struct {
	BasicLoginInformation       *BasicLoginInformation       `pulumi:"basicLoginInformation"`
	ClusterId                   *string                      `pulumi:"clusterId"`
	ExtensionId                 *string                      `pulumi:"extensionId"`
	Infrastructure              *Infrastructure              `pulumi:"infrastructure"`
	K8sRaw                      interface{}                  `pulumi:"k8sRaw"`
	LastUploadedDate            *string                      `pulumi:"lastUploadedDate"`
	LogAnalyticsWorkspaceConfig *LogAnalyticsWorkspaceConfig `pulumi:"logAnalyticsWorkspaceConfig"`
	LogsDashboardCredential     *BasicLoginInformation       `pulumi:"logsDashboardCredential"`
	MetricsDashboardCredential  *BasicLoginInformation       `pulumi:"metricsDashboardCredential"`
	OnPremiseProperty           *OnPremiseProperty           `pulumi:"onPremiseProperty"`
	UploadServicePrincipal      *UploadServicePrincipal      `pulumi:"uploadServicePrincipal"`
	UploadWatermark             *UploadWatermark             `pulumi:"uploadWatermark"`
}


func (val *DataControllerProperties) Defaults() *DataControllerProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Infrastructure) {
		infrastructure_ := Infrastructure("other")
		tmp.Infrastructure = &infrastructure_
	}
	return &tmp
}





type DataControllerPropertiesInput interface {
	pulumi.Input

	ToDataControllerPropertiesOutput() DataControllerPropertiesOutput
	ToDataControllerPropertiesOutputWithContext(context.Context) DataControllerPropertiesOutput
}

type DataControllerPropertiesArgs struct {
	BasicLoginInformation       BasicLoginInformationPtrInput       `pulumi:"basicLoginInformation"`
	ClusterId                   pulumi.StringPtrInput               `pulumi:"clusterId"`
	ExtensionId                 pulumi.StringPtrInput               `pulumi:"extensionId"`
	Infrastructure              InfrastructurePtrInput              `pulumi:"infrastructure"`
	K8sRaw                      pulumi.Input                        `pulumi:"k8sRaw"`
	LastUploadedDate            pulumi.StringPtrInput               `pulumi:"lastUploadedDate"`
	LogAnalyticsWorkspaceConfig LogAnalyticsWorkspaceConfigPtrInput `pulumi:"logAnalyticsWorkspaceConfig"`
	LogsDashboardCredential     BasicLoginInformationPtrInput       `pulumi:"logsDashboardCredential"`
	MetricsDashboardCredential  BasicLoginInformationPtrInput       `pulumi:"metricsDashboardCredential"`
	OnPremiseProperty           OnPremisePropertyPtrInput           `pulumi:"onPremiseProperty"`
	UploadServicePrincipal      UploadServicePrincipalPtrInput      `pulumi:"uploadServicePrincipal"`
	UploadWatermark             UploadWatermarkPtrInput             `pulumi:"uploadWatermark"`
}


func (val *DataControllerPropertiesArgs) Defaults() *DataControllerPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Infrastructure) {
		tmp.Infrastructure = Infrastructure("other")
	}
	return &tmp
}
func (DataControllerPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataControllerProperties)(nil)).Elem()
}

func (i DataControllerPropertiesArgs) ToDataControllerPropertiesOutput() DataControllerPropertiesOutput {
	return i.ToDataControllerPropertiesOutputWithContext(context.Background())
}

func (i DataControllerPropertiesArgs) ToDataControllerPropertiesOutputWithContext(ctx context.Context) DataControllerPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataControllerPropertiesOutput)
}

type DataControllerPropertiesOutput struct{ *pulumi.OutputState }

func (DataControllerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataControllerProperties)(nil)).Elem()
}

func (o DataControllerPropertiesOutput) ToDataControllerPropertiesOutput() DataControllerPropertiesOutput {
	return o
}

func (o DataControllerPropertiesOutput) ToDataControllerPropertiesOutputWithContext(ctx context.Context) DataControllerPropertiesOutput {
	return o
}

func (o DataControllerPropertiesOutput) BasicLoginInformation() BasicLoginInformationPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *BasicLoginInformation { return v.BasicLoginInformation }).(BasicLoginInformationPtrOutput)
}

func (o DataControllerPropertiesOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o DataControllerPropertiesOutput) ExtensionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *string { return v.ExtensionId }).(pulumi.StringPtrOutput)
}

func (o DataControllerPropertiesOutput) Infrastructure() InfrastructurePtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *Infrastructure { return v.Infrastructure }).(InfrastructurePtrOutput)
}

func (o DataControllerPropertiesOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v DataControllerProperties) interface{} { return v.K8sRaw }).(pulumi.AnyOutput)
}

func (o DataControllerPropertiesOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *string { return v.LastUploadedDate }).(pulumi.StringPtrOutput)
}

func (o DataControllerPropertiesOutput) LogAnalyticsWorkspaceConfig() LogAnalyticsWorkspaceConfigPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *LogAnalyticsWorkspaceConfig { return v.LogAnalyticsWorkspaceConfig }).(LogAnalyticsWorkspaceConfigPtrOutput)
}

func (o DataControllerPropertiesOutput) LogsDashboardCredential() BasicLoginInformationPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *BasicLoginInformation { return v.LogsDashboardCredential }).(BasicLoginInformationPtrOutput)
}

func (o DataControllerPropertiesOutput) MetricsDashboardCredential() BasicLoginInformationPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *BasicLoginInformation { return v.MetricsDashboardCredential }).(BasicLoginInformationPtrOutput)
}

func (o DataControllerPropertiesOutput) OnPremiseProperty() OnPremisePropertyPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *OnPremiseProperty { return v.OnPremiseProperty }).(OnPremisePropertyPtrOutput)
}

func (o DataControllerPropertiesOutput) UploadServicePrincipal() UploadServicePrincipalPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *UploadServicePrincipal { return v.UploadServicePrincipal }).(UploadServicePrincipalPtrOutput)
}

func (o DataControllerPropertiesOutput) UploadWatermark() UploadWatermarkPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *UploadWatermark { return v.UploadWatermark }).(UploadWatermarkPtrOutput)
}

type DataControllerPropertiesResponse struct {
	BasicLoginInformation       *BasicLoginInformationResponse       `pulumi:"basicLoginInformation"`
	ClusterId                   *string                              `pulumi:"clusterId"`
	ExtensionId                 *string                              `pulumi:"extensionId"`
	Infrastructure              *string                              `pulumi:"infrastructure"`
	K8sRaw                      interface{}                          `pulumi:"k8sRaw"`
	LastUploadedDate            *string                              `pulumi:"lastUploadedDate"`
	LogAnalyticsWorkspaceConfig *LogAnalyticsWorkspaceConfigResponse `pulumi:"logAnalyticsWorkspaceConfig"`
	LogsDashboardCredential     *BasicLoginInformationResponse       `pulumi:"logsDashboardCredential"`
	MetricsDashboardCredential  *BasicLoginInformationResponse       `pulumi:"metricsDashboardCredential"`
	OnPremiseProperty           *OnPremisePropertyResponse           `pulumi:"onPremiseProperty"`
	ProvisioningState           string                               `pulumi:"provisioningState"`
	UploadServicePrincipal      *UploadServicePrincipalResponse      `pulumi:"uploadServicePrincipal"`
	UploadWatermark             *UploadWatermarkResponse             `pulumi:"uploadWatermark"`
}


func (val *DataControllerPropertiesResponse) Defaults() *DataControllerPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Infrastructure) {
		infrastructure_ := "other"
		tmp.Infrastructure = &infrastructure_
	}
	return &tmp
}

type DataControllerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DataControllerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataControllerPropertiesResponse)(nil)).Elem()
}

func (o DataControllerPropertiesResponseOutput) ToDataControllerPropertiesResponseOutput() DataControllerPropertiesResponseOutput {
	return o
}

func (o DataControllerPropertiesResponseOutput) ToDataControllerPropertiesResponseOutputWithContext(ctx context.Context) DataControllerPropertiesResponseOutput {
	return o
}

func (o DataControllerPropertiesResponseOutput) BasicLoginInformation() BasicLoginInformationResponsePtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *BasicLoginInformationResponse {
		return v.BasicLoginInformation
	}).(BasicLoginInformationResponsePtrOutput)
}

func (o DataControllerPropertiesResponseOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o DataControllerPropertiesResponseOutput) ExtensionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *string { return v.ExtensionId }).(pulumi.StringPtrOutput)
}

func (o DataControllerPropertiesResponseOutput) Infrastructure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *string { return v.Infrastructure }).(pulumi.StringPtrOutput)
}

func (o DataControllerPropertiesResponseOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) interface{} { return v.K8sRaw }).(pulumi.AnyOutput)
}

func (o DataControllerPropertiesResponseOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *string { return v.LastUploadedDate }).(pulumi.StringPtrOutput)
}

func (o DataControllerPropertiesResponseOutput) LogAnalyticsWorkspaceConfig() LogAnalyticsWorkspaceConfigResponsePtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *LogAnalyticsWorkspaceConfigResponse {
		return v.LogAnalyticsWorkspaceConfig
	}).(LogAnalyticsWorkspaceConfigResponsePtrOutput)
}

func (o DataControllerPropertiesResponseOutput) LogsDashboardCredential() BasicLoginInformationResponsePtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *BasicLoginInformationResponse {
		return v.LogsDashboardCredential
	}).(BasicLoginInformationResponsePtrOutput)
}

func (o DataControllerPropertiesResponseOutput) MetricsDashboardCredential() BasicLoginInformationResponsePtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *BasicLoginInformationResponse {
		return v.MetricsDashboardCredential
	}).(BasicLoginInformationResponsePtrOutput)
}

func (o DataControllerPropertiesResponseOutput) OnPremiseProperty() OnPremisePropertyResponsePtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *OnPremisePropertyResponse { return v.OnPremiseProperty }).(OnPremisePropertyResponsePtrOutput)
}

func (o DataControllerPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DataControllerPropertiesResponseOutput) UploadServicePrincipal() UploadServicePrincipalResponsePtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *UploadServicePrincipalResponse {
		return v.UploadServicePrincipal
	}).(UploadServicePrincipalResponsePtrOutput)
}

func (o DataControllerPropertiesResponseOutput) UploadWatermark() UploadWatermarkResponsePtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *UploadWatermarkResponse { return v.UploadWatermark }).(UploadWatermarkResponsePtrOutput)
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

type K8sResourceRequirements struct {
	Limits   map[string]string `pulumi:"limits"`
	Requests map[string]string `pulumi:"requests"`
}





type K8sResourceRequirementsInput interface {
	pulumi.Input

	ToK8sResourceRequirementsOutput() K8sResourceRequirementsOutput
	ToK8sResourceRequirementsOutputWithContext(context.Context) K8sResourceRequirementsOutput
}

type K8sResourceRequirementsArgs struct {
	Limits   pulumi.StringMapInput `pulumi:"limits"`
	Requests pulumi.StringMapInput `pulumi:"requests"`
}

func (K8sResourceRequirementsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*K8sResourceRequirements)(nil)).Elem()
}

func (i K8sResourceRequirementsArgs) ToK8sResourceRequirementsOutput() K8sResourceRequirementsOutput {
	return i.ToK8sResourceRequirementsOutputWithContext(context.Background())
}

func (i K8sResourceRequirementsArgs) ToK8sResourceRequirementsOutputWithContext(ctx context.Context) K8sResourceRequirementsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8sResourceRequirementsOutput)
}

func (i K8sResourceRequirementsArgs) ToK8sResourceRequirementsPtrOutput() K8sResourceRequirementsPtrOutput {
	return i.ToK8sResourceRequirementsPtrOutputWithContext(context.Background())
}

func (i K8sResourceRequirementsArgs) ToK8sResourceRequirementsPtrOutputWithContext(ctx context.Context) K8sResourceRequirementsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8sResourceRequirementsOutput).ToK8sResourceRequirementsPtrOutputWithContext(ctx)
}









type K8sResourceRequirementsPtrInput interface {
	pulumi.Input

	ToK8sResourceRequirementsPtrOutput() K8sResourceRequirementsPtrOutput
	ToK8sResourceRequirementsPtrOutputWithContext(context.Context) K8sResourceRequirementsPtrOutput
}

type k8sResourceRequirementsPtrType K8sResourceRequirementsArgs

func K8sResourceRequirementsPtr(v *K8sResourceRequirementsArgs) K8sResourceRequirementsPtrInput {
	return (*k8sResourceRequirementsPtrType)(v)
}

func (*k8sResourceRequirementsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**K8sResourceRequirements)(nil)).Elem()
}

func (i *k8sResourceRequirementsPtrType) ToK8sResourceRequirementsPtrOutput() K8sResourceRequirementsPtrOutput {
	return i.ToK8sResourceRequirementsPtrOutputWithContext(context.Background())
}

func (i *k8sResourceRequirementsPtrType) ToK8sResourceRequirementsPtrOutputWithContext(ctx context.Context) K8sResourceRequirementsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8sResourceRequirementsPtrOutput)
}

type K8sResourceRequirementsOutput struct{ *pulumi.OutputState }

func (K8sResourceRequirementsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*K8sResourceRequirements)(nil)).Elem()
}

func (o K8sResourceRequirementsOutput) ToK8sResourceRequirementsOutput() K8sResourceRequirementsOutput {
	return o
}

func (o K8sResourceRequirementsOutput) ToK8sResourceRequirementsOutputWithContext(ctx context.Context) K8sResourceRequirementsOutput {
	return o
}

func (o K8sResourceRequirementsOutput) ToK8sResourceRequirementsPtrOutput() K8sResourceRequirementsPtrOutput {
	return o.ToK8sResourceRequirementsPtrOutputWithContext(context.Background())
}

func (o K8sResourceRequirementsOutput) ToK8sResourceRequirementsPtrOutputWithContext(ctx context.Context) K8sResourceRequirementsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v K8sResourceRequirements) *K8sResourceRequirements {
		return &v
	}).(K8sResourceRequirementsPtrOutput)
}

func (o K8sResourceRequirementsOutput) Limits() pulumi.StringMapOutput {
	return o.ApplyT(func(v K8sResourceRequirements) map[string]string { return v.Limits }).(pulumi.StringMapOutput)
}

func (o K8sResourceRequirementsOutput) Requests() pulumi.StringMapOutput {
	return o.ApplyT(func(v K8sResourceRequirements) map[string]string { return v.Requests }).(pulumi.StringMapOutput)
}

type K8sResourceRequirementsPtrOutput struct{ *pulumi.OutputState }

func (K8sResourceRequirementsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**K8sResourceRequirements)(nil)).Elem()
}

func (o K8sResourceRequirementsPtrOutput) ToK8sResourceRequirementsPtrOutput() K8sResourceRequirementsPtrOutput {
	return o
}

func (o K8sResourceRequirementsPtrOutput) ToK8sResourceRequirementsPtrOutputWithContext(ctx context.Context) K8sResourceRequirementsPtrOutput {
	return o
}

func (o K8sResourceRequirementsPtrOutput) Elem() K8sResourceRequirementsOutput {
	return o.ApplyT(func(v *K8sResourceRequirements) K8sResourceRequirements {
		if v != nil {
			return *v
		}
		var ret K8sResourceRequirements
		return ret
	}).(K8sResourceRequirementsOutput)
}

func (o K8sResourceRequirementsPtrOutput) Limits() pulumi.StringMapOutput {
	return o.ApplyT(func(v *K8sResourceRequirements) map[string]string {
		if v == nil {
			return nil
		}
		return v.Limits
	}).(pulumi.StringMapOutput)
}

func (o K8sResourceRequirementsPtrOutput) Requests() pulumi.StringMapOutput {
	return o.ApplyT(func(v *K8sResourceRequirements) map[string]string {
		if v == nil {
			return nil
		}
		return v.Requests
	}).(pulumi.StringMapOutput)
}

type K8sResourceRequirementsResponse struct {
	Limits   map[string]string `pulumi:"limits"`
	Requests map[string]string `pulumi:"requests"`
}

type K8sResourceRequirementsResponseOutput struct{ *pulumi.OutputState }

func (K8sResourceRequirementsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*K8sResourceRequirementsResponse)(nil)).Elem()
}

func (o K8sResourceRequirementsResponseOutput) ToK8sResourceRequirementsResponseOutput() K8sResourceRequirementsResponseOutput {
	return o
}

func (o K8sResourceRequirementsResponseOutput) ToK8sResourceRequirementsResponseOutputWithContext(ctx context.Context) K8sResourceRequirementsResponseOutput {
	return o
}

func (o K8sResourceRequirementsResponseOutput) Limits() pulumi.StringMapOutput {
	return o.ApplyT(func(v K8sResourceRequirementsResponse) map[string]string { return v.Limits }).(pulumi.StringMapOutput)
}

func (o K8sResourceRequirementsResponseOutput) Requests() pulumi.StringMapOutput {
	return o.ApplyT(func(v K8sResourceRequirementsResponse) map[string]string { return v.Requests }).(pulumi.StringMapOutput)
}

type K8sResourceRequirementsResponsePtrOutput struct{ *pulumi.OutputState }

func (K8sResourceRequirementsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**K8sResourceRequirementsResponse)(nil)).Elem()
}

func (o K8sResourceRequirementsResponsePtrOutput) ToK8sResourceRequirementsResponsePtrOutput() K8sResourceRequirementsResponsePtrOutput {
	return o
}

func (o K8sResourceRequirementsResponsePtrOutput) ToK8sResourceRequirementsResponsePtrOutputWithContext(ctx context.Context) K8sResourceRequirementsResponsePtrOutput {
	return o
}

func (o K8sResourceRequirementsResponsePtrOutput) Elem() K8sResourceRequirementsResponseOutput {
	return o.ApplyT(func(v *K8sResourceRequirementsResponse) K8sResourceRequirementsResponse {
		if v != nil {
			return *v
		}
		var ret K8sResourceRequirementsResponse
		return ret
	}).(K8sResourceRequirementsResponseOutput)
}

func (o K8sResourceRequirementsResponsePtrOutput) Limits() pulumi.StringMapOutput {
	return o.ApplyT(func(v *K8sResourceRequirementsResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Limits
	}).(pulumi.StringMapOutput)
}

func (o K8sResourceRequirementsResponsePtrOutput) Requests() pulumi.StringMapOutput {
	return o.ApplyT(func(v *K8sResourceRequirementsResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Requests
	}).(pulumi.StringMapOutput)
}

type K8sScheduling struct {
	Default *K8sSchedulingOptions `pulumi:"default"`
}





type K8sSchedulingInput interface {
	pulumi.Input

	ToK8sSchedulingOutput() K8sSchedulingOutput
	ToK8sSchedulingOutputWithContext(context.Context) K8sSchedulingOutput
}

type K8sSchedulingArgs struct {
	Default K8sSchedulingOptionsPtrInput `pulumi:"default"`
}

func (K8sSchedulingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*K8sScheduling)(nil)).Elem()
}

func (i K8sSchedulingArgs) ToK8sSchedulingOutput() K8sSchedulingOutput {
	return i.ToK8sSchedulingOutputWithContext(context.Background())
}

func (i K8sSchedulingArgs) ToK8sSchedulingOutputWithContext(ctx context.Context) K8sSchedulingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8sSchedulingOutput)
}

func (i K8sSchedulingArgs) ToK8sSchedulingPtrOutput() K8sSchedulingPtrOutput {
	return i.ToK8sSchedulingPtrOutputWithContext(context.Background())
}

func (i K8sSchedulingArgs) ToK8sSchedulingPtrOutputWithContext(ctx context.Context) K8sSchedulingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8sSchedulingOutput).ToK8sSchedulingPtrOutputWithContext(ctx)
}









type K8sSchedulingPtrInput interface {
	pulumi.Input

	ToK8sSchedulingPtrOutput() K8sSchedulingPtrOutput
	ToK8sSchedulingPtrOutputWithContext(context.Context) K8sSchedulingPtrOutput
}

type k8sSchedulingPtrType K8sSchedulingArgs

func K8sSchedulingPtr(v *K8sSchedulingArgs) K8sSchedulingPtrInput {
	return (*k8sSchedulingPtrType)(v)
}

func (*k8sSchedulingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**K8sScheduling)(nil)).Elem()
}

func (i *k8sSchedulingPtrType) ToK8sSchedulingPtrOutput() K8sSchedulingPtrOutput {
	return i.ToK8sSchedulingPtrOutputWithContext(context.Background())
}

func (i *k8sSchedulingPtrType) ToK8sSchedulingPtrOutputWithContext(ctx context.Context) K8sSchedulingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8sSchedulingPtrOutput)
}

type K8sSchedulingOutput struct{ *pulumi.OutputState }

func (K8sSchedulingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*K8sScheduling)(nil)).Elem()
}

func (o K8sSchedulingOutput) ToK8sSchedulingOutput() K8sSchedulingOutput {
	return o
}

func (o K8sSchedulingOutput) ToK8sSchedulingOutputWithContext(ctx context.Context) K8sSchedulingOutput {
	return o
}

func (o K8sSchedulingOutput) ToK8sSchedulingPtrOutput() K8sSchedulingPtrOutput {
	return o.ToK8sSchedulingPtrOutputWithContext(context.Background())
}

func (o K8sSchedulingOutput) ToK8sSchedulingPtrOutputWithContext(ctx context.Context) K8sSchedulingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v K8sScheduling) *K8sScheduling {
		return &v
	}).(K8sSchedulingPtrOutput)
}

func (o K8sSchedulingOutput) Default() K8sSchedulingOptionsPtrOutput {
	return o.ApplyT(func(v K8sScheduling) *K8sSchedulingOptions { return v.Default }).(K8sSchedulingOptionsPtrOutput)
}

type K8sSchedulingPtrOutput struct{ *pulumi.OutputState }

func (K8sSchedulingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**K8sScheduling)(nil)).Elem()
}

func (o K8sSchedulingPtrOutput) ToK8sSchedulingPtrOutput() K8sSchedulingPtrOutput {
	return o
}

func (o K8sSchedulingPtrOutput) ToK8sSchedulingPtrOutputWithContext(ctx context.Context) K8sSchedulingPtrOutput {
	return o
}

func (o K8sSchedulingPtrOutput) Elem() K8sSchedulingOutput {
	return o.ApplyT(func(v *K8sScheduling) K8sScheduling {
		if v != nil {
			return *v
		}
		var ret K8sScheduling
		return ret
	}).(K8sSchedulingOutput)
}

func (o K8sSchedulingPtrOutput) Default() K8sSchedulingOptionsPtrOutput {
	return o.ApplyT(func(v *K8sScheduling) *K8sSchedulingOptions {
		if v == nil {
			return nil
		}
		return v.Default
	}).(K8sSchedulingOptionsPtrOutput)
}

type K8sSchedulingOptions struct {
	Resources *K8sResourceRequirements `pulumi:"resources"`
}





type K8sSchedulingOptionsInput interface {
	pulumi.Input

	ToK8sSchedulingOptionsOutput() K8sSchedulingOptionsOutput
	ToK8sSchedulingOptionsOutputWithContext(context.Context) K8sSchedulingOptionsOutput
}

type K8sSchedulingOptionsArgs struct {
	Resources K8sResourceRequirementsPtrInput `pulumi:"resources"`
}

func (K8sSchedulingOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*K8sSchedulingOptions)(nil)).Elem()
}

func (i K8sSchedulingOptionsArgs) ToK8sSchedulingOptionsOutput() K8sSchedulingOptionsOutput {
	return i.ToK8sSchedulingOptionsOutputWithContext(context.Background())
}

func (i K8sSchedulingOptionsArgs) ToK8sSchedulingOptionsOutputWithContext(ctx context.Context) K8sSchedulingOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8sSchedulingOptionsOutput)
}

func (i K8sSchedulingOptionsArgs) ToK8sSchedulingOptionsPtrOutput() K8sSchedulingOptionsPtrOutput {
	return i.ToK8sSchedulingOptionsPtrOutputWithContext(context.Background())
}

func (i K8sSchedulingOptionsArgs) ToK8sSchedulingOptionsPtrOutputWithContext(ctx context.Context) K8sSchedulingOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8sSchedulingOptionsOutput).ToK8sSchedulingOptionsPtrOutputWithContext(ctx)
}









type K8sSchedulingOptionsPtrInput interface {
	pulumi.Input

	ToK8sSchedulingOptionsPtrOutput() K8sSchedulingOptionsPtrOutput
	ToK8sSchedulingOptionsPtrOutputWithContext(context.Context) K8sSchedulingOptionsPtrOutput
}

type k8sSchedulingOptionsPtrType K8sSchedulingOptionsArgs

func K8sSchedulingOptionsPtr(v *K8sSchedulingOptionsArgs) K8sSchedulingOptionsPtrInput {
	return (*k8sSchedulingOptionsPtrType)(v)
}

func (*k8sSchedulingOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**K8sSchedulingOptions)(nil)).Elem()
}

func (i *k8sSchedulingOptionsPtrType) ToK8sSchedulingOptionsPtrOutput() K8sSchedulingOptionsPtrOutput {
	return i.ToK8sSchedulingOptionsPtrOutputWithContext(context.Background())
}

func (i *k8sSchedulingOptionsPtrType) ToK8sSchedulingOptionsPtrOutputWithContext(ctx context.Context) K8sSchedulingOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8sSchedulingOptionsPtrOutput)
}

type K8sSchedulingOptionsOutput struct{ *pulumi.OutputState }

func (K8sSchedulingOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*K8sSchedulingOptions)(nil)).Elem()
}

func (o K8sSchedulingOptionsOutput) ToK8sSchedulingOptionsOutput() K8sSchedulingOptionsOutput {
	return o
}

func (o K8sSchedulingOptionsOutput) ToK8sSchedulingOptionsOutputWithContext(ctx context.Context) K8sSchedulingOptionsOutput {
	return o
}

func (o K8sSchedulingOptionsOutput) ToK8sSchedulingOptionsPtrOutput() K8sSchedulingOptionsPtrOutput {
	return o.ToK8sSchedulingOptionsPtrOutputWithContext(context.Background())
}

func (o K8sSchedulingOptionsOutput) ToK8sSchedulingOptionsPtrOutputWithContext(ctx context.Context) K8sSchedulingOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v K8sSchedulingOptions) *K8sSchedulingOptions {
		return &v
	}).(K8sSchedulingOptionsPtrOutput)
}

func (o K8sSchedulingOptionsOutput) Resources() K8sResourceRequirementsPtrOutput {
	return o.ApplyT(func(v K8sSchedulingOptions) *K8sResourceRequirements { return v.Resources }).(K8sResourceRequirementsPtrOutput)
}

type K8sSchedulingOptionsPtrOutput struct{ *pulumi.OutputState }

func (K8sSchedulingOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**K8sSchedulingOptions)(nil)).Elem()
}

func (o K8sSchedulingOptionsPtrOutput) ToK8sSchedulingOptionsPtrOutput() K8sSchedulingOptionsPtrOutput {
	return o
}

func (o K8sSchedulingOptionsPtrOutput) ToK8sSchedulingOptionsPtrOutputWithContext(ctx context.Context) K8sSchedulingOptionsPtrOutput {
	return o
}

func (o K8sSchedulingOptionsPtrOutput) Elem() K8sSchedulingOptionsOutput {
	return o.ApplyT(func(v *K8sSchedulingOptions) K8sSchedulingOptions {
		if v != nil {
			return *v
		}
		var ret K8sSchedulingOptions
		return ret
	}).(K8sSchedulingOptionsOutput)
}

func (o K8sSchedulingOptionsPtrOutput) Resources() K8sResourceRequirementsPtrOutput {
	return o.ApplyT(func(v *K8sSchedulingOptions) *K8sResourceRequirements {
		if v == nil {
			return nil
		}
		return v.Resources
	}).(K8sResourceRequirementsPtrOutput)
}

type K8sSchedulingOptionsResponse struct {
	Resources *K8sResourceRequirementsResponse `pulumi:"resources"`
}

type K8sSchedulingOptionsResponseOutput struct{ *pulumi.OutputState }

func (K8sSchedulingOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*K8sSchedulingOptionsResponse)(nil)).Elem()
}

func (o K8sSchedulingOptionsResponseOutput) ToK8sSchedulingOptionsResponseOutput() K8sSchedulingOptionsResponseOutput {
	return o
}

func (o K8sSchedulingOptionsResponseOutput) ToK8sSchedulingOptionsResponseOutputWithContext(ctx context.Context) K8sSchedulingOptionsResponseOutput {
	return o
}

func (o K8sSchedulingOptionsResponseOutput) Resources() K8sResourceRequirementsResponsePtrOutput {
	return o.ApplyT(func(v K8sSchedulingOptionsResponse) *K8sResourceRequirementsResponse { return v.Resources }).(K8sResourceRequirementsResponsePtrOutput)
}

type K8sSchedulingOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (K8sSchedulingOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**K8sSchedulingOptionsResponse)(nil)).Elem()
}

func (o K8sSchedulingOptionsResponsePtrOutput) ToK8sSchedulingOptionsResponsePtrOutput() K8sSchedulingOptionsResponsePtrOutput {
	return o
}

func (o K8sSchedulingOptionsResponsePtrOutput) ToK8sSchedulingOptionsResponsePtrOutputWithContext(ctx context.Context) K8sSchedulingOptionsResponsePtrOutput {
	return o
}

func (o K8sSchedulingOptionsResponsePtrOutput) Elem() K8sSchedulingOptionsResponseOutput {
	return o.ApplyT(func(v *K8sSchedulingOptionsResponse) K8sSchedulingOptionsResponse {
		if v != nil {
			return *v
		}
		var ret K8sSchedulingOptionsResponse
		return ret
	}).(K8sSchedulingOptionsResponseOutput)
}

func (o K8sSchedulingOptionsResponsePtrOutput) Resources() K8sResourceRequirementsResponsePtrOutput {
	return o.ApplyT(func(v *K8sSchedulingOptionsResponse) *K8sResourceRequirementsResponse {
		if v == nil {
			return nil
		}
		return v.Resources
	}).(K8sResourceRequirementsResponsePtrOutput)
}

type K8sSchedulingResponse struct {
	Default *K8sSchedulingOptionsResponse `pulumi:"default"`
}

type K8sSchedulingResponseOutput struct{ *pulumi.OutputState }

func (K8sSchedulingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*K8sSchedulingResponse)(nil)).Elem()
}

func (o K8sSchedulingResponseOutput) ToK8sSchedulingResponseOutput() K8sSchedulingResponseOutput {
	return o
}

func (o K8sSchedulingResponseOutput) ToK8sSchedulingResponseOutputWithContext(ctx context.Context) K8sSchedulingResponseOutput {
	return o
}

func (o K8sSchedulingResponseOutput) Default() K8sSchedulingOptionsResponsePtrOutput {
	return o.ApplyT(func(v K8sSchedulingResponse) *K8sSchedulingOptionsResponse { return v.Default }).(K8sSchedulingOptionsResponsePtrOutput)
}

type K8sSchedulingResponsePtrOutput struct{ *pulumi.OutputState }

func (K8sSchedulingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**K8sSchedulingResponse)(nil)).Elem()
}

func (o K8sSchedulingResponsePtrOutput) ToK8sSchedulingResponsePtrOutput() K8sSchedulingResponsePtrOutput {
	return o
}

func (o K8sSchedulingResponsePtrOutput) ToK8sSchedulingResponsePtrOutputWithContext(ctx context.Context) K8sSchedulingResponsePtrOutput {
	return o
}

func (o K8sSchedulingResponsePtrOutput) Elem() K8sSchedulingResponseOutput {
	return o.ApplyT(func(v *K8sSchedulingResponse) K8sSchedulingResponse {
		if v != nil {
			return *v
		}
		var ret K8sSchedulingResponse
		return ret
	}).(K8sSchedulingResponseOutput)
}

func (o K8sSchedulingResponsePtrOutput) Default() K8sSchedulingOptionsResponsePtrOutput {
	return o.ApplyT(func(v *K8sSchedulingResponse) *K8sSchedulingOptionsResponse {
		if v == nil {
			return nil
		}
		return v.Default
	}).(K8sSchedulingOptionsResponsePtrOutput)
}

type KeytabInformation struct {
	Keytab *string `pulumi:"keytab"`
}





type KeytabInformationInput interface {
	pulumi.Input

	ToKeytabInformationOutput() KeytabInformationOutput
	ToKeytabInformationOutputWithContext(context.Context) KeytabInformationOutput
}

type KeytabInformationArgs struct {
	Keytab pulumi.StringPtrInput `pulumi:"keytab"`
}

func (KeytabInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeytabInformation)(nil)).Elem()
}

func (i KeytabInformationArgs) ToKeytabInformationOutput() KeytabInformationOutput {
	return i.ToKeytabInformationOutputWithContext(context.Background())
}

func (i KeytabInformationArgs) ToKeytabInformationOutputWithContext(ctx context.Context) KeytabInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeytabInformationOutput)
}

func (i KeytabInformationArgs) ToKeytabInformationPtrOutput() KeytabInformationPtrOutput {
	return i.ToKeytabInformationPtrOutputWithContext(context.Background())
}

func (i KeytabInformationArgs) ToKeytabInformationPtrOutputWithContext(ctx context.Context) KeytabInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeytabInformationOutput).ToKeytabInformationPtrOutputWithContext(ctx)
}









type KeytabInformationPtrInput interface {
	pulumi.Input

	ToKeytabInformationPtrOutput() KeytabInformationPtrOutput
	ToKeytabInformationPtrOutputWithContext(context.Context) KeytabInformationPtrOutput
}

type keytabInformationPtrType KeytabInformationArgs

func KeytabInformationPtr(v *KeytabInformationArgs) KeytabInformationPtrInput {
	return (*keytabInformationPtrType)(v)
}

func (*keytabInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeytabInformation)(nil)).Elem()
}

func (i *keytabInformationPtrType) ToKeytabInformationPtrOutput() KeytabInformationPtrOutput {
	return i.ToKeytabInformationPtrOutputWithContext(context.Background())
}

func (i *keytabInformationPtrType) ToKeytabInformationPtrOutputWithContext(ctx context.Context) KeytabInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeytabInformationPtrOutput)
}

type KeytabInformationOutput struct{ *pulumi.OutputState }

func (KeytabInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeytabInformation)(nil)).Elem()
}

func (o KeytabInformationOutput) ToKeytabInformationOutput() KeytabInformationOutput {
	return o
}

func (o KeytabInformationOutput) ToKeytabInformationOutputWithContext(ctx context.Context) KeytabInformationOutput {
	return o
}

func (o KeytabInformationOutput) ToKeytabInformationPtrOutput() KeytabInformationPtrOutput {
	return o.ToKeytabInformationPtrOutputWithContext(context.Background())
}

func (o KeytabInformationOutput) ToKeytabInformationPtrOutputWithContext(ctx context.Context) KeytabInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeytabInformation) *KeytabInformation {
		return &v
	}).(KeytabInformationPtrOutput)
}

func (o KeytabInformationOutput) Keytab() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeytabInformation) *string { return v.Keytab }).(pulumi.StringPtrOutput)
}

type KeytabInformationPtrOutput struct{ *pulumi.OutputState }

func (KeytabInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeytabInformation)(nil)).Elem()
}

func (o KeytabInformationPtrOutput) ToKeytabInformationPtrOutput() KeytabInformationPtrOutput {
	return o
}

func (o KeytabInformationPtrOutput) ToKeytabInformationPtrOutputWithContext(ctx context.Context) KeytabInformationPtrOutput {
	return o
}

func (o KeytabInformationPtrOutput) Elem() KeytabInformationOutput {
	return o.ApplyT(func(v *KeytabInformation) KeytabInformation {
		if v != nil {
			return *v
		}
		var ret KeytabInformation
		return ret
	}).(KeytabInformationOutput)
}

func (o KeytabInformationPtrOutput) Keytab() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeytabInformation) *string {
		if v == nil {
			return nil
		}
		return v.Keytab
	}).(pulumi.StringPtrOutput)
}

type LogAnalyticsWorkspaceConfig struct {
	PrimaryKey  *string `pulumi:"primaryKey"`
	WorkspaceId *string `pulumi:"workspaceId"`
}





type LogAnalyticsWorkspaceConfigInput interface {
	pulumi.Input

	ToLogAnalyticsWorkspaceConfigOutput() LogAnalyticsWorkspaceConfigOutput
	ToLogAnalyticsWorkspaceConfigOutputWithContext(context.Context) LogAnalyticsWorkspaceConfigOutput
}

type LogAnalyticsWorkspaceConfigArgs struct {
	PrimaryKey  pulumi.StringPtrInput `pulumi:"primaryKey"`
	WorkspaceId pulumi.StringPtrInput `pulumi:"workspaceId"`
}

func (LogAnalyticsWorkspaceConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsWorkspaceConfig)(nil)).Elem()
}

func (i LogAnalyticsWorkspaceConfigArgs) ToLogAnalyticsWorkspaceConfigOutput() LogAnalyticsWorkspaceConfigOutput {
	return i.ToLogAnalyticsWorkspaceConfigOutputWithContext(context.Background())
}

func (i LogAnalyticsWorkspaceConfigArgs) ToLogAnalyticsWorkspaceConfigOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsWorkspaceConfigOutput)
}

func (i LogAnalyticsWorkspaceConfigArgs) ToLogAnalyticsWorkspaceConfigPtrOutput() LogAnalyticsWorkspaceConfigPtrOutput {
	return i.ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(context.Background())
}

func (i LogAnalyticsWorkspaceConfigArgs) ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsWorkspaceConfigOutput).ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(ctx)
}









type LogAnalyticsWorkspaceConfigPtrInput interface {
	pulumi.Input

	ToLogAnalyticsWorkspaceConfigPtrOutput() LogAnalyticsWorkspaceConfigPtrOutput
	ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(context.Context) LogAnalyticsWorkspaceConfigPtrOutput
}

type logAnalyticsWorkspaceConfigPtrType LogAnalyticsWorkspaceConfigArgs

func LogAnalyticsWorkspaceConfigPtr(v *LogAnalyticsWorkspaceConfigArgs) LogAnalyticsWorkspaceConfigPtrInput {
	return (*logAnalyticsWorkspaceConfigPtrType)(v)
}

func (*logAnalyticsWorkspaceConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsWorkspaceConfig)(nil)).Elem()
}

func (i *logAnalyticsWorkspaceConfigPtrType) ToLogAnalyticsWorkspaceConfigPtrOutput() LogAnalyticsWorkspaceConfigPtrOutput {
	return i.ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(context.Background())
}

func (i *logAnalyticsWorkspaceConfigPtrType) ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsWorkspaceConfigPtrOutput)
}

type LogAnalyticsWorkspaceConfigOutput struct{ *pulumi.OutputState }

func (LogAnalyticsWorkspaceConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsWorkspaceConfig)(nil)).Elem()
}

func (o LogAnalyticsWorkspaceConfigOutput) ToLogAnalyticsWorkspaceConfigOutput() LogAnalyticsWorkspaceConfigOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigOutput) ToLogAnalyticsWorkspaceConfigOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigOutput) ToLogAnalyticsWorkspaceConfigPtrOutput() LogAnalyticsWorkspaceConfigPtrOutput {
	return o.ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(context.Background())
}

func (o LogAnalyticsWorkspaceConfigOutput) ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogAnalyticsWorkspaceConfig) *LogAnalyticsWorkspaceConfig {
		return &v
	}).(LogAnalyticsWorkspaceConfigPtrOutput)
}

func (o LogAnalyticsWorkspaceConfigOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsWorkspaceConfig) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsWorkspaceConfigOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsWorkspaceConfig) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

type LogAnalyticsWorkspaceConfigPtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsWorkspaceConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsWorkspaceConfig)(nil)).Elem()
}

func (o LogAnalyticsWorkspaceConfigPtrOutput) ToLogAnalyticsWorkspaceConfigPtrOutput() LogAnalyticsWorkspaceConfigPtrOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigPtrOutput) ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigPtrOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigPtrOutput) Elem() LogAnalyticsWorkspaceConfigOutput {
	return o.ApplyT(func(v *LogAnalyticsWorkspaceConfig) LogAnalyticsWorkspaceConfig {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsWorkspaceConfig
		return ret
	}).(LogAnalyticsWorkspaceConfigOutput)
}

func (o LogAnalyticsWorkspaceConfigPtrOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsWorkspaceConfig) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryKey
	}).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsWorkspaceConfigPtrOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsWorkspaceConfig) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceId
	}).(pulumi.StringPtrOutput)
}

type LogAnalyticsWorkspaceConfigResponse struct {
	WorkspaceId *string `pulumi:"workspaceId"`
}

type LogAnalyticsWorkspaceConfigResponseOutput struct{ *pulumi.OutputState }

func (LogAnalyticsWorkspaceConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsWorkspaceConfigResponse)(nil)).Elem()
}

func (o LogAnalyticsWorkspaceConfigResponseOutput) ToLogAnalyticsWorkspaceConfigResponseOutput() LogAnalyticsWorkspaceConfigResponseOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigResponseOutput) ToLogAnalyticsWorkspaceConfigResponseOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigResponseOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigResponseOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsWorkspaceConfigResponse) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

type LogAnalyticsWorkspaceConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsWorkspaceConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsWorkspaceConfigResponse)(nil)).Elem()
}

func (o LogAnalyticsWorkspaceConfigResponsePtrOutput) ToLogAnalyticsWorkspaceConfigResponsePtrOutput() LogAnalyticsWorkspaceConfigResponsePtrOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigResponsePtrOutput) ToLogAnalyticsWorkspaceConfigResponsePtrOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigResponsePtrOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigResponsePtrOutput) Elem() LogAnalyticsWorkspaceConfigResponseOutput {
	return o.ApplyT(func(v *LogAnalyticsWorkspaceConfigResponse) LogAnalyticsWorkspaceConfigResponse {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsWorkspaceConfigResponse
		return ret
	}).(LogAnalyticsWorkspaceConfigResponseOutput)
}

func (o LogAnalyticsWorkspaceConfigResponsePtrOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsWorkspaceConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceId
	}).(pulumi.StringPtrOutput)
}

type OnPremiseProperty struct {
	Id                           string  `pulumi:"id"`
	PublicSigningKey             string  `pulumi:"publicSigningKey"`
	SigningCertificateThumbprint *string `pulumi:"signingCertificateThumbprint"`
}





type OnPremisePropertyInput interface {
	pulumi.Input

	ToOnPremisePropertyOutput() OnPremisePropertyOutput
	ToOnPremisePropertyOutputWithContext(context.Context) OnPremisePropertyOutput
}

type OnPremisePropertyArgs struct {
	Id                           pulumi.StringInput    `pulumi:"id"`
	PublicSigningKey             pulumi.StringInput    `pulumi:"publicSigningKey"`
	SigningCertificateThumbprint pulumi.StringPtrInput `pulumi:"signingCertificateThumbprint"`
}

func (OnPremisePropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseProperty)(nil)).Elem()
}

func (i OnPremisePropertyArgs) ToOnPremisePropertyOutput() OnPremisePropertyOutput {
	return i.ToOnPremisePropertyOutputWithContext(context.Background())
}

func (i OnPremisePropertyArgs) ToOnPremisePropertyOutputWithContext(ctx context.Context) OnPremisePropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremisePropertyOutput)
}

func (i OnPremisePropertyArgs) ToOnPremisePropertyPtrOutput() OnPremisePropertyPtrOutput {
	return i.ToOnPremisePropertyPtrOutputWithContext(context.Background())
}

func (i OnPremisePropertyArgs) ToOnPremisePropertyPtrOutputWithContext(ctx context.Context) OnPremisePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremisePropertyOutput).ToOnPremisePropertyPtrOutputWithContext(ctx)
}









type OnPremisePropertyPtrInput interface {
	pulumi.Input

	ToOnPremisePropertyPtrOutput() OnPremisePropertyPtrOutput
	ToOnPremisePropertyPtrOutputWithContext(context.Context) OnPremisePropertyPtrOutput
}

type onPremisePropertyPtrType OnPremisePropertyArgs

func OnPremisePropertyPtr(v *OnPremisePropertyArgs) OnPremisePropertyPtrInput {
	return (*onPremisePropertyPtrType)(v)
}

func (*onPremisePropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OnPremiseProperty)(nil)).Elem()
}

func (i *onPremisePropertyPtrType) ToOnPremisePropertyPtrOutput() OnPremisePropertyPtrOutput {
	return i.ToOnPremisePropertyPtrOutputWithContext(context.Background())
}

func (i *onPremisePropertyPtrType) ToOnPremisePropertyPtrOutputWithContext(ctx context.Context) OnPremisePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremisePropertyPtrOutput)
}

type OnPremisePropertyOutput struct{ *pulumi.OutputState }

func (OnPremisePropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseProperty)(nil)).Elem()
}

func (o OnPremisePropertyOutput) ToOnPremisePropertyOutput() OnPremisePropertyOutput {
	return o
}

func (o OnPremisePropertyOutput) ToOnPremisePropertyOutputWithContext(ctx context.Context) OnPremisePropertyOutput {
	return o
}

func (o OnPremisePropertyOutput) ToOnPremisePropertyPtrOutput() OnPremisePropertyPtrOutput {
	return o.ToOnPremisePropertyPtrOutputWithContext(context.Background())
}

func (o OnPremisePropertyOutput) ToOnPremisePropertyPtrOutputWithContext(ctx context.Context) OnPremisePropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OnPremiseProperty) *OnPremiseProperty {
		return &v
	}).(OnPremisePropertyPtrOutput)
}

func (o OnPremisePropertyOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseProperty) string { return v.Id }).(pulumi.StringOutput)
}

func (o OnPremisePropertyOutput) PublicSigningKey() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseProperty) string { return v.PublicSigningKey }).(pulumi.StringOutput)
}

func (o OnPremisePropertyOutput) SigningCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OnPremiseProperty) *string { return v.SigningCertificateThumbprint }).(pulumi.StringPtrOutput)
}

type OnPremisePropertyPtrOutput struct{ *pulumi.OutputState }

func (OnPremisePropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OnPremiseProperty)(nil)).Elem()
}

func (o OnPremisePropertyPtrOutput) ToOnPremisePropertyPtrOutput() OnPremisePropertyPtrOutput {
	return o
}

func (o OnPremisePropertyPtrOutput) ToOnPremisePropertyPtrOutputWithContext(ctx context.Context) OnPremisePropertyPtrOutput {
	return o
}

func (o OnPremisePropertyPtrOutput) Elem() OnPremisePropertyOutput {
	return o.ApplyT(func(v *OnPremiseProperty) OnPremiseProperty {
		if v != nil {
			return *v
		}
		var ret OnPremiseProperty
		return ret
	}).(OnPremisePropertyOutput)
}

func (o OnPremisePropertyPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnPremiseProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o OnPremisePropertyPtrOutput) PublicSigningKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnPremiseProperty) *string {
		if v == nil {
			return nil
		}
		return &v.PublicSigningKey
	}).(pulumi.StringPtrOutput)
}

func (o OnPremisePropertyPtrOutput) SigningCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnPremiseProperty) *string {
		if v == nil {
			return nil
		}
		return v.SigningCertificateThumbprint
	}).(pulumi.StringPtrOutput)
}

type OnPremisePropertyResponse struct {
	Id                           string  `pulumi:"id"`
	PublicSigningKey             string  `pulumi:"publicSigningKey"`
	SigningCertificateThumbprint *string `pulumi:"signingCertificateThumbprint"`
}

type OnPremisePropertyResponseOutput struct{ *pulumi.OutputState }

func (OnPremisePropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremisePropertyResponse)(nil)).Elem()
}

func (o OnPremisePropertyResponseOutput) ToOnPremisePropertyResponseOutput() OnPremisePropertyResponseOutput {
	return o
}

func (o OnPremisePropertyResponseOutput) ToOnPremisePropertyResponseOutputWithContext(ctx context.Context) OnPremisePropertyResponseOutput {
	return o
}

func (o OnPremisePropertyResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremisePropertyResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o OnPremisePropertyResponseOutput) PublicSigningKey() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremisePropertyResponse) string { return v.PublicSigningKey }).(pulumi.StringOutput)
}

func (o OnPremisePropertyResponseOutput) SigningCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OnPremisePropertyResponse) *string { return v.SigningCertificateThumbprint }).(pulumi.StringPtrOutput)
}

type OnPremisePropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (OnPremisePropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OnPremisePropertyResponse)(nil)).Elem()
}

func (o OnPremisePropertyResponsePtrOutput) ToOnPremisePropertyResponsePtrOutput() OnPremisePropertyResponsePtrOutput {
	return o
}

func (o OnPremisePropertyResponsePtrOutput) ToOnPremisePropertyResponsePtrOutputWithContext(ctx context.Context) OnPremisePropertyResponsePtrOutput {
	return o
}

func (o OnPremisePropertyResponsePtrOutput) Elem() OnPremisePropertyResponseOutput {
	return o.ApplyT(func(v *OnPremisePropertyResponse) OnPremisePropertyResponse {
		if v != nil {
			return *v
		}
		var ret OnPremisePropertyResponse
		return ret
	}).(OnPremisePropertyResponseOutput)
}

func (o OnPremisePropertyResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnPremisePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o OnPremisePropertyResponsePtrOutput) PublicSigningKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnPremisePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PublicSigningKey
	}).(pulumi.StringPtrOutput)
}

func (o OnPremisePropertyResponsePtrOutput) SigningCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnPremisePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.SigningCertificateThumbprint
	}).(pulumi.StringPtrOutput)
}

type PostgresInstanceProperties struct {
	Admin                 *string                `pulumi:"admin"`
	BasicLoginInformation *BasicLoginInformation `pulumi:"basicLoginInformation"`
	DataControllerId      *string                `pulumi:"dataControllerId"`
	K8sRaw                interface{}            `pulumi:"k8sRaw"`
	LastUploadedDate      *string                `pulumi:"lastUploadedDate"`
}





type PostgresInstancePropertiesInput interface {
	pulumi.Input

	ToPostgresInstancePropertiesOutput() PostgresInstancePropertiesOutput
	ToPostgresInstancePropertiesOutputWithContext(context.Context) PostgresInstancePropertiesOutput
}

type PostgresInstancePropertiesArgs struct {
	Admin                 pulumi.StringPtrInput         `pulumi:"admin"`
	BasicLoginInformation BasicLoginInformationPtrInput `pulumi:"basicLoginInformation"`
	DataControllerId      pulumi.StringPtrInput         `pulumi:"dataControllerId"`
	K8sRaw                pulumi.Input                  `pulumi:"k8sRaw"`
	LastUploadedDate      pulumi.StringPtrInput         `pulumi:"lastUploadedDate"`
}

func (PostgresInstancePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PostgresInstanceProperties)(nil)).Elem()
}

func (i PostgresInstancePropertiesArgs) ToPostgresInstancePropertiesOutput() PostgresInstancePropertiesOutput {
	return i.ToPostgresInstancePropertiesOutputWithContext(context.Background())
}

func (i PostgresInstancePropertiesArgs) ToPostgresInstancePropertiesOutputWithContext(ctx context.Context) PostgresInstancePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstancePropertiesOutput)
}

type PostgresInstancePropertiesOutput struct{ *pulumi.OutputState }

func (PostgresInstancePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PostgresInstanceProperties)(nil)).Elem()
}

func (o PostgresInstancePropertiesOutput) ToPostgresInstancePropertiesOutput() PostgresInstancePropertiesOutput {
	return o
}

func (o PostgresInstancePropertiesOutput) ToPostgresInstancePropertiesOutputWithContext(ctx context.Context) PostgresInstancePropertiesOutput {
	return o
}

func (o PostgresInstancePropertiesOutput) Admin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceProperties) *string { return v.Admin }).(pulumi.StringPtrOutput)
}

func (o PostgresInstancePropertiesOutput) BasicLoginInformation() BasicLoginInformationPtrOutput {
	return o.ApplyT(func(v PostgresInstanceProperties) *BasicLoginInformation { return v.BasicLoginInformation }).(BasicLoginInformationPtrOutput)
}

func (o PostgresInstancePropertiesOutput) DataControllerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceProperties) *string { return v.DataControllerId }).(pulumi.StringPtrOutput)
}

func (o PostgresInstancePropertiesOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v PostgresInstanceProperties) interface{} { return v.K8sRaw }).(pulumi.AnyOutput)
}

func (o PostgresInstancePropertiesOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceProperties) *string { return v.LastUploadedDate }).(pulumi.StringPtrOutput)
}

type PostgresInstancePropertiesResponse struct {
	Admin                 *string                        `pulumi:"admin"`
	BasicLoginInformation *BasicLoginInformationResponse `pulumi:"basicLoginInformation"`
	DataControllerId      *string                        `pulumi:"dataControllerId"`
	K8sRaw                interface{}                    `pulumi:"k8sRaw"`
	LastUploadedDate      *string                        `pulumi:"lastUploadedDate"`
	ProvisioningState     string                         `pulumi:"provisioningState"`
}

type PostgresInstancePropertiesResponseOutput struct{ *pulumi.OutputState }

func (PostgresInstancePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PostgresInstancePropertiesResponse)(nil)).Elem()
}

func (o PostgresInstancePropertiesResponseOutput) ToPostgresInstancePropertiesResponseOutput() PostgresInstancePropertiesResponseOutput {
	return o
}

func (o PostgresInstancePropertiesResponseOutput) ToPostgresInstancePropertiesResponseOutputWithContext(ctx context.Context) PostgresInstancePropertiesResponseOutput {
	return o
}

func (o PostgresInstancePropertiesResponseOutput) Admin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstancePropertiesResponse) *string { return v.Admin }).(pulumi.StringPtrOutput)
}

func (o PostgresInstancePropertiesResponseOutput) BasicLoginInformation() BasicLoginInformationResponsePtrOutput {
	return o.ApplyT(func(v PostgresInstancePropertiesResponse) *BasicLoginInformationResponse {
		return v.BasicLoginInformation
	}).(BasicLoginInformationResponsePtrOutput)
}

func (o PostgresInstancePropertiesResponseOutput) DataControllerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstancePropertiesResponse) *string { return v.DataControllerId }).(pulumi.StringPtrOutput)
}

func (o PostgresInstancePropertiesResponseOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v PostgresInstancePropertiesResponse) interface{} { return v.K8sRaw }).(pulumi.AnyOutput)
}

func (o PostgresInstancePropertiesResponseOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstancePropertiesResponse) *string { return v.LastUploadedDate }).(pulumi.StringPtrOutput)
}

func (o PostgresInstancePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PostgresInstancePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type PostgresInstanceSku struct {
	Capacity *int                     `pulumi:"capacity"`
	Dev      *bool                    `pulumi:"dev"`
	Family   *string                  `pulumi:"family"`
	Name     string                   `pulumi:"name"`
	Size     *string                  `pulumi:"size"`
	Tier     *PostgresInstanceSkuTier `pulumi:"tier"`
}


func (val *PostgresInstanceSku) Defaults() *PostgresInstanceSku {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Dev) {
		dev_ := true
		tmp.Dev = &dev_
	}
	if isZero(tmp.Tier) {
		tier_ := PostgresInstanceSkuTier("Hyperscale")
		tmp.Tier = &tier_
	}
	return &tmp
}





type PostgresInstanceSkuInput interface {
	pulumi.Input

	ToPostgresInstanceSkuOutput() PostgresInstanceSkuOutput
	ToPostgresInstanceSkuOutputWithContext(context.Context) PostgresInstanceSkuOutput
}

type PostgresInstanceSkuArgs struct {
	Capacity pulumi.IntPtrInput              `pulumi:"capacity"`
	Dev      pulumi.BoolPtrInput             `pulumi:"dev"`
	Family   pulumi.StringPtrInput           `pulumi:"family"`
	Name     pulumi.StringInput              `pulumi:"name"`
	Size     pulumi.StringPtrInput           `pulumi:"size"`
	Tier     PostgresInstanceSkuTierPtrInput `pulumi:"tier"`
}


func (val *PostgresInstanceSkuArgs) Defaults() *PostgresInstanceSkuArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Dev) {
		tmp.Dev = pulumi.BoolPtr(true)
	}
	if isZero(tmp.Tier) {
		tmp.Tier = PostgresInstanceSkuTier("Hyperscale")
	}
	return &tmp
}
func (PostgresInstanceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PostgresInstanceSku)(nil)).Elem()
}

func (i PostgresInstanceSkuArgs) ToPostgresInstanceSkuOutput() PostgresInstanceSkuOutput {
	return i.ToPostgresInstanceSkuOutputWithContext(context.Background())
}

func (i PostgresInstanceSkuArgs) ToPostgresInstanceSkuOutputWithContext(ctx context.Context) PostgresInstanceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstanceSkuOutput)
}

func (i PostgresInstanceSkuArgs) ToPostgresInstanceSkuPtrOutput() PostgresInstanceSkuPtrOutput {
	return i.ToPostgresInstanceSkuPtrOutputWithContext(context.Background())
}

func (i PostgresInstanceSkuArgs) ToPostgresInstanceSkuPtrOutputWithContext(ctx context.Context) PostgresInstanceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstanceSkuOutput).ToPostgresInstanceSkuPtrOutputWithContext(ctx)
}









type PostgresInstanceSkuPtrInput interface {
	pulumi.Input

	ToPostgresInstanceSkuPtrOutput() PostgresInstanceSkuPtrOutput
	ToPostgresInstanceSkuPtrOutputWithContext(context.Context) PostgresInstanceSkuPtrOutput
}

type postgresInstanceSkuPtrType PostgresInstanceSkuArgs

func PostgresInstanceSkuPtr(v *PostgresInstanceSkuArgs) PostgresInstanceSkuPtrInput {
	return (*postgresInstanceSkuPtrType)(v)
}

func (*postgresInstanceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PostgresInstanceSku)(nil)).Elem()
}

func (i *postgresInstanceSkuPtrType) ToPostgresInstanceSkuPtrOutput() PostgresInstanceSkuPtrOutput {
	return i.ToPostgresInstanceSkuPtrOutputWithContext(context.Background())
}

func (i *postgresInstanceSkuPtrType) ToPostgresInstanceSkuPtrOutputWithContext(ctx context.Context) PostgresInstanceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstanceSkuPtrOutput)
}

type PostgresInstanceSkuOutput struct{ *pulumi.OutputState }

func (PostgresInstanceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PostgresInstanceSku)(nil)).Elem()
}

func (o PostgresInstanceSkuOutput) ToPostgresInstanceSkuOutput() PostgresInstanceSkuOutput {
	return o
}

func (o PostgresInstanceSkuOutput) ToPostgresInstanceSkuOutputWithContext(ctx context.Context) PostgresInstanceSkuOutput {
	return o
}

func (o PostgresInstanceSkuOutput) ToPostgresInstanceSkuPtrOutput() PostgresInstanceSkuPtrOutput {
	return o.ToPostgresInstanceSkuPtrOutputWithContext(context.Background())
}

func (o PostgresInstanceSkuOutput) ToPostgresInstanceSkuPtrOutputWithContext(ctx context.Context) PostgresInstanceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PostgresInstanceSku) *PostgresInstanceSku {
		return &v
	}).(PostgresInstanceSkuPtrOutput)
}

func (o PostgresInstanceSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o PostgresInstanceSkuOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSku) *bool { return v.Dev }).(pulumi.BoolPtrOutput)
}

func (o PostgresInstanceSkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PostgresInstanceSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o PostgresInstanceSkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuOutput) Tier() PostgresInstanceSkuTierPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSku) *PostgresInstanceSkuTier { return v.Tier }).(PostgresInstanceSkuTierPtrOutput)
}

type PostgresInstanceSkuPtrOutput struct{ *pulumi.OutputState }

func (PostgresInstanceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PostgresInstanceSku)(nil)).Elem()
}

func (o PostgresInstanceSkuPtrOutput) ToPostgresInstanceSkuPtrOutput() PostgresInstanceSkuPtrOutput {
	return o
}

func (o PostgresInstanceSkuPtrOutput) ToPostgresInstanceSkuPtrOutputWithContext(ctx context.Context) PostgresInstanceSkuPtrOutput {
	return o
}

func (o PostgresInstanceSkuPtrOutput) Elem() PostgresInstanceSkuOutput {
	return o.ApplyT(func(v *PostgresInstanceSku) PostgresInstanceSku {
		if v != nil {
			return *v
		}
		var ret PostgresInstanceSku
		return ret
	}).(PostgresInstanceSkuOutput)
}

func (o PostgresInstanceSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o PostgresInstanceSkuPtrOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSku) *bool {
		if v == nil {
			return nil
		}
		return v.Dev
	}).(pulumi.BoolPtrOutput)
}

func (o PostgresInstanceSkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuPtrOutput) Tier() PostgresInstanceSkuTierPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSku) *PostgresInstanceSkuTier {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(PostgresInstanceSkuTierPtrOutput)
}

type PostgresInstanceSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Dev      *bool   `pulumi:"dev"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}


func (val *PostgresInstanceSkuResponse) Defaults() *PostgresInstanceSkuResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Dev) {
		dev_ := true
		tmp.Dev = &dev_
	}
	if isZero(tmp.Tier) {
		tier_ := "Hyperscale"
		tmp.Tier = &tier_
	}
	return &tmp
}

type PostgresInstanceSkuResponseOutput struct{ *pulumi.OutputState }

func (PostgresInstanceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PostgresInstanceSkuResponse)(nil)).Elem()
}

func (o PostgresInstanceSkuResponseOutput) ToPostgresInstanceSkuResponseOutput() PostgresInstanceSkuResponseOutput {
	return o
}

func (o PostgresInstanceSkuResponseOutput) ToPostgresInstanceSkuResponseOutputWithContext(ctx context.Context) PostgresInstanceSkuResponseOutput {
	return o
}

func (o PostgresInstanceSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o PostgresInstanceSkuResponseOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSkuResponse) *bool { return v.Dev }).(pulumi.BoolPtrOutput)
}

func (o PostgresInstanceSkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PostgresInstanceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PostgresInstanceSkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type PostgresInstanceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (PostgresInstanceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PostgresInstanceSkuResponse)(nil)).Elem()
}

func (o PostgresInstanceSkuResponsePtrOutput) ToPostgresInstanceSkuResponsePtrOutput() PostgresInstanceSkuResponsePtrOutput {
	return o
}

func (o PostgresInstanceSkuResponsePtrOutput) ToPostgresInstanceSkuResponsePtrOutputWithContext(ctx context.Context) PostgresInstanceSkuResponsePtrOutput {
	return o
}

func (o PostgresInstanceSkuResponsePtrOutput) Elem() PostgresInstanceSkuResponseOutput {
	return o.ApplyT(func(v *PostgresInstanceSkuResponse) PostgresInstanceSkuResponse {
		if v != nil {
			return *v
		}
		var ret PostgresInstanceSkuResponse
		return ret
	}).(PostgresInstanceSkuResponseOutput)
}

func (o PostgresInstanceSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o PostgresInstanceSkuResponsePtrOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSkuResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Dev
	}).(pulumi.BoolPtrOutput)
}

func (o PostgresInstanceSkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SqlManagedInstanceK8sRaw struct {
	Spec *SqlManagedInstanceK8sSpec `pulumi:"spec"`
}





type SqlManagedInstanceK8sRawInput interface {
	pulumi.Input

	ToSqlManagedInstanceK8sRawOutput() SqlManagedInstanceK8sRawOutput
	ToSqlManagedInstanceK8sRawOutputWithContext(context.Context) SqlManagedInstanceK8sRawOutput
}

type SqlManagedInstanceK8sRawArgs struct {
	Spec SqlManagedInstanceK8sSpecPtrInput `pulumi:"spec"`
}

func (SqlManagedInstanceK8sRawArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceK8sRaw)(nil)).Elem()
}

func (i SqlManagedInstanceK8sRawArgs) ToSqlManagedInstanceK8sRawOutput() SqlManagedInstanceK8sRawOutput {
	return i.ToSqlManagedInstanceK8sRawOutputWithContext(context.Background())
}

func (i SqlManagedInstanceK8sRawArgs) ToSqlManagedInstanceK8sRawOutputWithContext(ctx context.Context) SqlManagedInstanceK8sRawOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceK8sRawOutput)
}

func (i SqlManagedInstanceK8sRawArgs) ToSqlManagedInstanceK8sRawPtrOutput() SqlManagedInstanceK8sRawPtrOutput {
	return i.ToSqlManagedInstanceK8sRawPtrOutputWithContext(context.Background())
}

func (i SqlManagedInstanceK8sRawArgs) ToSqlManagedInstanceK8sRawPtrOutputWithContext(ctx context.Context) SqlManagedInstanceK8sRawPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceK8sRawOutput).ToSqlManagedInstanceK8sRawPtrOutputWithContext(ctx)
}









type SqlManagedInstanceK8sRawPtrInput interface {
	pulumi.Input

	ToSqlManagedInstanceK8sRawPtrOutput() SqlManagedInstanceK8sRawPtrOutput
	ToSqlManagedInstanceK8sRawPtrOutputWithContext(context.Context) SqlManagedInstanceK8sRawPtrOutput
}

type sqlManagedInstanceK8sRawPtrType SqlManagedInstanceK8sRawArgs

func SqlManagedInstanceK8sRawPtr(v *SqlManagedInstanceK8sRawArgs) SqlManagedInstanceK8sRawPtrInput {
	return (*sqlManagedInstanceK8sRawPtrType)(v)
}

func (*sqlManagedInstanceK8sRawPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceK8sRaw)(nil)).Elem()
}

func (i *sqlManagedInstanceK8sRawPtrType) ToSqlManagedInstanceK8sRawPtrOutput() SqlManagedInstanceK8sRawPtrOutput {
	return i.ToSqlManagedInstanceK8sRawPtrOutputWithContext(context.Background())
}

func (i *sqlManagedInstanceK8sRawPtrType) ToSqlManagedInstanceK8sRawPtrOutputWithContext(ctx context.Context) SqlManagedInstanceK8sRawPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceK8sRawPtrOutput)
}

type SqlManagedInstanceK8sRawOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceK8sRawOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceK8sRaw)(nil)).Elem()
}

func (o SqlManagedInstanceK8sRawOutput) ToSqlManagedInstanceK8sRawOutput() SqlManagedInstanceK8sRawOutput {
	return o
}

func (o SqlManagedInstanceK8sRawOutput) ToSqlManagedInstanceK8sRawOutputWithContext(ctx context.Context) SqlManagedInstanceK8sRawOutput {
	return o
}

func (o SqlManagedInstanceK8sRawOutput) ToSqlManagedInstanceK8sRawPtrOutput() SqlManagedInstanceK8sRawPtrOutput {
	return o.ToSqlManagedInstanceK8sRawPtrOutputWithContext(context.Background())
}

func (o SqlManagedInstanceK8sRawOutput) ToSqlManagedInstanceK8sRawPtrOutputWithContext(ctx context.Context) SqlManagedInstanceK8sRawPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlManagedInstanceK8sRaw) *SqlManagedInstanceK8sRaw {
		return &v
	}).(SqlManagedInstanceK8sRawPtrOutput)
}

func (o SqlManagedInstanceK8sRawOutput) Spec() SqlManagedInstanceK8sSpecPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceK8sRaw) *SqlManagedInstanceK8sSpec { return v.Spec }).(SqlManagedInstanceK8sSpecPtrOutput)
}

type SqlManagedInstanceK8sRawPtrOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceK8sRawPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceK8sRaw)(nil)).Elem()
}

func (o SqlManagedInstanceK8sRawPtrOutput) ToSqlManagedInstanceK8sRawPtrOutput() SqlManagedInstanceK8sRawPtrOutput {
	return o
}

func (o SqlManagedInstanceK8sRawPtrOutput) ToSqlManagedInstanceK8sRawPtrOutputWithContext(ctx context.Context) SqlManagedInstanceK8sRawPtrOutput {
	return o
}

func (o SqlManagedInstanceK8sRawPtrOutput) Elem() SqlManagedInstanceK8sRawOutput {
	return o.ApplyT(func(v *SqlManagedInstanceK8sRaw) SqlManagedInstanceK8sRaw {
		if v != nil {
			return *v
		}
		var ret SqlManagedInstanceK8sRaw
		return ret
	}).(SqlManagedInstanceK8sRawOutput)
}

func (o SqlManagedInstanceK8sRawPtrOutput) Spec() SqlManagedInstanceK8sSpecPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceK8sRaw) *SqlManagedInstanceK8sSpec {
		if v == nil {
			return nil
		}
		return v.Spec
	}).(SqlManagedInstanceK8sSpecPtrOutput)
}

type SqlManagedInstanceK8sRawResponse struct {
	Spec *SqlManagedInstanceK8sSpecResponse `pulumi:"spec"`
}

type SqlManagedInstanceK8sRawResponseOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceK8sRawResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceK8sRawResponse)(nil)).Elem()
}

func (o SqlManagedInstanceK8sRawResponseOutput) ToSqlManagedInstanceK8sRawResponseOutput() SqlManagedInstanceK8sRawResponseOutput {
	return o
}

func (o SqlManagedInstanceK8sRawResponseOutput) ToSqlManagedInstanceK8sRawResponseOutputWithContext(ctx context.Context) SqlManagedInstanceK8sRawResponseOutput {
	return o
}

func (o SqlManagedInstanceK8sRawResponseOutput) Spec() SqlManagedInstanceK8sSpecResponsePtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceK8sRawResponse) *SqlManagedInstanceK8sSpecResponse { return v.Spec }).(SqlManagedInstanceK8sSpecResponsePtrOutput)
}

type SqlManagedInstanceK8sRawResponsePtrOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceK8sRawResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceK8sRawResponse)(nil)).Elem()
}

func (o SqlManagedInstanceK8sRawResponsePtrOutput) ToSqlManagedInstanceK8sRawResponsePtrOutput() SqlManagedInstanceK8sRawResponsePtrOutput {
	return o
}

func (o SqlManagedInstanceK8sRawResponsePtrOutput) ToSqlManagedInstanceK8sRawResponsePtrOutputWithContext(ctx context.Context) SqlManagedInstanceK8sRawResponsePtrOutput {
	return o
}

func (o SqlManagedInstanceK8sRawResponsePtrOutput) Elem() SqlManagedInstanceK8sRawResponseOutput {
	return o.ApplyT(func(v *SqlManagedInstanceK8sRawResponse) SqlManagedInstanceK8sRawResponse {
		if v != nil {
			return *v
		}
		var ret SqlManagedInstanceK8sRawResponse
		return ret
	}).(SqlManagedInstanceK8sRawResponseOutput)
}

func (o SqlManagedInstanceK8sRawResponsePtrOutput) Spec() SqlManagedInstanceK8sSpecResponsePtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceK8sRawResponse) *SqlManagedInstanceK8sSpecResponse {
		if v == nil {
			return nil
		}
		return v.Spec
	}).(SqlManagedInstanceK8sSpecResponsePtrOutput)
}

type SqlManagedInstanceK8sSpec struct {
	Replicas   *int           `pulumi:"replicas"`
	Scheduling *K8sScheduling `pulumi:"scheduling"`
}





type SqlManagedInstanceK8sSpecInput interface {
	pulumi.Input

	ToSqlManagedInstanceK8sSpecOutput() SqlManagedInstanceK8sSpecOutput
	ToSqlManagedInstanceK8sSpecOutputWithContext(context.Context) SqlManagedInstanceK8sSpecOutput
}

type SqlManagedInstanceK8sSpecArgs struct {
	Replicas   pulumi.IntPtrInput    `pulumi:"replicas"`
	Scheduling K8sSchedulingPtrInput `pulumi:"scheduling"`
}

func (SqlManagedInstanceK8sSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceK8sSpec)(nil)).Elem()
}

func (i SqlManagedInstanceK8sSpecArgs) ToSqlManagedInstanceK8sSpecOutput() SqlManagedInstanceK8sSpecOutput {
	return i.ToSqlManagedInstanceK8sSpecOutputWithContext(context.Background())
}

func (i SqlManagedInstanceK8sSpecArgs) ToSqlManagedInstanceK8sSpecOutputWithContext(ctx context.Context) SqlManagedInstanceK8sSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceK8sSpecOutput)
}

func (i SqlManagedInstanceK8sSpecArgs) ToSqlManagedInstanceK8sSpecPtrOutput() SqlManagedInstanceK8sSpecPtrOutput {
	return i.ToSqlManagedInstanceK8sSpecPtrOutputWithContext(context.Background())
}

func (i SqlManagedInstanceK8sSpecArgs) ToSqlManagedInstanceK8sSpecPtrOutputWithContext(ctx context.Context) SqlManagedInstanceK8sSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceK8sSpecOutput).ToSqlManagedInstanceK8sSpecPtrOutputWithContext(ctx)
}









type SqlManagedInstanceK8sSpecPtrInput interface {
	pulumi.Input

	ToSqlManagedInstanceK8sSpecPtrOutput() SqlManagedInstanceK8sSpecPtrOutput
	ToSqlManagedInstanceK8sSpecPtrOutputWithContext(context.Context) SqlManagedInstanceK8sSpecPtrOutput
}

type sqlManagedInstanceK8sSpecPtrType SqlManagedInstanceK8sSpecArgs

func SqlManagedInstanceK8sSpecPtr(v *SqlManagedInstanceK8sSpecArgs) SqlManagedInstanceK8sSpecPtrInput {
	return (*sqlManagedInstanceK8sSpecPtrType)(v)
}

func (*sqlManagedInstanceK8sSpecPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceK8sSpec)(nil)).Elem()
}

func (i *sqlManagedInstanceK8sSpecPtrType) ToSqlManagedInstanceK8sSpecPtrOutput() SqlManagedInstanceK8sSpecPtrOutput {
	return i.ToSqlManagedInstanceK8sSpecPtrOutputWithContext(context.Background())
}

func (i *sqlManagedInstanceK8sSpecPtrType) ToSqlManagedInstanceK8sSpecPtrOutputWithContext(ctx context.Context) SqlManagedInstanceK8sSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceK8sSpecPtrOutput)
}

type SqlManagedInstanceK8sSpecOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceK8sSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceK8sSpec)(nil)).Elem()
}

func (o SqlManagedInstanceK8sSpecOutput) ToSqlManagedInstanceK8sSpecOutput() SqlManagedInstanceK8sSpecOutput {
	return o
}

func (o SqlManagedInstanceK8sSpecOutput) ToSqlManagedInstanceK8sSpecOutputWithContext(ctx context.Context) SqlManagedInstanceK8sSpecOutput {
	return o
}

func (o SqlManagedInstanceK8sSpecOutput) ToSqlManagedInstanceK8sSpecPtrOutput() SqlManagedInstanceK8sSpecPtrOutput {
	return o.ToSqlManagedInstanceK8sSpecPtrOutputWithContext(context.Background())
}

func (o SqlManagedInstanceK8sSpecOutput) ToSqlManagedInstanceK8sSpecPtrOutputWithContext(ctx context.Context) SqlManagedInstanceK8sSpecPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlManagedInstanceK8sSpec) *SqlManagedInstanceK8sSpec {
		return &v
	}).(SqlManagedInstanceK8sSpecPtrOutput)
}

func (o SqlManagedInstanceK8sSpecOutput) Replicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceK8sSpec) *int { return v.Replicas }).(pulumi.IntPtrOutput)
}

func (o SqlManagedInstanceK8sSpecOutput) Scheduling() K8sSchedulingPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceK8sSpec) *K8sScheduling { return v.Scheduling }).(K8sSchedulingPtrOutput)
}

type SqlManagedInstanceK8sSpecPtrOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceK8sSpecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceK8sSpec)(nil)).Elem()
}

func (o SqlManagedInstanceK8sSpecPtrOutput) ToSqlManagedInstanceK8sSpecPtrOutput() SqlManagedInstanceK8sSpecPtrOutput {
	return o
}

func (o SqlManagedInstanceK8sSpecPtrOutput) ToSqlManagedInstanceK8sSpecPtrOutputWithContext(ctx context.Context) SqlManagedInstanceK8sSpecPtrOutput {
	return o
}

func (o SqlManagedInstanceK8sSpecPtrOutput) Elem() SqlManagedInstanceK8sSpecOutput {
	return o.ApplyT(func(v *SqlManagedInstanceK8sSpec) SqlManagedInstanceK8sSpec {
		if v != nil {
			return *v
		}
		var ret SqlManagedInstanceK8sSpec
		return ret
	}).(SqlManagedInstanceK8sSpecOutput)
}

func (o SqlManagedInstanceK8sSpecPtrOutput) Replicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceK8sSpec) *int {
		if v == nil {
			return nil
		}
		return v.Replicas
	}).(pulumi.IntPtrOutput)
}

func (o SqlManagedInstanceK8sSpecPtrOutput) Scheduling() K8sSchedulingPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceK8sSpec) *K8sScheduling {
		if v == nil {
			return nil
		}
		return v.Scheduling
	}).(K8sSchedulingPtrOutput)
}

type SqlManagedInstanceK8sSpecResponse struct {
	Replicas   *int                   `pulumi:"replicas"`
	Scheduling *K8sSchedulingResponse `pulumi:"scheduling"`
}

type SqlManagedInstanceK8sSpecResponseOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceK8sSpecResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceK8sSpecResponse)(nil)).Elem()
}

func (o SqlManagedInstanceK8sSpecResponseOutput) ToSqlManagedInstanceK8sSpecResponseOutput() SqlManagedInstanceK8sSpecResponseOutput {
	return o
}

func (o SqlManagedInstanceK8sSpecResponseOutput) ToSqlManagedInstanceK8sSpecResponseOutputWithContext(ctx context.Context) SqlManagedInstanceK8sSpecResponseOutput {
	return o
}

func (o SqlManagedInstanceK8sSpecResponseOutput) Replicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceK8sSpecResponse) *int { return v.Replicas }).(pulumi.IntPtrOutput)
}

func (o SqlManagedInstanceK8sSpecResponseOutput) Scheduling() K8sSchedulingResponsePtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceK8sSpecResponse) *K8sSchedulingResponse { return v.Scheduling }).(K8sSchedulingResponsePtrOutput)
}

type SqlManagedInstanceK8sSpecResponsePtrOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceK8sSpecResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceK8sSpecResponse)(nil)).Elem()
}

func (o SqlManagedInstanceK8sSpecResponsePtrOutput) ToSqlManagedInstanceK8sSpecResponsePtrOutput() SqlManagedInstanceK8sSpecResponsePtrOutput {
	return o
}

func (o SqlManagedInstanceK8sSpecResponsePtrOutput) ToSqlManagedInstanceK8sSpecResponsePtrOutputWithContext(ctx context.Context) SqlManagedInstanceK8sSpecResponsePtrOutput {
	return o
}

func (o SqlManagedInstanceK8sSpecResponsePtrOutput) Elem() SqlManagedInstanceK8sSpecResponseOutput {
	return o.ApplyT(func(v *SqlManagedInstanceK8sSpecResponse) SqlManagedInstanceK8sSpecResponse {
		if v != nil {
			return *v
		}
		var ret SqlManagedInstanceK8sSpecResponse
		return ret
	}).(SqlManagedInstanceK8sSpecResponseOutput)
}

func (o SqlManagedInstanceK8sSpecResponsePtrOutput) Replicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceK8sSpecResponse) *int {
		if v == nil {
			return nil
		}
		return v.Replicas
	}).(pulumi.IntPtrOutput)
}

func (o SqlManagedInstanceK8sSpecResponsePtrOutput) Scheduling() K8sSchedulingResponsePtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceK8sSpecResponse) *K8sSchedulingResponse {
		if v == nil {
			return nil
		}
		return v.Scheduling
	}).(K8sSchedulingResponsePtrOutput)
}

type SqlManagedInstanceProperties struct {
	ActiveDirectoryInformation *ActiveDirectoryInformation `pulumi:"activeDirectoryInformation"`
	Admin                      *string                     `pulumi:"admin"`
	BasicLoginInformation      *BasicLoginInformation      `pulumi:"basicLoginInformation"`
	ClusterId                  *string                     `pulumi:"clusterId"`
	DataControllerId           *string                     `pulumi:"dataControllerId"`
	EndTime                    *string                     `pulumi:"endTime"`
	ExtensionId                *string                     `pulumi:"extensionId"`
	K8sRaw                     *SqlManagedInstanceK8sRaw   `pulumi:"k8sRaw"`
	LastUploadedDate           *string                     `pulumi:"lastUploadedDate"`
	LicenseType                *string                     `pulumi:"licenseType"`
	StartTime                  *string                     `pulumi:"startTime"`
}


func (val *SqlManagedInstanceProperties) Defaults() *SqlManagedInstanceProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.LicenseType) {
		licenseType_ := "BasePrice"
		tmp.LicenseType = &licenseType_
	}
	return &tmp
}





type SqlManagedInstancePropertiesInput interface {
	pulumi.Input

	ToSqlManagedInstancePropertiesOutput() SqlManagedInstancePropertiesOutput
	ToSqlManagedInstancePropertiesOutputWithContext(context.Context) SqlManagedInstancePropertiesOutput
}

type SqlManagedInstancePropertiesArgs struct {
	ActiveDirectoryInformation ActiveDirectoryInformationPtrInput `pulumi:"activeDirectoryInformation"`
	Admin                      pulumi.StringPtrInput              `pulumi:"admin"`
	BasicLoginInformation      BasicLoginInformationPtrInput      `pulumi:"basicLoginInformation"`
	ClusterId                  pulumi.StringPtrInput              `pulumi:"clusterId"`
	DataControllerId           pulumi.StringPtrInput              `pulumi:"dataControllerId"`
	EndTime                    pulumi.StringPtrInput              `pulumi:"endTime"`
	ExtensionId                pulumi.StringPtrInput              `pulumi:"extensionId"`
	K8sRaw                     SqlManagedInstanceK8sRawPtrInput   `pulumi:"k8sRaw"`
	LastUploadedDate           pulumi.StringPtrInput              `pulumi:"lastUploadedDate"`
	LicenseType                pulumi.StringPtrInput              `pulumi:"licenseType"`
	StartTime                  pulumi.StringPtrInput              `pulumi:"startTime"`
}


func (val *SqlManagedInstancePropertiesArgs) Defaults() *SqlManagedInstancePropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.LicenseType) {
		tmp.LicenseType = pulumi.StringPtr("BasePrice")
	}
	return &tmp
}
func (SqlManagedInstancePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceProperties)(nil)).Elem()
}

func (i SqlManagedInstancePropertiesArgs) ToSqlManagedInstancePropertiesOutput() SqlManagedInstancePropertiesOutput {
	return i.ToSqlManagedInstancePropertiesOutputWithContext(context.Background())
}

func (i SqlManagedInstancePropertiesArgs) ToSqlManagedInstancePropertiesOutputWithContext(ctx context.Context) SqlManagedInstancePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstancePropertiesOutput)
}

type SqlManagedInstancePropertiesOutput struct{ *pulumi.OutputState }

func (SqlManagedInstancePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceProperties)(nil)).Elem()
}

func (o SqlManagedInstancePropertiesOutput) ToSqlManagedInstancePropertiesOutput() SqlManagedInstancePropertiesOutput {
	return o
}

func (o SqlManagedInstancePropertiesOutput) ToSqlManagedInstancePropertiesOutputWithContext(ctx context.Context) SqlManagedInstancePropertiesOutput {
	return o
}

func (o SqlManagedInstancePropertiesOutput) ActiveDirectoryInformation() ActiveDirectoryInformationPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *ActiveDirectoryInformation { return v.ActiveDirectoryInformation }).(ActiveDirectoryInformationPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) Admin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *string { return v.Admin }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) BasicLoginInformation() BasicLoginInformationPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *BasicLoginInformation { return v.BasicLoginInformation }).(BasicLoginInformationPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) DataControllerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *string { return v.DataControllerId }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) ExtensionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *string { return v.ExtensionId }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) K8sRaw() SqlManagedInstanceK8sRawPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *SqlManagedInstanceK8sRaw { return v.K8sRaw }).(SqlManagedInstanceK8sRawPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *string { return v.LastUploadedDate }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type SqlManagedInstancePropertiesResponse struct {
	Admin                 *string                           `pulumi:"admin"`
	BasicLoginInformation *BasicLoginInformationResponse    `pulumi:"basicLoginInformation"`
	ClusterId             *string                           `pulumi:"clusterId"`
	DataControllerId      *string                           `pulumi:"dataControllerId"`
	EndTime               *string                           `pulumi:"endTime"`
	ExtensionId           *string                           `pulumi:"extensionId"`
	K8sRaw                *SqlManagedInstanceK8sRawResponse `pulumi:"k8sRaw"`
	LastUploadedDate      *string                           `pulumi:"lastUploadedDate"`
	LicenseType           *string                           `pulumi:"licenseType"`
	ProvisioningState     string                            `pulumi:"provisioningState"`
	StartTime             *string                           `pulumi:"startTime"`
}


func (val *SqlManagedInstancePropertiesResponse) Defaults() *SqlManagedInstancePropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.LicenseType) {
		licenseType_ := "BasePrice"
		tmp.LicenseType = &licenseType_
	}
	return &tmp
}

type SqlManagedInstancePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SqlManagedInstancePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstancePropertiesResponse)(nil)).Elem()
}

func (o SqlManagedInstancePropertiesResponseOutput) ToSqlManagedInstancePropertiesResponseOutput() SqlManagedInstancePropertiesResponseOutput {
	return o
}

func (o SqlManagedInstancePropertiesResponseOutput) ToSqlManagedInstancePropertiesResponseOutputWithContext(ctx context.Context) SqlManagedInstancePropertiesResponseOutput {
	return o
}

func (o SqlManagedInstancePropertiesResponseOutput) Admin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *string { return v.Admin }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) BasicLoginInformation() BasicLoginInformationResponsePtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *BasicLoginInformationResponse {
		return v.BasicLoginInformation
	}).(BasicLoginInformationResponsePtrOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) DataControllerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *string { return v.DataControllerId }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) ExtensionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *string { return v.ExtensionId }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) K8sRaw() SqlManagedInstanceK8sRawResponsePtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *SqlManagedInstanceK8sRawResponse { return v.K8sRaw }).(SqlManagedInstanceK8sRawResponsePtrOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *string { return v.LastUploadedDate }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type SqlManagedInstanceSku struct {
	Capacity *int                       `pulumi:"capacity"`
	Dev      *bool                      `pulumi:"dev"`
	Family   *string                    `pulumi:"family"`
	Name     SqlManagedInstanceSkuName  `pulumi:"name"`
	Size     *string                    `pulumi:"size"`
	Tier     *SqlManagedInstanceSkuTier `pulumi:"tier"`
}


func (val *SqlManagedInstanceSku) Defaults() *SqlManagedInstanceSku {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Dev) {
		dev_ := true
		tmp.Dev = &dev_
	}
	if isZero(tmp.Tier) {
		tier_ := SqlManagedInstanceSkuTier("GeneralPurpose")
		tmp.Tier = &tier_
	}
	return &tmp
}





type SqlManagedInstanceSkuInput interface {
	pulumi.Input

	ToSqlManagedInstanceSkuOutput() SqlManagedInstanceSkuOutput
	ToSqlManagedInstanceSkuOutputWithContext(context.Context) SqlManagedInstanceSkuOutput
}

type SqlManagedInstanceSkuArgs struct {
	Capacity pulumi.IntPtrInput                `pulumi:"capacity"`
	Dev      pulumi.BoolPtrInput               `pulumi:"dev"`
	Family   pulumi.StringPtrInput             `pulumi:"family"`
	Name     SqlManagedInstanceSkuNameInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput             `pulumi:"size"`
	Tier     SqlManagedInstanceSkuTierPtrInput `pulumi:"tier"`
}


func (val *SqlManagedInstanceSkuArgs) Defaults() *SqlManagedInstanceSkuArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Dev) {
		tmp.Dev = pulumi.BoolPtr(true)
	}
	if isZero(tmp.Tier) {
		tmp.Tier = SqlManagedInstanceSkuTier("GeneralPurpose")
	}
	return &tmp
}
func (SqlManagedInstanceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceSku)(nil)).Elem()
}

func (i SqlManagedInstanceSkuArgs) ToSqlManagedInstanceSkuOutput() SqlManagedInstanceSkuOutput {
	return i.ToSqlManagedInstanceSkuOutputWithContext(context.Background())
}

func (i SqlManagedInstanceSkuArgs) ToSqlManagedInstanceSkuOutputWithContext(ctx context.Context) SqlManagedInstanceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceSkuOutput)
}

func (i SqlManagedInstanceSkuArgs) ToSqlManagedInstanceSkuPtrOutput() SqlManagedInstanceSkuPtrOutput {
	return i.ToSqlManagedInstanceSkuPtrOutputWithContext(context.Background())
}

func (i SqlManagedInstanceSkuArgs) ToSqlManagedInstanceSkuPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceSkuOutput).ToSqlManagedInstanceSkuPtrOutputWithContext(ctx)
}









type SqlManagedInstanceSkuPtrInput interface {
	pulumi.Input

	ToSqlManagedInstanceSkuPtrOutput() SqlManagedInstanceSkuPtrOutput
	ToSqlManagedInstanceSkuPtrOutputWithContext(context.Context) SqlManagedInstanceSkuPtrOutput
}

type sqlManagedInstanceSkuPtrType SqlManagedInstanceSkuArgs

func SqlManagedInstanceSkuPtr(v *SqlManagedInstanceSkuArgs) SqlManagedInstanceSkuPtrInput {
	return (*sqlManagedInstanceSkuPtrType)(v)
}

func (*sqlManagedInstanceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceSku)(nil)).Elem()
}

func (i *sqlManagedInstanceSkuPtrType) ToSqlManagedInstanceSkuPtrOutput() SqlManagedInstanceSkuPtrOutput {
	return i.ToSqlManagedInstanceSkuPtrOutputWithContext(context.Background())
}

func (i *sqlManagedInstanceSkuPtrType) ToSqlManagedInstanceSkuPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceSkuPtrOutput)
}

type SqlManagedInstanceSkuOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceSku)(nil)).Elem()
}

func (o SqlManagedInstanceSkuOutput) ToSqlManagedInstanceSkuOutput() SqlManagedInstanceSkuOutput {
	return o
}

func (o SqlManagedInstanceSkuOutput) ToSqlManagedInstanceSkuOutputWithContext(ctx context.Context) SqlManagedInstanceSkuOutput {
	return o
}

func (o SqlManagedInstanceSkuOutput) ToSqlManagedInstanceSkuPtrOutput() SqlManagedInstanceSkuPtrOutput {
	return o.ToSqlManagedInstanceSkuPtrOutputWithContext(context.Background())
}

func (o SqlManagedInstanceSkuOutput) ToSqlManagedInstanceSkuPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlManagedInstanceSku) *SqlManagedInstanceSku {
		return &v
	}).(SqlManagedInstanceSkuPtrOutput)
}

func (o SqlManagedInstanceSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SqlManagedInstanceSkuOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSku) *bool { return v.Dev }).(pulumi.BoolPtrOutput)
}

func (o SqlManagedInstanceSkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuOutput) Name() SqlManagedInstanceSkuNameOutput {
	return o.ApplyT(func(v SqlManagedInstanceSku) SqlManagedInstanceSkuName { return v.Name }).(SqlManagedInstanceSkuNameOutput)
}

func (o SqlManagedInstanceSkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuOutput) Tier() SqlManagedInstanceSkuTierPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSku) *SqlManagedInstanceSkuTier { return v.Tier }).(SqlManagedInstanceSkuTierPtrOutput)
}

type SqlManagedInstanceSkuPtrOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceSku)(nil)).Elem()
}

func (o SqlManagedInstanceSkuPtrOutput) ToSqlManagedInstanceSkuPtrOutput() SqlManagedInstanceSkuPtrOutput {
	return o
}

func (o SqlManagedInstanceSkuPtrOutput) ToSqlManagedInstanceSkuPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuPtrOutput {
	return o
}

func (o SqlManagedInstanceSkuPtrOutput) Elem() SqlManagedInstanceSkuOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSku) SqlManagedInstanceSku {
		if v != nil {
			return *v
		}
		var ret SqlManagedInstanceSku
		return ret
	}).(SqlManagedInstanceSkuOutput)
}

func (o SqlManagedInstanceSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SqlManagedInstanceSkuPtrOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSku) *bool {
		if v == nil {
			return nil
		}
		return v.Dev
	}).(pulumi.BoolPtrOutput)
}

func (o SqlManagedInstanceSkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuPtrOutput) Name() SqlManagedInstanceSkuNamePtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSku) *SqlManagedInstanceSkuName {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(SqlManagedInstanceSkuNamePtrOutput)
}

func (o SqlManagedInstanceSkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuPtrOutput) Tier() SqlManagedInstanceSkuTierPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSku) *SqlManagedInstanceSkuTier {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(SqlManagedInstanceSkuTierPtrOutput)
}

type SqlManagedInstanceSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Dev      *bool   `pulumi:"dev"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}


func (val *SqlManagedInstanceSkuResponse) Defaults() *SqlManagedInstanceSkuResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Dev) {
		dev_ := true
		tmp.Dev = &dev_
	}
	if isZero(tmp.Tier) {
		tier_ := "GeneralPurpose"
		tmp.Tier = &tier_
	}
	return &tmp
}

type SqlManagedInstanceSkuResponseOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceSkuResponse)(nil)).Elem()
}

func (o SqlManagedInstanceSkuResponseOutput) ToSqlManagedInstanceSkuResponseOutput() SqlManagedInstanceSkuResponseOutput {
	return o
}

func (o SqlManagedInstanceSkuResponseOutput) ToSqlManagedInstanceSkuResponseOutputWithContext(ctx context.Context) SqlManagedInstanceSkuResponseOutput {
	return o
}

func (o SqlManagedInstanceSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SqlManagedInstanceSkuResponseOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSkuResponse) *bool { return v.Dev }).(pulumi.BoolPtrOutput)
}

func (o SqlManagedInstanceSkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SqlManagedInstanceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SqlManagedInstanceSkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SqlManagedInstanceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceSkuResponse)(nil)).Elem()
}

func (o SqlManagedInstanceSkuResponsePtrOutput) ToSqlManagedInstanceSkuResponsePtrOutput() SqlManagedInstanceSkuResponsePtrOutput {
	return o
}

func (o SqlManagedInstanceSkuResponsePtrOutput) ToSqlManagedInstanceSkuResponsePtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuResponsePtrOutput {
	return o
}

func (o SqlManagedInstanceSkuResponsePtrOutput) Elem() SqlManagedInstanceSkuResponseOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuResponse) SqlManagedInstanceSkuResponse {
		if v != nil {
			return *v
		}
		var ret SqlManagedInstanceSkuResponse
		return ret
	}).(SqlManagedInstanceSkuResponseOutput)
}

func (o SqlManagedInstanceSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SqlManagedInstanceSkuResponsePtrOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Dev
	}).(pulumi.BoolPtrOutput)
}

func (o SqlManagedInstanceSkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SqlServerInstanceProperties struct {
	AzureDefenderStatus            *string `pulumi:"azureDefenderStatus"`
	AzureDefenderStatusLastUpdated *string `pulumi:"azureDefenderStatusLastUpdated"`
	Collation                      *string `pulumi:"collation"`
	ContainerResourceId            string  `pulumi:"containerResourceId"`
	CurrentVersion                 *string `pulumi:"currentVersion"`
	Edition                        *string `pulumi:"edition"`
	HostType                       *string `pulumi:"hostType"`
	InstanceName                   *string `pulumi:"instanceName"`
	LicenseType                    *string `pulumi:"licenseType"`
	PatchLevel                     *string `pulumi:"patchLevel"`
	ProductId                      *string `pulumi:"productId"`
	Status                         string  `pulumi:"status"`
	TcpDynamicPorts                *string `pulumi:"tcpDynamicPorts"`
	TcpStaticPorts                 *string `pulumi:"tcpStaticPorts"`
	VCore                          *string `pulumi:"vCore"`
	Version                        *string `pulumi:"version"`
}





type SqlServerInstancePropertiesInput interface {
	pulumi.Input

	ToSqlServerInstancePropertiesOutput() SqlServerInstancePropertiesOutput
	ToSqlServerInstancePropertiesOutputWithContext(context.Context) SqlServerInstancePropertiesOutput
}

type SqlServerInstancePropertiesArgs struct {
	AzureDefenderStatus            pulumi.StringPtrInput `pulumi:"azureDefenderStatus"`
	AzureDefenderStatusLastUpdated pulumi.StringPtrInput `pulumi:"azureDefenderStatusLastUpdated"`
	Collation                      pulumi.StringPtrInput `pulumi:"collation"`
	ContainerResourceId            pulumi.StringInput    `pulumi:"containerResourceId"`
	CurrentVersion                 pulumi.StringPtrInput `pulumi:"currentVersion"`
	Edition                        pulumi.StringPtrInput `pulumi:"edition"`
	HostType                       pulumi.StringPtrInput `pulumi:"hostType"`
	InstanceName                   pulumi.StringPtrInput `pulumi:"instanceName"`
	LicenseType                    pulumi.StringPtrInput `pulumi:"licenseType"`
	PatchLevel                     pulumi.StringPtrInput `pulumi:"patchLevel"`
	ProductId                      pulumi.StringPtrInput `pulumi:"productId"`
	Status                         pulumi.StringInput    `pulumi:"status"`
	TcpDynamicPorts                pulumi.StringPtrInput `pulumi:"tcpDynamicPorts"`
	TcpStaticPorts                 pulumi.StringPtrInput `pulumi:"tcpStaticPorts"`
	VCore                          pulumi.StringPtrInput `pulumi:"vCore"`
	Version                        pulumi.StringPtrInput `pulumi:"version"`
}

func (SqlServerInstancePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerInstanceProperties)(nil)).Elem()
}

func (i SqlServerInstancePropertiesArgs) ToSqlServerInstancePropertiesOutput() SqlServerInstancePropertiesOutput {
	return i.ToSqlServerInstancePropertiesOutputWithContext(context.Background())
}

func (i SqlServerInstancePropertiesArgs) ToSqlServerInstancePropertiesOutputWithContext(ctx context.Context) SqlServerInstancePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerInstancePropertiesOutput)
}

func (i SqlServerInstancePropertiesArgs) ToSqlServerInstancePropertiesPtrOutput() SqlServerInstancePropertiesPtrOutput {
	return i.ToSqlServerInstancePropertiesPtrOutputWithContext(context.Background())
}

func (i SqlServerInstancePropertiesArgs) ToSqlServerInstancePropertiesPtrOutputWithContext(ctx context.Context) SqlServerInstancePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerInstancePropertiesOutput).ToSqlServerInstancePropertiesPtrOutputWithContext(ctx)
}









type SqlServerInstancePropertiesPtrInput interface {
	pulumi.Input

	ToSqlServerInstancePropertiesPtrOutput() SqlServerInstancePropertiesPtrOutput
	ToSqlServerInstancePropertiesPtrOutputWithContext(context.Context) SqlServerInstancePropertiesPtrOutput
}

type sqlServerInstancePropertiesPtrType SqlServerInstancePropertiesArgs

func SqlServerInstancePropertiesPtr(v *SqlServerInstancePropertiesArgs) SqlServerInstancePropertiesPtrInput {
	return (*sqlServerInstancePropertiesPtrType)(v)
}

func (*sqlServerInstancePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerInstanceProperties)(nil)).Elem()
}

func (i *sqlServerInstancePropertiesPtrType) ToSqlServerInstancePropertiesPtrOutput() SqlServerInstancePropertiesPtrOutput {
	return i.ToSqlServerInstancePropertiesPtrOutputWithContext(context.Background())
}

func (i *sqlServerInstancePropertiesPtrType) ToSqlServerInstancePropertiesPtrOutputWithContext(ctx context.Context) SqlServerInstancePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerInstancePropertiesPtrOutput)
}

type SqlServerInstancePropertiesOutput struct{ *pulumi.OutputState }

func (SqlServerInstancePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerInstanceProperties)(nil)).Elem()
}

func (o SqlServerInstancePropertiesOutput) ToSqlServerInstancePropertiesOutput() SqlServerInstancePropertiesOutput {
	return o
}

func (o SqlServerInstancePropertiesOutput) ToSqlServerInstancePropertiesOutputWithContext(ctx context.Context) SqlServerInstancePropertiesOutput {
	return o
}

func (o SqlServerInstancePropertiesOutput) ToSqlServerInstancePropertiesPtrOutput() SqlServerInstancePropertiesPtrOutput {
	return o.ToSqlServerInstancePropertiesPtrOutputWithContext(context.Background())
}

func (o SqlServerInstancePropertiesOutput) ToSqlServerInstancePropertiesPtrOutputWithContext(ctx context.Context) SqlServerInstancePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlServerInstanceProperties) *SqlServerInstanceProperties {
		return &v
	}).(SqlServerInstancePropertiesPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) AzureDefenderStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.AzureDefenderStatus }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) AzureDefenderStatusLastUpdated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.AzureDefenderStatusLastUpdated }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) ContainerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) string { return v.ContainerResourceId }).(pulumi.StringOutput)
}

func (o SqlServerInstancePropertiesOutput) CurrentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.CurrentVersion }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.Edition }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) HostType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.HostType }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.InstanceName }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) PatchLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.PatchLevel }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) ProductId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.ProductId }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) string { return v.Status }).(pulumi.StringOutput)
}

func (o SqlServerInstancePropertiesOutput) TcpDynamicPorts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.TcpDynamicPorts }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) TcpStaticPorts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.TcpStaticPorts }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) VCore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.VCore }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type SqlServerInstancePropertiesPtrOutput struct{ *pulumi.OutputState }

func (SqlServerInstancePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerInstanceProperties)(nil)).Elem()
}

func (o SqlServerInstancePropertiesPtrOutput) ToSqlServerInstancePropertiesPtrOutput() SqlServerInstancePropertiesPtrOutput {
	return o
}

func (o SqlServerInstancePropertiesPtrOutput) ToSqlServerInstancePropertiesPtrOutputWithContext(ctx context.Context) SqlServerInstancePropertiesPtrOutput {
	return o
}

func (o SqlServerInstancePropertiesPtrOutput) Elem() SqlServerInstancePropertiesOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) SqlServerInstanceProperties {
		if v != nil {
			return *v
		}
		var ret SqlServerInstanceProperties
		return ret
	}).(SqlServerInstancePropertiesOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) AzureDefenderStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.AzureDefenderStatus
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) AzureDefenderStatusLastUpdated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.AzureDefenderStatusLastUpdated
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Collation
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) ContainerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ContainerResourceId
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) CurrentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.CurrentVersion
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Edition
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) HostType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.HostType
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.InstanceName
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.LicenseType
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) PatchLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.PatchLevel
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) ProductId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProductId
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) TcpDynamicPorts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.TcpDynamicPorts
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) TcpStaticPorts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.TcpStaticPorts
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) VCore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.VCore
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type SqlServerInstancePropertiesResponse struct {
	AzureDefenderStatus            *string `pulumi:"azureDefenderStatus"`
	AzureDefenderStatusLastUpdated *string `pulumi:"azureDefenderStatusLastUpdated"`
	Collation                      *string `pulumi:"collation"`
	ContainerResourceId            string  `pulumi:"containerResourceId"`
	CreateTime                     string  `pulumi:"createTime"`
	CurrentVersion                 *string `pulumi:"currentVersion"`
	Edition                        *string `pulumi:"edition"`
	HostType                       *string `pulumi:"hostType"`
	InstanceName                   *string `pulumi:"instanceName"`
	LicenseType                    *string `pulumi:"licenseType"`
	PatchLevel                     *string `pulumi:"patchLevel"`
	ProductId                      *string `pulumi:"productId"`
	ProvisioningState              string  `pulumi:"provisioningState"`
	Status                         string  `pulumi:"status"`
	TcpDynamicPorts                *string `pulumi:"tcpDynamicPorts"`
	TcpStaticPorts                 *string `pulumi:"tcpStaticPorts"`
	VCore                          *string `pulumi:"vCore"`
	Version                        *string `pulumi:"version"`
}

type SqlServerInstancePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SqlServerInstancePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerInstancePropertiesResponse)(nil)).Elem()
}

func (o SqlServerInstancePropertiesResponseOutput) ToSqlServerInstancePropertiesResponseOutput() SqlServerInstancePropertiesResponseOutput {
	return o
}

func (o SqlServerInstancePropertiesResponseOutput) ToSqlServerInstancePropertiesResponseOutputWithContext(ctx context.Context) SqlServerInstancePropertiesResponseOutput {
	return o
}

func (o SqlServerInstancePropertiesResponseOutput) AzureDefenderStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.AzureDefenderStatus }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) AzureDefenderStatusLastUpdated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.AzureDefenderStatusLastUpdated }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) ContainerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) string { return v.ContainerResourceId }).(pulumi.StringOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) CurrentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.CurrentVersion }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.Edition }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) HostType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.HostType }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.InstanceName }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) PatchLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.PatchLevel }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) ProductId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.ProductId }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) TcpDynamicPorts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.TcpDynamicPorts }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) TcpStaticPorts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.TcpStaticPorts }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) VCore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.VCore }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
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

type UploadServicePrincipal struct {
	Authority    *string `pulumi:"authority"`
	ClientId     *string `pulumi:"clientId"`
	ClientSecret *string `pulumi:"clientSecret"`
	TenantId     *string `pulumi:"tenantId"`
}





type UploadServicePrincipalInput interface {
	pulumi.Input

	ToUploadServicePrincipalOutput() UploadServicePrincipalOutput
	ToUploadServicePrincipalOutputWithContext(context.Context) UploadServicePrincipalOutput
}

type UploadServicePrincipalArgs struct {
	Authority    pulumi.StringPtrInput `pulumi:"authority"`
	ClientId     pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecret pulumi.StringPtrInput `pulumi:"clientSecret"`
	TenantId     pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (UploadServicePrincipalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadServicePrincipal)(nil)).Elem()
}

func (i UploadServicePrincipalArgs) ToUploadServicePrincipalOutput() UploadServicePrincipalOutput {
	return i.ToUploadServicePrincipalOutputWithContext(context.Background())
}

func (i UploadServicePrincipalArgs) ToUploadServicePrincipalOutputWithContext(ctx context.Context) UploadServicePrincipalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadServicePrincipalOutput)
}

func (i UploadServicePrincipalArgs) ToUploadServicePrincipalPtrOutput() UploadServicePrincipalPtrOutput {
	return i.ToUploadServicePrincipalPtrOutputWithContext(context.Background())
}

func (i UploadServicePrincipalArgs) ToUploadServicePrincipalPtrOutputWithContext(ctx context.Context) UploadServicePrincipalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadServicePrincipalOutput).ToUploadServicePrincipalPtrOutputWithContext(ctx)
}









type UploadServicePrincipalPtrInput interface {
	pulumi.Input

	ToUploadServicePrincipalPtrOutput() UploadServicePrincipalPtrOutput
	ToUploadServicePrincipalPtrOutputWithContext(context.Context) UploadServicePrincipalPtrOutput
}

type uploadServicePrincipalPtrType UploadServicePrincipalArgs

func UploadServicePrincipalPtr(v *UploadServicePrincipalArgs) UploadServicePrincipalPtrInput {
	return (*uploadServicePrincipalPtrType)(v)
}

func (*uploadServicePrincipalPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadServicePrincipal)(nil)).Elem()
}

func (i *uploadServicePrincipalPtrType) ToUploadServicePrincipalPtrOutput() UploadServicePrincipalPtrOutput {
	return i.ToUploadServicePrincipalPtrOutputWithContext(context.Background())
}

func (i *uploadServicePrincipalPtrType) ToUploadServicePrincipalPtrOutputWithContext(ctx context.Context) UploadServicePrincipalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadServicePrincipalPtrOutput)
}

type UploadServicePrincipalOutput struct{ *pulumi.OutputState }

func (UploadServicePrincipalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadServicePrincipal)(nil)).Elem()
}

func (o UploadServicePrincipalOutput) ToUploadServicePrincipalOutput() UploadServicePrincipalOutput {
	return o
}

func (o UploadServicePrincipalOutput) ToUploadServicePrincipalOutputWithContext(ctx context.Context) UploadServicePrincipalOutput {
	return o
}

func (o UploadServicePrincipalOutput) ToUploadServicePrincipalPtrOutput() UploadServicePrincipalPtrOutput {
	return o.ToUploadServicePrincipalPtrOutputWithContext(context.Background())
}

func (o UploadServicePrincipalOutput) ToUploadServicePrincipalPtrOutputWithContext(ctx context.Context) UploadServicePrincipalPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UploadServicePrincipal) *UploadServicePrincipal {
		return &v
	}).(UploadServicePrincipalPtrOutput)
}

func (o UploadServicePrincipalOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadServicePrincipal) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadServicePrincipal) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadServicePrincipal) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadServicePrincipal) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type UploadServicePrincipalPtrOutput struct{ *pulumi.OutputState }

func (UploadServicePrincipalPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadServicePrincipal)(nil)).Elem()
}

func (o UploadServicePrincipalPtrOutput) ToUploadServicePrincipalPtrOutput() UploadServicePrincipalPtrOutput {
	return o
}

func (o UploadServicePrincipalPtrOutput) ToUploadServicePrincipalPtrOutputWithContext(ctx context.Context) UploadServicePrincipalPtrOutput {
	return o
}

func (o UploadServicePrincipalPtrOutput) Elem() UploadServicePrincipalOutput {
	return o.ApplyT(func(v *UploadServicePrincipal) UploadServicePrincipal {
		if v != nil {
			return *v
		}
		var ret UploadServicePrincipal
		return ret
	}).(UploadServicePrincipalOutput)
}

func (o UploadServicePrincipalPtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadServicePrincipal) *string {
		if v == nil {
			return nil
		}
		return v.Authority
	}).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadServicePrincipal) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalPtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadServicePrincipal) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadServicePrincipal) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type UploadServicePrincipalResponse struct {
	Authority *string `pulumi:"authority"`
	ClientId  *string `pulumi:"clientId"`
	TenantId  *string `pulumi:"tenantId"`
}

type UploadServicePrincipalResponseOutput struct{ *pulumi.OutputState }

func (UploadServicePrincipalResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadServicePrincipalResponse)(nil)).Elem()
}

func (o UploadServicePrincipalResponseOutput) ToUploadServicePrincipalResponseOutput() UploadServicePrincipalResponseOutput {
	return o
}

func (o UploadServicePrincipalResponseOutput) ToUploadServicePrincipalResponseOutputWithContext(ctx context.Context) UploadServicePrincipalResponseOutput {
	return o
}

func (o UploadServicePrincipalResponseOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadServicePrincipalResponse) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadServicePrincipalResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadServicePrincipalResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type UploadServicePrincipalResponsePtrOutput struct{ *pulumi.OutputState }

func (UploadServicePrincipalResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadServicePrincipalResponse)(nil)).Elem()
}

func (o UploadServicePrincipalResponsePtrOutput) ToUploadServicePrincipalResponsePtrOutput() UploadServicePrincipalResponsePtrOutput {
	return o
}

func (o UploadServicePrincipalResponsePtrOutput) ToUploadServicePrincipalResponsePtrOutputWithContext(ctx context.Context) UploadServicePrincipalResponsePtrOutput {
	return o
}

func (o UploadServicePrincipalResponsePtrOutput) Elem() UploadServicePrincipalResponseOutput {
	return o.ApplyT(func(v *UploadServicePrincipalResponse) UploadServicePrincipalResponse {
		if v != nil {
			return *v
		}
		var ret UploadServicePrincipalResponse
		return ret
	}).(UploadServicePrincipalResponseOutput)
}

func (o UploadServicePrincipalResponsePtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadServicePrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return v.Authority
	}).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadServicePrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadServicePrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type UploadWatermark struct {
	Logs    *string `pulumi:"logs"`
	Metrics *string `pulumi:"metrics"`
	Usages  *string `pulumi:"usages"`
}





type UploadWatermarkInput interface {
	pulumi.Input

	ToUploadWatermarkOutput() UploadWatermarkOutput
	ToUploadWatermarkOutputWithContext(context.Context) UploadWatermarkOutput
}

type UploadWatermarkArgs struct {
	Logs    pulumi.StringPtrInput `pulumi:"logs"`
	Metrics pulumi.StringPtrInput `pulumi:"metrics"`
	Usages  pulumi.StringPtrInput `pulumi:"usages"`
}

func (UploadWatermarkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadWatermark)(nil)).Elem()
}

func (i UploadWatermarkArgs) ToUploadWatermarkOutput() UploadWatermarkOutput {
	return i.ToUploadWatermarkOutputWithContext(context.Background())
}

func (i UploadWatermarkArgs) ToUploadWatermarkOutputWithContext(ctx context.Context) UploadWatermarkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadWatermarkOutput)
}

func (i UploadWatermarkArgs) ToUploadWatermarkPtrOutput() UploadWatermarkPtrOutput {
	return i.ToUploadWatermarkPtrOutputWithContext(context.Background())
}

func (i UploadWatermarkArgs) ToUploadWatermarkPtrOutputWithContext(ctx context.Context) UploadWatermarkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadWatermarkOutput).ToUploadWatermarkPtrOutputWithContext(ctx)
}









type UploadWatermarkPtrInput interface {
	pulumi.Input

	ToUploadWatermarkPtrOutput() UploadWatermarkPtrOutput
	ToUploadWatermarkPtrOutputWithContext(context.Context) UploadWatermarkPtrOutput
}

type uploadWatermarkPtrType UploadWatermarkArgs

func UploadWatermarkPtr(v *UploadWatermarkArgs) UploadWatermarkPtrInput {
	return (*uploadWatermarkPtrType)(v)
}

func (*uploadWatermarkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadWatermark)(nil)).Elem()
}

func (i *uploadWatermarkPtrType) ToUploadWatermarkPtrOutput() UploadWatermarkPtrOutput {
	return i.ToUploadWatermarkPtrOutputWithContext(context.Background())
}

func (i *uploadWatermarkPtrType) ToUploadWatermarkPtrOutputWithContext(ctx context.Context) UploadWatermarkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadWatermarkPtrOutput)
}

type UploadWatermarkOutput struct{ *pulumi.OutputState }

func (UploadWatermarkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadWatermark)(nil)).Elem()
}

func (o UploadWatermarkOutput) ToUploadWatermarkOutput() UploadWatermarkOutput {
	return o
}

func (o UploadWatermarkOutput) ToUploadWatermarkOutputWithContext(ctx context.Context) UploadWatermarkOutput {
	return o
}

func (o UploadWatermarkOutput) ToUploadWatermarkPtrOutput() UploadWatermarkPtrOutput {
	return o.ToUploadWatermarkPtrOutputWithContext(context.Background())
}

func (o UploadWatermarkOutput) ToUploadWatermarkPtrOutputWithContext(ctx context.Context) UploadWatermarkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UploadWatermark) *UploadWatermark {
		return &v
	}).(UploadWatermarkPtrOutput)
}

func (o UploadWatermarkOutput) Logs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadWatermark) *string { return v.Logs }).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkOutput) Metrics() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadWatermark) *string { return v.Metrics }).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkOutput) Usages() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadWatermark) *string { return v.Usages }).(pulumi.StringPtrOutput)
}

type UploadWatermarkPtrOutput struct{ *pulumi.OutputState }

func (UploadWatermarkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadWatermark)(nil)).Elem()
}

func (o UploadWatermarkPtrOutput) ToUploadWatermarkPtrOutput() UploadWatermarkPtrOutput {
	return o
}

func (o UploadWatermarkPtrOutput) ToUploadWatermarkPtrOutputWithContext(ctx context.Context) UploadWatermarkPtrOutput {
	return o
}

func (o UploadWatermarkPtrOutput) Elem() UploadWatermarkOutput {
	return o.ApplyT(func(v *UploadWatermark) UploadWatermark {
		if v != nil {
			return *v
		}
		var ret UploadWatermark
		return ret
	}).(UploadWatermarkOutput)
}

func (o UploadWatermarkPtrOutput) Logs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadWatermark) *string {
		if v == nil {
			return nil
		}
		return v.Logs
	}).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkPtrOutput) Metrics() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadWatermark) *string {
		if v == nil {
			return nil
		}
		return v.Metrics
	}).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkPtrOutput) Usages() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadWatermark) *string {
		if v == nil {
			return nil
		}
		return v.Usages
	}).(pulumi.StringPtrOutput)
}

type UploadWatermarkResponse struct {
	Logs    *string `pulumi:"logs"`
	Metrics *string `pulumi:"metrics"`
	Usages  *string `pulumi:"usages"`
}

type UploadWatermarkResponseOutput struct{ *pulumi.OutputState }

func (UploadWatermarkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadWatermarkResponse)(nil)).Elem()
}

func (o UploadWatermarkResponseOutput) ToUploadWatermarkResponseOutput() UploadWatermarkResponseOutput {
	return o
}

func (o UploadWatermarkResponseOutput) ToUploadWatermarkResponseOutputWithContext(ctx context.Context) UploadWatermarkResponseOutput {
	return o
}

func (o UploadWatermarkResponseOutput) Logs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadWatermarkResponse) *string { return v.Logs }).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkResponseOutput) Metrics() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadWatermarkResponse) *string { return v.Metrics }).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkResponseOutput) Usages() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadWatermarkResponse) *string { return v.Usages }).(pulumi.StringPtrOutput)
}

type UploadWatermarkResponsePtrOutput struct{ *pulumi.OutputState }

func (UploadWatermarkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadWatermarkResponse)(nil)).Elem()
}

func (o UploadWatermarkResponsePtrOutput) ToUploadWatermarkResponsePtrOutput() UploadWatermarkResponsePtrOutput {
	return o
}

func (o UploadWatermarkResponsePtrOutput) ToUploadWatermarkResponsePtrOutputWithContext(ctx context.Context) UploadWatermarkResponsePtrOutput {
	return o
}

func (o UploadWatermarkResponsePtrOutput) Elem() UploadWatermarkResponseOutput {
	return o.ApplyT(func(v *UploadWatermarkResponse) UploadWatermarkResponse {
		if v != nil {
			return *v
		}
		var ret UploadWatermarkResponse
		return ret
	}).(UploadWatermarkResponseOutput)
}

func (o UploadWatermarkResponsePtrOutput) Logs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadWatermarkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Logs
	}).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkResponsePtrOutput) Metrics() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadWatermarkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Metrics
	}).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkResponsePtrOutput) Usages() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadWatermarkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Usages
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActiveDirectoryConnectorDNSDetailsOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryConnectorDNSDetailsResponseOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryConnectorDomainDetailsOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryConnectorDomainDetailsResponseOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryConnectorPropertiesOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryConnectorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryConnectorSpecOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryConnectorSpecResponseOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryConnectorStatusOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryConnectorStatusPtrOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryConnectorStatusResponseOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryConnectorStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryDomainControllerOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryDomainControllerPtrOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryDomainControllerArrayOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryDomainControllerResponseOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryDomainControllerResponsePtrOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryDomainControllerResponseArrayOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryDomainControllersOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryDomainControllersResponseOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryInformationOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryInformationPtrOutput{})
	pulumi.RegisterOutputType(BasicLoginInformationOutput{})
	pulumi.RegisterOutputType(BasicLoginInformationPtrOutput{})
	pulumi.RegisterOutputType(BasicLoginInformationResponseOutput{})
	pulumi.RegisterOutputType(BasicLoginInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(DataControllerPropertiesOutput{})
	pulumi.RegisterOutputType(DataControllerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(K8sResourceRequirementsOutput{})
	pulumi.RegisterOutputType(K8sResourceRequirementsPtrOutput{})
	pulumi.RegisterOutputType(K8sResourceRequirementsResponseOutput{})
	pulumi.RegisterOutputType(K8sResourceRequirementsResponsePtrOutput{})
	pulumi.RegisterOutputType(K8sSchedulingOutput{})
	pulumi.RegisterOutputType(K8sSchedulingPtrOutput{})
	pulumi.RegisterOutputType(K8sSchedulingOptionsOutput{})
	pulumi.RegisterOutputType(K8sSchedulingOptionsPtrOutput{})
	pulumi.RegisterOutputType(K8sSchedulingOptionsResponseOutput{})
	pulumi.RegisterOutputType(K8sSchedulingOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(K8sSchedulingResponseOutput{})
	pulumi.RegisterOutputType(K8sSchedulingResponsePtrOutput{})
	pulumi.RegisterOutputType(KeytabInformationOutput{})
	pulumi.RegisterOutputType(KeytabInformationPtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsWorkspaceConfigOutput{})
	pulumi.RegisterOutputType(LogAnalyticsWorkspaceConfigPtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsWorkspaceConfigResponseOutput{})
	pulumi.RegisterOutputType(LogAnalyticsWorkspaceConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(OnPremisePropertyOutput{})
	pulumi.RegisterOutputType(OnPremisePropertyPtrOutput{})
	pulumi.RegisterOutputType(OnPremisePropertyResponseOutput{})
	pulumi.RegisterOutputType(OnPremisePropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(PostgresInstancePropertiesOutput{})
	pulumi.RegisterOutputType(PostgresInstancePropertiesResponseOutput{})
	pulumi.RegisterOutputType(PostgresInstanceSkuOutput{})
	pulumi.RegisterOutputType(PostgresInstanceSkuPtrOutput{})
	pulumi.RegisterOutputType(PostgresInstanceSkuResponseOutput{})
	pulumi.RegisterOutputType(PostgresInstanceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceK8sRawOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceK8sRawPtrOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceK8sRawResponseOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceK8sRawResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceK8sSpecOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceK8sSpecPtrOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceK8sSpecResponseOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceK8sSpecResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlManagedInstancePropertiesOutput{})
	pulumi.RegisterOutputType(SqlManagedInstancePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceSkuOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceSkuPtrOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceSkuResponseOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlServerInstancePropertiesOutput{})
	pulumi.RegisterOutputType(SqlServerInstancePropertiesPtrOutput{})
	pulumi.RegisterOutputType(SqlServerInstancePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UploadServicePrincipalOutput{})
	pulumi.RegisterOutputType(UploadServicePrincipalPtrOutput{})
	pulumi.RegisterOutputType(UploadServicePrincipalResponseOutput{})
	pulumi.RegisterOutputType(UploadServicePrincipalResponsePtrOutput{})
	pulumi.RegisterOutputType(UploadWatermarkOutput{})
	pulumi.RegisterOutputType(UploadWatermarkPtrOutput{})
	pulumi.RegisterOutputType(UploadWatermarkResponseOutput{})
	pulumi.RegisterOutputType(UploadWatermarkResponsePtrOutput{})
}

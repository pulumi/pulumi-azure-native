


package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Address struct {
	Address1   string  `pulumi:"address1"`
	Address2   *string `pulumi:"address2"`
	City       string  `pulumi:"city"`
	Country    string  `pulumi:"country"`
	PostalCode string  `pulumi:"postalCode"`
	State      string  `pulumi:"state"`
}





type AddressInput interface {
	pulumi.Input

	ToAddressOutput() AddressOutput
	ToAddressOutputWithContext(context.Context) AddressOutput
}

type AddressArgs struct {
	Address1   pulumi.StringInput    `pulumi:"address1"`
	Address2   pulumi.StringPtrInput `pulumi:"address2"`
	City       pulumi.StringInput    `pulumi:"city"`
	Country    pulumi.StringInput    `pulumi:"country"`
	PostalCode pulumi.StringInput    `pulumi:"postalCode"`
	State      pulumi.StringInput    `pulumi:"state"`
}

func (AddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Address)(nil)).Elem()
}

func (i AddressArgs) ToAddressOutput() AddressOutput {
	return i.ToAddressOutputWithContext(context.Background())
}

func (i AddressArgs) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressOutput)
}

func (i AddressArgs) ToAddressPtrOutput() AddressPtrOutput {
	return i.ToAddressPtrOutputWithContext(context.Background())
}

func (i AddressArgs) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressOutput).ToAddressPtrOutputWithContext(ctx)
}









type AddressPtrInput interface {
	pulumi.Input

	ToAddressPtrOutput() AddressPtrOutput
	ToAddressPtrOutputWithContext(context.Context) AddressPtrOutput
}

type addressPtrType AddressArgs

func AddressPtr(v *AddressArgs) AddressPtrInput {
	return (*addressPtrType)(v)
}

func (*addressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (i *addressPtrType) ToAddressPtrOutput() AddressPtrOutput {
	return i.ToAddressPtrOutputWithContext(context.Background())
}

func (i *addressPtrType) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPtrOutput)
}

type AddressOutput struct{ *pulumi.OutputState }

func (AddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Address)(nil)).Elem()
}

func (o AddressOutput) ToAddressOutput() AddressOutput {
	return o
}

func (o AddressOutput) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return o
}

func (o AddressOutput) ToAddressPtrOutput() AddressPtrOutput {
	return o.ToAddressPtrOutputWithContext(context.Background())
}

func (o AddressOutput) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Address) *Address {
		return &v
	}).(AddressPtrOutput)
}

func (o AddressOutput) Address1() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.Address1 }).(pulumi.StringOutput)
}

func (o AddressOutput) Address2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Address) *string { return v.Address2 }).(pulumi.StringPtrOutput)
}

func (o AddressOutput) City() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.City }).(pulumi.StringOutput)
}

func (o AddressOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.Country }).(pulumi.StringOutput)
}

func (o AddressOutput) PostalCode() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.PostalCode }).(pulumi.StringOutput)
}

func (o AddressOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.State }).(pulumi.StringOutput)
}

type AddressPtrOutput struct{ *pulumi.OutputState }

func (AddressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (o AddressPtrOutput) ToAddressPtrOutput() AddressPtrOutput {
	return o
}

func (o AddressPtrOutput) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return o
}

func (o AddressPtrOutput) Elem() AddressOutput {
	return o.ApplyT(func(v *Address) Address {
		if v != nil {
			return *v
		}
		var ret Address
		return ret
	}).(AddressOutput)
}

func (o AddressPtrOutput) Address1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return &v.Address1
	}).(pulumi.StringPtrOutput)
}

func (o AddressPtrOutput) Address2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return v.Address2
	}).(pulumi.StringPtrOutput)
}

func (o AddressPtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return &v.City
	}).(pulumi.StringPtrOutput)
}

func (o AddressPtrOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return &v.Country
	}).(pulumi.StringPtrOutput)
}

func (o AddressPtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return &v.PostalCode
	}).(pulumi.StringPtrOutput)
}

func (o AddressPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type Contact struct {
	AddressMailing *Address `pulumi:"addressMailing"`
	Email          string   `pulumi:"email"`
	Fax            *string  `pulumi:"fax"`
	JobTitle       *string  `pulumi:"jobTitle"`
	NameFirst      string   `pulumi:"nameFirst"`
	NameLast       string   `pulumi:"nameLast"`
	NameMiddle     *string  `pulumi:"nameMiddle"`
	Organization   *string  `pulumi:"organization"`
	Phone          string   `pulumi:"phone"`
}





type ContactInput interface {
	pulumi.Input

	ToContactOutput() ContactOutput
	ToContactOutputWithContext(context.Context) ContactOutput
}

type ContactArgs struct {
	AddressMailing AddressPtrInput       `pulumi:"addressMailing"`
	Email          pulumi.StringInput    `pulumi:"email"`
	Fax            pulumi.StringPtrInput `pulumi:"fax"`
	JobTitle       pulumi.StringPtrInput `pulumi:"jobTitle"`
	NameFirst      pulumi.StringInput    `pulumi:"nameFirst"`
	NameLast       pulumi.StringInput    `pulumi:"nameLast"`
	NameMiddle     pulumi.StringPtrInput `pulumi:"nameMiddle"`
	Organization   pulumi.StringPtrInput `pulumi:"organization"`
	Phone          pulumi.StringInput    `pulumi:"phone"`
}

func (ContactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Contact)(nil)).Elem()
}

func (i ContactArgs) ToContactOutput() ContactOutput {
	return i.ToContactOutputWithContext(context.Background())
}

func (i ContactArgs) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactOutput)
}

func (i ContactArgs) ToContactPtrOutput() ContactPtrOutput {
	return i.ToContactPtrOutputWithContext(context.Background())
}

func (i ContactArgs) ToContactPtrOutputWithContext(ctx context.Context) ContactPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactOutput).ToContactPtrOutputWithContext(ctx)
}









type ContactPtrInput interface {
	pulumi.Input

	ToContactPtrOutput() ContactPtrOutput
	ToContactPtrOutputWithContext(context.Context) ContactPtrOutput
}

type contactPtrType ContactArgs

func ContactPtr(v *ContactArgs) ContactPtrInput {
	return (*contactPtrType)(v)
}

func (*contactPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Contact)(nil)).Elem()
}

func (i *contactPtrType) ToContactPtrOutput() ContactPtrOutput {
	return i.ToContactPtrOutputWithContext(context.Background())
}

func (i *contactPtrType) ToContactPtrOutputWithContext(ctx context.Context) ContactPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactPtrOutput)
}

type ContactOutput struct{ *pulumi.OutputState }

func (ContactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Contact)(nil)).Elem()
}

func (o ContactOutput) ToContactOutput() ContactOutput {
	return o
}

func (o ContactOutput) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return o
}

func (o ContactOutput) ToContactPtrOutput() ContactPtrOutput {
	return o.ToContactPtrOutputWithContext(context.Background())
}

func (o ContactOutput) ToContactPtrOutputWithContext(ctx context.Context) ContactPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Contact) *Contact {
		return &v
	}).(ContactPtrOutput)
}

func (o ContactOutput) AddressMailing() AddressPtrOutput {
	return o.ApplyT(func(v Contact) *Address { return v.AddressMailing }).(AddressPtrOutput)
}

func (o ContactOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v Contact) string { return v.Email }).(pulumi.StringOutput)
}

func (o ContactOutput) Fax() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Contact) *string { return v.Fax }).(pulumi.StringPtrOutput)
}

func (o ContactOutput) JobTitle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Contact) *string { return v.JobTitle }).(pulumi.StringPtrOutput)
}

func (o ContactOutput) NameFirst() pulumi.StringOutput {
	return o.ApplyT(func(v Contact) string { return v.NameFirst }).(pulumi.StringOutput)
}

func (o ContactOutput) NameLast() pulumi.StringOutput {
	return o.ApplyT(func(v Contact) string { return v.NameLast }).(pulumi.StringOutput)
}

func (o ContactOutput) NameMiddle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Contact) *string { return v.NameMiddle }).(pulumi.StringPtrOutput)
}

func (o ContactOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Contact) *string { return v.Organization }).(pulumi.StringPtrOutput)
}

func (o ContactOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v Contact) string { return v.Phone }).(pulumi.StringOutput)
}

type ContactPtrOutput struct{ *pulumi.OutputState }

func (ContactPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Contact)(nil)).Elem()
}

func (o ContactPtrOutput) ToContactPtrOutput() ContactPtrOutput {
	return o
}

func (o ContactPtrOutput) ToContactPtrOutputWithContext(ctx context.Context) ContactPtrOutput {
	return o
}

func (o ContactPtrOutput) Elem() ContactOutput {
	return o.ApplyT(func(v *Contact) Contact {
		if v != nil {
			return *v
		}
		var ret Contact
		return ret
	}).(ContactOutput)
}

func (o ContactPtrOutput) AddressMailing() AddressPtrOutput {
	return o.ApplyT(func(v *Contact) *Address {
		if v == nil {
			return nil
		}
		return v.AddressMailing
	}).(AddressPtrOutput)
}

func (o ContactPtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Contact) *string {
		if v == nil {
			return nil
		}
		return &v.Email
	}).(pulumi.StringPtrOutput)
}

func (o ContactPtrOutput) Fax() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Contact) *string {
		if v == nil {
			return nil
		}
		return v.Fax
	}).(pulumi.StringPtrOutput)
}

func (o ContactPtrOutput) JobTitle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Contact) *string {
		if v == nil {
			return nil
		}
		return v.JobTitle
	}).(pulumi.StringPtrOutput)
}

func (o ContactPtrOutput) NameFirst() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Contact) *string {
		if v == nil {
			return nil
		}
		return &v.NameFirst
	}).(pulumi.StringPtrOutput)
}

func (o ContactPtrOutput) NameLast() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Contact) *string {
		if v == nil {
			return nil
		}
		return &v.NameLast
	}).(pulumi.StringPtrOutput)
}

func (o ContactPtrOutput) NameMiddle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Contact) *string {
		if v == nil {
			return nil
		}
		return v.NameMiddle
	}).(pulumi.StringPtrOutput)
}

func (o ContactPtrOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Contact) *string {
		if v == nil {
			return nil
		}
		return v.Organization
	}).(pulumi.StringPtrOutput)
}

func (o ContactPtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Contact) *string {
		if v == nil {
			return nil
		}
		return &v.Phone
	}).(pulumi.StringPtrOutput)
}

type DomainPurchaseConsent struct {
	AgreedAt      *string  `pulumi:"agreedAt"`
	AgreedBy      *string  `pulumi:"agreedBy"`
	AgreementKeys []string `pulumi:"agreementKeys"`
}





type DomainPurchaseConsentInput interface {
	pulumi.Input

	ToDomainPurchaseConsentOutput() DomainPurchaseConsentOutput
	ToDomainPurchaseConsentOutputWithContext(context.Context) DomainPurchaseConsentOutput
}

type DomainPurchaseConsentArgs struct {
	AgreedAt      pulumi.StringPtrInput   `pulumi:"agreedAt"`
	AgreedBy      pulumi.StringPtrInput   `pulumi:"agreedBy"`
	AgreementKeys pulumi.StringArrayInput `pulumi:"agreementKeys"`
}

func (DomainPurchaseConsentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainPurchaseConsent)(nil)).Elem()
}

func (i DomainPurchaseConsentArgs) ToDomainPurchaseConsentOutput() DomainPurchaseConsentOutput {
	return i.ToDomainPurchaseConsentOutputWithContext(context.Background())
}

func (i DomainPurchaseConsentArgs) ToDomainPurchaseConsentOutputWithContext(ctx context.Context) DomainPurchaseConsentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPurchaseConsentOutput)
}

func (i DomainPurchaseConsentArgs) ToDomainPurchaseConsentPtrOutput() DomainPurchaseConsentPtrOutput {
	return i.ToDomainPurchaseConsentPtrOutputWithContext(context.Background())
}

func (i DomainPurchaseConsentArgs) ToDomainPurchaseConsentPtrOutputWithContext(ctx context.Context) DomainPurchaseConsentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPurchaseConsentOutput).ToDomainPurchaseConsentPtrOutputWithContext(ctx)
}









type DomainPurchaseConsentPtrInput interface {
	pulumi.Input

	ToDomainPurchaseConsentPtrOutput() DomainPurchaseConsentPtrOutput
	ToDomainPurchaseConsentPtrOutputWithContext(context.Context) DomainPurchaseConsentPtrOutput
}

type domainPurchaseConsentPtrType DomainPurchaseConsentArgs

func DomainPurchaseConsentPtr(v *DomainPurchaseConsentArgs) DomainPurchaseConsentPtrInput {
	return (*domainPurchaseConsentPtrType)(v)
}

func (*domainPurchaseConsentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainPurchaseConsent)(nil)).Elem()
}

func (i *domainPurchaseConsentPtrType) ToDomainPurchaseConsentPtrOutput() DomainPurchaseConsentPtrOutput {
	return i.ToDomainPurchaseConsentPtrOutputWithContext(context.Background())
}

func (i *domainPurchaseConsentPtrType) ToDomainPurchaseConsentPtrOutputWithContext(ctx context.Context) DomainPurchaseConsentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPurchaseConsentPtrOutput)
}

type DomainPurchaseConsentOutput struct{ *pulumi.OutputState }

func (DomainPurchaseConsentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainPurchaseConsent)(nil)).Elem()
}

func (o DomainPurchaseConsentOutput) ToDomainPurchaseConsentOutput() DomainPurchaseConsentOutput {
	return o
}

func (o DomainPurchaseConsentOutput) ToDomainPurchaseConsentOutputWithContext(ctx context.Context) DomainPurchaseConsentOutput {
	return o
}

func (o DomainPurchaseConsentOutput) ToDomainPurchaseConsentPtrOutput() DomainPurchaseConsentPtrOutput {
	return o.ToDomainPurchaseConsentPtrOutputWithContext(context.Background())
}

func (o DomainPurchaseConsentOutput) ToDomainPurchaseConsentPtrOutputWithContext(ctx context.Context) DomainPurchaseConsentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainPurchaseConsent) *DomainPurchaseConsent {
		return &v
	}).(DomainPurchaseConsentPtrOutput)
}

func (o DomainPurchaseConsentOutput) AgreedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainPurchaseConsent) *string { return v.AgreedAt }).(pulumi.StringPtrOutput)
}

func (o DomainPurchaseConsentOutput) AgreedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainPurchaseConsent) *string { return v.AgreedBy }).(pulumi.StringPtrOutput)
}

func (o DomainPurchaseConsentOutput) AgreementKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DomainPurchaseConsent) []string { return v.AgreementKeys }).(pulumi.StringArrayOutput)
}

type DomainPurchaseConsentPtrOutput struct{ *pulumi.OutputState }

func (DomainPurchaseConsentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainPurchaseConsent)(nil)).Elem()
}

func (o DomainPurchaseConsentPtrOutput) ToDomainPurchaseConsentPtrOutput() DomainPurchaseConsentPtrOutput {
	return o
}

func (o DomainPurchaseConsentPtrOutput) ToDomainPurchaseConsentPtrOutputWithContext(ctx context.Context) DomainPurchaseConsentPtrOutput {
	return o
}

func (o DomainPurchaseConsentPtrOutput) Elem() DomainPurchaseConsentOutput {
	return o.ApplyT(func(v *DomainPurchaseConsent) DomainPurchaseConsent {
		if v != nil {
			return *v
		}
		var ret DomainPurchaseConsent
		return ret
	}).(DomainPurchaseConsentOutput)
}

func (o DomainPurchaseConsentPtrOutput) AgreedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainPurchaseConsent) *string {
		if v == nil {
			return nil
		}
		return v.AgreedAt
	}).(pulumi.StringPtrOutput)
}

func (o DomainPurchaseConsentPtrOutput) AgreedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainPurchaseConsent) *string {
		if v == nil {
			return nil
		}
		return v.AgreedBy
	}).(pulumi.StringPtrOutput)
}

func (o DomainPurchaseConsentPtrOutput) AgreementKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainPurchaseConsent) []string {
		if v == nil {
			return nil
		}
		return v.AgreementKeys
	}).(pulumi.StringArrayOutput)
}

type HostNameResponse struct {
	AzureResourceName           *string  `pulumi:"azureResourceName"`
	AzureResourceType           *string  `pulumi:"azureResourceType"`
	CustomHostNameDnsRecordType *string  `pulumi:"customHostNameDnsRecordType"`
	HostNameType                *string  `pulumi:"hostNameType"`
	Name                        *string  `pulumi:"name"`
	SiteNames                   []string `pulumi:"siteNames"`
}





type HostNameResponseInput interface {
	pulumi.Input

	ToHostNameResponseOutput() HostNameResponseOutput
	ToHostNameResponseOutputWithContext(context.Context) HostNameResponseOutput
}

type HostNameResponseArgs struct {
	AzureResourceName           pulumi.StringPtrInput   `pulumi:"azureResourceName"`
	AzureResourceType           pulumi.StringPtrInput   `pulumi:"azureResourceType"`
	CustomHostNameDnsRecordType pulumi.StringPtrInput   `pulumi:"customHostNameDnsRecordType"`
	HostNameType                pulumi.StringPtrInput   `pulumi:"hostNameType"`
	Name                        pulumi.StringPtrInput   `pulumi:"name"`
	SiteNames                   pulumi.StringArrayInput `pulumi:"siteNames"`
}

func (HostNameResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameResponse)(nil)).Elem()
}

func (i HostNameResponseArgs) ToHostNameResponseOutput() HostNameResponseOutput {
	return i.ToHostNameResponseOutputWithContext(context.Background())
}

func (i HostNameResponseArgs) ToHostNameResponseOutputWithContext(ctx context.Context) HostNameResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostNameResponseOutput)
}





type HostNameResponseArrayInput interface {
	pulumi.Input

	ToHostNameResponseArrayOutput() HostNameResponseArrayOutput
	ToHostNameResponseArrayOutputWithContext(context.Context) HostNameResponseArrayOutput
}

type HostNameResponseArray []HostNameResponseInput

func (HostNameResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostNameResponse)(nil)).Elem()
}

func (i HostNameResponseArray) ToHostNameResponseArrayOutput() HostNameResponseArrayOutput {
	return i.ToHostNameResponseArrayOutputWithContext(context.Background())
}

func (i HostNameResponseArray) ToHostNameResponseArrayOutputWithContext(ctx context.Context) HostNameResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostNameResponseArrayOutput)
}

type HostNameResponseOutput struct{ *pulumi.OutputState }

func (HostNameResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameResponse)(nil)).Elem()
}

func (o HostNameResponseOutput) ToHostNameResponseOutput() HostNameResponseOutput {
	return o
}

func (o HostNameResponseOutput) ToHostNameResponseOutputWithContext(ctx context.Context) HostNameResponseOutput {
	return o
}

func (o HostNameResponseOutput) AzureResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameResponse) *string { return v.AzureResourceName }).(pulumi.StringPtrOutput)
}

func (o HostNameResponseOutput) AzureResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameResponse) *string { return v.AzureResourceType }).(pulumi.StringPtrOutput)
}

func (o HostNameResponseOutput) CustomHostNameDnsRecordType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameResponse) *string { return v.CustomHostNameDnsRecordType }).(pulumi.StringPtrOutput)
}

func (o HostNameResponseOutput) HostNameType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameResponse) *string { return v.HostNameType }).(pulumi.StringPtrOutput)
}

func (o HostNameResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HostNameResponseOutput) SiteNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v HostNameResponse) []string { return v.SiteNames }).(pulumi.StringArrayOutput)
}

type HostNameResponseArrayOutput struct{ *pulumi.OutputState }

func (HostNameResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostNameResponse)(nil)).Elem()
}

func (o HostNameResponseArrayOutput) ToHostNameResponseArrayOutput() HostNameResponseArrayOutput {
	return o
}

func (o HostNameResponseArrayOutput) ToHostNameResponseArrayOutputWithContext(ctx context.Context) HostNameResponseArrayOutput {
	return o
}

func (o HostNameResponseArrayOutput) Index(i pulumi.IntInput) HostNameResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostNameResponse {
		return vs[0].([]HostNameResponse)[vs[1].(int)]
	}).(HostNameResponseOutput)
}

type NameIdentifierResponse struct {
	Name *string `pulumi:"name"`
}





type NameIdentifierResponseInput interface {
	pulumi.Input

	ToNameIdentifierResponseOutput() NameIdentifierResponseOutput
	ToNameIdentifierResponseOutputWithContext(context.Context) NameIdentifierResponseOutput
}

type NameIdentifierResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (NameIdentifierResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NameIdentifierResponse)(nil)).Elem()
}

func (i NameIdentifierResponseArgs) ToNameIdentifierResponseOutput() NameIdentifierResponseOutput {
	return i.ToNameIdentifierResponseOutputWithContext(context.Background())
}

func (i NameIdentifierResponseArgs) ToNameIdentifierResponseOutputWithContext(ctx context.Context) NameIdentifierResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameIdentifierResponseOutput)
}





type NameIdentifierResponseArrayInput interface {
	pulumi.Input

	ToNameIdentifierResponseArrayOutput() NameIdentifierResponseArrayOutput
	ToNameIdentifierResponseArrayOutputWithContext(context.Context) NameIdentifierResponseArrayOutput
}

type NameIdentifierResponseArray []NameIdentifierResponseInput

func (NameIdentifierResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameIdentifierResponse)(nil)).Elem()
}

func (i NameIdentifierResponseArray) ToNameIdentifierResponseArrayOutput() NameIdentifierResponseArrayOutput {
	return i.ToNameIdentifierResponseArrayOutputWithContext(context.Background())
}

func (i NameIdentifierResponseArray) ToNameIdentifierResponseArrayOutputWithContext(ctx context.Context) NameIdentifierResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameIdentifierResponseArrayOutput)
}

type NameIdentifierResponseOutput struct{ *pulumi.OutputState }

func (NameIdentifierResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NameIdentifierResponse)(nil)).Elem()
}

func (o NameIdentifierResponseOutput) ToNameIdentifierResponseOutput() NameIdentifierResponseOutput {
	return o
}

func (o NameIdentifierResponseOutput) ToNameIdentifierResponseOutputWithContext(ctx context.Context) NameIdentifierResponseOutput {
	return o
}

func (o NameIdentifierResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameIdentifierResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type NameIdentifierResponseArrayOutput struct{ *pulumi.OutputState }

func (NameIdentifierResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameIdentifierResponse)(nil)).Elem()
}

func (o NameIdentifierResponseArrayOutput) ToNameIdentifierResponseArrayOutput() NameIdentifierResponseArrayOutput {
	return o
}

func (o NameIdentifierResponseArrayOutput) ToNameIdentifierResponseArrayOutputWithContext(ctx context.Context) NameIdentifierResponseArrayOutput {
	return o
}

func (o NameIdentifierResponseArrayOutput) Index(i pulumi.IntInput) NameIdentifierResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NameIdentifierResponse {
		return vs[0].([]NameIdentifierResponse)[vs[1].(int)]
	}).(NameIdentifierResponseOutput)
}

type TldLegalAgreementResponse struct {
	AgreementKey string  `pulumi:"agreementKey"`
	Content      string  `pulumi:"content"`
	Title        string  `pulumi:"title"`
	Url          *string `pulumi:"url"`
}





type TldLegalAgreementResponseInput interface {
	pulumi.Input

	ToTldLegalAgreementResponseOutput() TldLegalAgreementResponseOutput
	ToTldLegalAgreementResponseOutputWithContext(context.Context) TldLegalAgreementResponseOutput
}

type TldLegalAgreementResponseArgs struct {
	AgreementKey pulumi.StringInput    `pulumi:"agreementKey"`
	Content      pulumi.StringInput    `pulumi:"content"`
	Title        pulumi.StringInput    `pulumi:"title"`
	Url          pulumi.StringPtrInput `pulumi:"url"`
}

func (TldLegalAgreementResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TldLegalAgreementResponse)(nil)).Elem()
}

func (i TldLegalAgreementResponseArgs) ToTldLegalAgreementResponseOutput() TldLegalAgreementResponseOutput {
	return i.ToTldLegalAgreementResponseOutputWithContext(context.Background())
}

func (i TldLegalAgreementResponseArgs) ToTldLegalAgreementResponseOutputWithContext(ctx context.Context) TldLegalAgreementResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TldLegalAgreementResponseOutput)
}





type TldLegalAgreementResponseArrayInput interface {
	pulumi.Input

	ToTldLegalAgreementResponseArrayOutput() TldLegalAgreementResponseArrayOutput
	ToTldLegalAgreementResponseArrayOutputWithContext(context.Context) TldLegalAgreementResponseArrayOutput
}

type TldLegalAgreementResponseArray []TldLegalAgreementResponseInput

func (TldLegalAgreementResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TldLegalAgreementResponse)(nil)).Elem()
}

func (i TldLegalAgreementResponseArray) ToTldLegalAgreementResponseArrayOutput() TldLegalAgreementResponseArrayOutput {
	return i.ToTldLegalAgreementResponseArrayOutputWithContext(context.Background())
}

func (i TldLegalAgreementResponseArray) ToTldLegalAgreementResponseArrayOutputWithContext(ctx context.Context) TldLegalAgreementResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TldLegalAgreementResponseArrayOutput)
}

type TldLegalAgreementResponseOutput struct{ *pulumi.OutputState }

func (TldLegalAgreementResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TldLegalAgreementResponse)(nil)).Elem()
}

func (o TldLegalAgreementResponseOutput) ToTldLegalAgreementResponseOutput() TldLegalAgreementResponseOutput {
	return o
}

func (o TldLegalAgreementResponseOutput) ToTldLegalAgreementResponseOutputWithContext(ctx context.Context) TldLegalAgreementResponseOutput {
	return o
}

func (o TldLegalAgreementResponseOutput) AgreementKey() pulumi.StringOutput {
	return o.ApplyT(func(v TldLegalAgreementResponse) string { return v.AgreementKey }).(pulumi.StringOutput)
}

func (o TldLegalAgreementResponseOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v TldLegalAgreementResponse) string { return v.Content }).(pulumi.StringOutput)
}

func (o TldLegalAgreementResponseOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v TldLegalAgreementResponse) string { return v.Title }).(pulumi.StringOutput)
}

func (o TldLegalAgreementResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TldLegalAgreementResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type TldLegalAgreementResponseArrayOutput struct{ *pulumi.OutputState }

func (TldLegalAgreementResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TldLegalAgreementResponse)(nil)).Elem()
}

func (o TldLegalAgreementResponseArrayOutput) ToTldLegalAgreementResponseArrayOutput() TldLegalAgreementResponseArrayOutput {
	return o
}

func (o TldLegalAgreementResponseArrayOutput) ToTldLegalAgreementResponseArrayOutputWithContext(ctx context.Context) TldLegalAgreementResponseArrayOutput {
	return o
}

func (o TldLegalAgreementResponseArrayOutput) Index(i pulumi.IntInput) TldLegalAgreementResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TldLegalAgreementResponse {
		return vs[0].([]TldLegalAgreementResponse)[vs[1].(int)]
	}).(TldLegalAgreementResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AddressOutput{})
	pulumi.RegisterOutputType(AddressPtrOutput{})
	pulumi.RegisterOutputType(ContactOutput{})
	pulumi.RegisterOutputType(ContactPtrOutput{})
	pulumi.RegisterOutputType(DomainPurchaseConsentOutput{})
	pulumi.RegisterOutputType(DomainPurchaseConsentPtrOutput{})
	pulumi.RegisterOutputType(HostNameResponseOutput{})
	pulumi.RegisterOutputType(HostNameResponseArrayOutput{})
	pulumi.RegisterOutputType(NameIdentifierResponseOutput{})
	pulumi.RegisterOutputType(NameIdentifierResponseArrayOutput{})
	pulumi.RegisterOutputType(TldLegalAgreementResponseOutput{})
	pulumi.RegisterOutputType(TldLegalAgreementResponseArrayOutput{})
}




package domainregistration

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

func (o DomainPurchaseConsentOutput) AgreedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainPurchaseConsent) *string { return v.AgreedAt }).(pulumi.StringPtrOutput)
}

func (o DomainPurchaseConsentOutput) AgreedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainPurchaseConsent) *string { return v.AgreedBy }).(pulumi.StringPtrOutput)
}

func (o DomainPurchaseConsentOutput) AgreementKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DomainPurchaseConsent) []string { return v.AgreementKeys }).(pulumi.StringArrayOutput)
}

type HostNameResponse struct {
	AzureResourceName           *string  `pulumi:"azureResourceName"`
	AzureResourceType           *string  `pulumi:"azureResourceType"`
	CustomHostNameDnsRecordType *string  `pulumi:"customHostNameDnsRecordType"`
	HostNameType                *string  `pulumi:"hostNameType"`
	Name                        *string  `pulumi:"name"`
	SiteNames                   []string `pulumi:"siteNames"`
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

type TldLegalAgreementResponse struct {
	AgreementKey string  `pulumi:"agreementKey"`
	Content      string  `pulumi:"content"`
	Title        string  `pulumi:"title"`
	Url          *string `pulumi:"url"`
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
	pulumi.RegisterOutputType(DomainPurchaseConsentOutput{})
	pulumi.RegisterOutputType(HostNameResponseOutput{})
	pulumi.RegisterOutputType(HostNameResponseArrayOutput{})
	pulumi.RegisterOutputType(NameIdentifierResponseOutput{})
	pulumi.RegisterOutputType(NameIdentifierResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TldLegalAgreementResponseOutput{})
	pulumi.RegisterOutputType(TldLegalAgreementResponseArrayOutput{})
}
